package cmd

import (
	"github.com/dm03514/wait-for/poller"
	"github.com/urfave/cli"
)

var PostgresCommand = cli.Command{
	Name: "postgres",
	Action: func(c *cli.Context) error {
		postgres, err := poller.NewPostgres(
			c.String("connection-string"),
		)
		if err != nil {
			return err
		}

		p := poller.Poller{
			Timeout:    c.GlobalDuration("timeout"),
			Interval:   c.GlobalDuration("poll-interval"),
			CheckReady: postgres.CheckReady,
		}

		return p.WaitReady()
	},
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:   "connection-string, cs",
			Value:  "",
			Usage:  "psql connection string",
			EnvVar: "WAIT_FOR_POSTGRES_CONNECTION_STRING",
		},
	},
}
