package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/pkg/errors"

	"github.com/ernoaapa/fetch-ssh-keys/fetch"
	"github.com/ernoaapa/fetch-ssh-keys/output"
	"github.com/ernoaapa/fetch-ssh-keys/utils"
	"github.com/urfave/cli"
)

// Version string to be set at compile time via command line (-ldflags "-X main.VersionString=1.2.3")
var (
	VersionString string
)

func main() {
	app := cli.NewApp()
	app.Name = "fetch-ssh-keys"
	app.Usage = "Fetch user public SSH keys"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "format, f",
			Usage: "Output format. One of: ssh",
			Value: "ssh",
		},
		cli.StringFlag{
			Name:  "file-mode",
			Usage: "File permissions for file",
			Value: "0600",
		},
		cli.StringFlag{
			Name:  "comment",
			Usage: "Include `COMMENT` at top and bottom",
		},
	}
	app.Version = VersionString
	app.Commands = []cli.Command{
		{
			Name:  "github",
			Usage: "Get user GitHub public SSH key",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "organization, o",
					Usage:  "GitHub organization which users public keys to get",
					EnvVar: "GITHUB_ORGANIZATION",
				},
				cli.StringFlag{
					Name:   "token, t",
					Usage:  "GitHub access token",
					EnvVar: "GITHUB_TOKEN",
				},
				cli.BoolFlag{
					Name:  "public-only",
					Usage: "Return only public members of organization",
				},
				cli.StringSliceFlag{
					Name:  "team",
					Usage: "Return only members of `TEAM` (this option can be used multiple times for multiple teams)",
				},
				cli.StringSliceFlag{
					Name:  "user",
					Usage: "Return given `USER` public ssh keys (this option can be used multiple times for multiple users)",
				},
				cli.StringSliceFlag{
					Name:  "deploy-key",
					Usage: "Return given `OWNER/REPO` deploy ssh keys (this option can be used multiple times for multiple repositories)",
				},
			},
			Action: func(c *cli.Context) error {
				var (
					token        = c.String("token")
					organisation = c.String("organization")
					teams        = c.StringSlice("team")
					users        = c.StringSlice("user")
					ownerRepos   = c.StringSlice("deploy-key")
					publicOnly   = c.Bool("public-only")

					deployKeys map[string][]string
					orgKeys    map[string][]string
					userKeys   map[string][]string

					target   = c.Args().Get(0)
					fileMode = os.FileMode(c.GlobalInt("file-mode"))
					format   = c.GlobalString("format")
					comment  = c.GlobalString("comment")

					err error
				)

				if organisation == "" && len(users) == 0 && len(ownerRepos) == 0 {
					return fmt.Errorf("You must give either --organisation or --user or --deploy-key parameter")
				}

				if c.IsSet("organization") {
					orgKeys, err = fetch.GitHubOrganisationKeys(organisation, fetch.GithubFetchParams{
						Token:             token,
						TeamNames:         teams,
						PublicMembersOnly: publicOnly,
					})
					if err != nil {
						return errors.Wrapf(err, "Failed to fetch keys from organisation %s", organisation)
					}
				}

				if c.IsSet("user") {
					userKeys, err = fetch.GitHubUsers(users, token)
					if err != nil {
						return errors.Wrap(err, "Failed to fetch GitHub user(s) keys")
					}
				}

				if c.IsSet("deploy-key") {
					deployKeys, err = fetch.GitHubDeployKeys(ownerRepos, token)
					if err != nil {
						return errors.Wrap(err, "Failed to fetch GitHub repositories' deploy keys")
					}
				}

				return output.Write(format, target, fileMode, utils.MergeKeys(orgKeys, userKeys, deployKeys), comment)
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
