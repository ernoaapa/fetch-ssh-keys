package fetch

import (
	"fmt"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// GithubFetchParams contains all parameters what are required for fetching tokens from GitHub
type GithubFetchParams struct {
	Token             string
	TeamName          string
	PublicMembersOnly bool
}

// GitHubKeys fetches organization users public SSH key from GitHub
func GitHubKeys(organizationName string, params GithubFetchParams) (map[string][]string, error) {
	client := getClient(params)
	users, err := fetchUsers(client, organizationName, params)
	if err != nil {
		return map[string][]string{}, err
	}
	log.Debugf("Users found: %d", len(users))

	result := map[string][]string{}
	for _, user := range users {
		username := *user.Login
		keys, _, err := client.Users.ListKeys(username, &github.ListOptions{})
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

func getClient(params GithubFetchParams) *github.Client {
	if len(params.Token) > 0 {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: params.Token},
		)
		return github.NewClient(oauth2.NewClient(oauth2.NoContext, ts))
	}
	return github.NewClient(nil)
}

func fetchUsers(client *github.Client, organizationName string, params GithubFetchParams) ([]*github.User, error) {
	if params.TeamName != "" {
		teamID, err := resolveTeamID(client, organizationName, params.TeamName)
		if err != nil {
			return []*github.User{}, err
		}
		users, _, err := client.Organizations.ListTeamMembers(teamID, &github.OrganizationListTeamMembersOptions{})
		return users, err
	}

	users, _, err := client.Organizations.ListMembers(organizationName, &github.ListMembersOptions{
		PublicOnly: params.PublicMembersOnly,
	})
	return users, err
}

func resolveTeamID(client *github.Client, organizationName, teamName string) (int, error) {
	teams, _, err := client.Organizations.ListTeams(organizationName, &github.ListOptions{})
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
