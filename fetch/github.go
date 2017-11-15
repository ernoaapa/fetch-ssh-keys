package fetch

import (
	"context"
	"fmt"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// GithubFetchParams contains all parameters what are required for fetching tokens from GitHub
type GithubFetchParams struct {
	Token             string
	TeamNames         []string
	PublicMembersOnly bool
}

// GitHubOrganisationKeys fetches organization users public SSH key from GitHub
func GitHubOrganisationKeys(organizationName string, params GithubFetchParams) (map[string][]string, error) {
	client := getClient(params.Token)
	users, err := fetchUsers(client, organizationName, params)
	if err != nil {
		return map[string][]string{}, err
	}
	log.Debugf("Users found: %d", len(users))

	usernames := []string{}
	for _, user := range users {
		usernames = append(usernames, *user.Login)
	}

	return fetchUserKeys(client, usernames, params.Token)
}

// GitHubUsers fetches users public SSH keys from GitHub
func GitHubUsers(usernames []string, token string) (map[string][]string, error) {
	client := getClient(token)
	return fetchUserKeys(client, usernames, token)
}

func getClient(token string) *github.Client {
	if len(token) > 0 {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		return github.NewClient(oauth2.NewClient(oauth2.NoContext, ts))
	}
	return github.NewClient(nil)
}

func fetchUsers(client *github.Client, organizationName string, params GithubFetchParams) ([]*github.User, error) {
	ctx := context.Background()

	if len(params.TeamNames) > 0 {
		var users []*github.User
		for _, teamName := range params.TeamNames {
			teamID, err := resolveTeamID(client, organizationName, teamName)
			if err != nil {
				return []*github.User{}, err
			}

			teamUsers, _, err := client.Organizations.ListTeamMembers(ctx, teamID, &github.OrganizationListTeamMembersOptions{})
			if err != nil {
				return []*github.User{}, err
			}

			users = append(users, teamUsers...)
		}

		return users, nil
	}

	users, _, err := client.Organizations.ListMembers(ctx, organizationName, &github.ListMembersOptions{
		PublicOnly: params.PublicMembersOnly,
	})
	return users, err
}

func resolveTeamID(client *github.Client, organizationName, teamName string) (int, error) {
	ctx := context.Background()

	teams, _, err := client.Organizations.ListTeams(ctx, organizationName, &github.ListOptions{})
	if err != nil {
		return -1, err
	}

	for _, team := range teams {
		if strings.EqualFold(*team.Name, teamName) {
			return *team.ID, nil
		}
	}

	return -1, fmt.Errorf("Unable to find team [%s] from organization [%s]", teamName, organizationName)
}

func fetchUserKeys(client *github.Client, usernames []string, token string) (map[string][]string, error) {
	ctx := context.Background()

	result := map[string][]string{}
	for _, username := range usernames {
		keys, _, err := client.Users.ListKeys(ctx, username, &github.ListOptions{})
		if err != nil {
			return map[string][]string{}, err
		}

		result[username] = make([]string, len(keys))

		for index, key := range keys {
			result[username][index] = *key.Key
		}
	}

	return result, nil
}
