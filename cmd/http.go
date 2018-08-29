package cmd

import (
	"github.com/dm03514/wait-for/poller"
	"github.com/dm03514/wait-for/poller/http"
	"github.com/urfave/cli"
)

var HTTPCommand = cli.Command{
	Name: "http",
	Action: func(c *cli.Context) error {
		h, err := http.New(
			c.String("method"),
			c.String("url"),
			c.String("body"),
		)
		if err != nil {
			return err
		}

		p := poller.Poller{
			Timeout:    c.GlobalDuration("timeout"),
			Interval:   c.GlobalDuration("poll-interval"),
			CheckReady: h.CheckReady,
		}

		return p.WaitReady()
	},
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "method, m",
			Value: "GET",
			Usage: "http request method to use for polling",
		},
		cli.StringFlag{
			Name:  "url",
			Value: "",
			Usage: "http uri to poll status of",
		},
		cli.StringFlag{
			Name:  "body",
			Value: "",
			Usage: "optional body to send",
		},
	},
}
