package cmd

import (
	"github.com/dm03514/wait-for/poller"
	"github.com/dm03514/wait-for/poller/net"
	"github.com/urfave/cli"
)

var NetCommand = cli.Command{
	Name: "net",
	Action: func(c *cli.Context) error {
		n, err := net.New(
			c.String("network"),
			c.String("address"),
		)
		if err != nil {
			return err
		}

		p := poller.Poller{
			Timeout:    c.GlobalDuration("timeout"),
			Interval:   c.GlobalDuration("poll-interval"),
			CheckReady: n.CheckReady,
		}

		return p.WaitReady()
	},
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "network, n",
			Value: "tcp",
			Usage: "Protocol: (\"tcp\", \"udp\"), any value that go net.Dial accepts",
		},
		cli.StringFlag{
			Name:  "address, a",
			Value: "",
			Usage: "address and port ie localhost:8000",
		},
	},
}
