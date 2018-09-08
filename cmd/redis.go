package cmd

import (
	"github.com/dm03514/wait-for/poller"
	"github.com/dm03514/wait-for/poller/redis"
	"github.com/urfave/cli"
)

var RedisCommand = cli.Command{
	Name: "redis",
	Action: func(c *cli.Context) error {
		r, err := redis.New(
			c.String("address"),
			c.String("password"),
			c.Int("db"),
		)
		if err != nil {
			return err
		}

		p := poller.Poller{
			Timeout:    c.GlobalDuration("timeout"),
			Interval:   c.GlobalDuration("poll-interval"),
			CheckReady: r.CheckReady,
		}

		return p.WaitReady()
	},
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "address, a",
			Value: "localhost:6379",
			Usage: "redis address and port \"addr:port\"",
		},
		cli.StringFlag{
			Name:  "password, p",
			Value: "",
			Usage: "redis password",
		},
		cli.IntFlag{
			Name:  "db",
			Value: 0,
			Usage: "redis db to connect to",
		},
	},
}
