package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/dancannon/go-micro/server"
	"github.com/gin-gonic/gin"

	"github.com/dancannon/k8s_dev/services/hello-api/handler"
)

func main() {
	app := cli.NewApp()

	app.Name = "hello-api"
	app.Version = "0.1.0"
	app.Author = "Daniel Cannon"
	app.Email = "daniel@danielcannon.co.uk"
	app.Usage = "API for talking to hello-service"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "env",
			Value:  "dev",
			Usage:  "environment the application is running in",
			EnvVar: "APP_ENV",
		},
		cli.StringFlag{
			Name:   "region",
			Value:  "dev",
			Usage:  "region the application is running in",
			EnvVar: "APP_REGION",
		},
		cli.BoolFlag{
			Name:  "verbose, v",
			Usage: "sets log level to DEBUG",
		},

		cli.StringFlag{
			Name:  "registry",
			Value: "consul",
			Usage: "registry for discovery. kubernetes, consul, etc",
		},
		cli.StringFlag{
			Name:  "bind_address",
			Value: ":0",
			Usage: "bind address for the server. 127.0.0.1:8080",
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

		server.Registry = c.String("registry")
		server.BindAddress = c.String("bind_address")
		server.Init(false)

		startServer(c.String("bind_address"))
	}

	app.Run(os.Args)
}

func startServer(addr string) {
	server := gin.Default()

	server.GET("/ping", handler.Ping)
	server.GET("/hello/:name", handler.Hello)

	// Run server
	server.Run(addr)
}

func init() {
	cli.VersionFlag = cli.BoolFlag{
		Name:  "version",
		Usage: "print the version",
	}
}
