package main

import (
	"os"

	log "github.com/Sirupsen/logrus"

	"github.com/ernoaapa/fetch-ssh-keys/fetch"
	"github.com/ernoaapa/fetch-ssh-keys/output"
	"github.com/urfave/cli"
)

// StatsdConfig for statsd client
type StatsdConfig struct {
	Host       string `default:"localhost"`
	Port       int    `default:"8125"`
	Prefix     string `default:"stats"`
	MetricName string `envconfig:"metric_name",required:"true"`
}

func main() {
	app := cli.NewApp()
	app.Name = "fetch-ssh-keys"
	app.Usage = "Fetch user public SSH keys"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "output, o",
			Usage: "Output format. One of: ssh",
			Value: "ssh",
		},
	}
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
				cli.StringFlag{
					Name:  "team",
					Usage: "Return only members of one team",
				},
			},
			Action: func(c *cli.Context) error {
				if c.String("organization") == "" {
					log.Fatalln("You must give --organization value")
				}

				keys, err := fetch.GitHubKeys(c.String("organization"), fetch.GithubFetchParams{
					Token:             c.String("token"),
					TeamName:          c.String("team"),
					PublicMembersOnly: c.Bool("public-only"),
				})
				if err != nil {
					log.Fatalln("Failed to fetch keys", err)
				}

				println(output.Build(c.GlobalString("output"), keys))
				return nil
			},
		},
	}

	app.Run(os.Args)
}
