package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/asim/go-micro/server"
	"github.com/codegangsta/cli"

	"github.com/dancannon/k8s_dev/services/hello-service/handler"
)

func main() {
	app := cli.NewApp()

	app.Name = "app-hello-service"
	app.Version = "0.1.0"
	app.Author = "Daniel Cannon"
	app.Email = "daniel@danielcannon.co.uk"
	app.Usage = "Welcomes users"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "env",
			Value:  "dev",
			Usage:  "environment the service is running in, defaults to 'dev'",
			EnvVar: "APP_ENV",
		},
		cli.StringFlag{
			Name:   "region",
			Value:  "dev",
			Usage:  "region the service is running in, defaults to 'dev'",
			EnvVar: "APP_REGION",
		},
		cli.BoolFlag{
			Name:  "verbose, v",
			Usage: "sets log level to DEBUG",
		},
	}

	app.Action = func(c *cli.Context) {
		if c.Bool("verbose") {
			logrus.SetLevel(logrus.DebugLevel)
		}
		if c.String("env") != "dev" {
			logrus.SetFormatter(new(logrus.JSONFormatter))
		} else {
			logrus.SetFormatter(new(logrus.TextFormatter))
		}

		startServer()
	}

	app.Run(os.Args)
}

func startServer() {
	server.Name = "app.service.hello"
	server.Init()

	// Register Handlers
	server.Register(
		server.NewReceiver(new(handler.Hello)),
	)

	// Run server
	if err := server.Run(); err != nil {
		logrus.Fatal(err.Error())
	}
}

func init() {
	cli.VersionFlag = cli.BoolFlag{
		Name:  "version",
		Usage: "print the version",
	}
}
