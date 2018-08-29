package cmd

import (
	"github.com/dm03514/wait-for/poller"
	"github.com/dm03514/wait-for/poller/mysql"
	"github.com/urfave/cli"
)

var MySQLCommand = cli.Command{
	Name: "mysql",
	Action: func(c *cli.Context) error {
		postgres, err := mysql.New(
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
			Usage:  "mysql connection string",
			EnvVar: "WAIT_FOR_MYSQL_CONNECTION_STRING",
		},
	},
}
