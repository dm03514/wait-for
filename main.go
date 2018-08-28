package main

import (
	"os"
	"time"

	"github.com/dm03514/wait-for/cmd"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

func main() {
	app := cli.NewApp()

	app.Name = "wait-for"
	app.Usage = "wait for a service to become available"
	app.Commands = []cli.Command{
		cmd.HTTPCommand,
		cmd.PostgresCommand,
	}
	app.Flags = []cli.Flag{
		cli.DurationFlag{
			Name:  "timeout, t",
			Value: time.Minute * 5,
			Usage: "duration to wait until marking as failure and returning",
		},
		cli.DurationFlag{
			Name:  "poll-interval, pi",
			Value: time.Millisecond * 100,
			Usage: "interval",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
