package main

import (
	"github.com/juju/errors"
	"os"
	"fmt"
	"github.com/urfave/cli"
	"strings"
	"github.com/Sirupsen/logrus"
)

func main() {
	app := cli.NewApp()
	app.Name = "homeapi"
	app.Usage = "API for Home Services"
	app.Version = "1.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "configfile",
			Usage: "Configuration File",
			EnvVar: "CONFIGFILE",
		},
		cli.StringFlag{
			Name: "port",
			Usage: "Post to listen on",
			EnvVar: "PORT",
		},
		cli.StringFlag{
			Name: "loglevel",
			Usage: "Logging level",
			Value: "info",
			EnvVar: "LOGLEVEL",
		},
	}
	app.Action = verbose(ConfigureAndStart)
	app.Run(os.Args)
}

func verbose(next func(*cli.Context) error) func(*cli.Context) error {
	return func(c *cli.Context) error {
		err := next(c)

		if err != nil {
			fmt.Println(errors.ErrorStack(err))
		}

		return err
	}
}

func initLogging(level string) error {
	switch strings.ToLower(level) {
	case "debug": logrus.SetLevel(logrus.DebugLevel)
	case "info": logrus.SetLevel(logrus.InfoLevel)
	case "warn": logrus.SetLevel(logrus.WarnLevel)
	case "error": logrus.SetLevel(logrus.ErrorLevel)
	case "fatal": logrus.SetLevel(logrus.FatalLevel)
	default:
		return errors.Errorf("Unknown logging level: %v", level)
	}

	return nil
}

func ConfigureAndStart(c *cli.Context) error {

	err := initLogging(c.String("loglevel"))

	if err != nil {
		return errors.Annotate(err, "Failed to initialize logging")
	}

	config, err := loadConfiguration(c.String("configfile"))

	if err != nil {
		return errors.Annotate(err, "Failed loading config")
	}

	setupHandlers();
	setupRouting(c.String("port"), config);

	return nil
}