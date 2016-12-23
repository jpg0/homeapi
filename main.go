package main

import (
	"io/ioutil"
	"encoding/json"
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
		},
		cli.StringFlag{
			Name: "port",
			Usage: "Post to listen on",
		},
		cli.StringFlag{
			Name: "loglevel",
			Usage: "Logging level",
			Value: "info",
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

func loadConfiguration(path string) (*Configuration, error) {
	rv := &Configuration{}
	file, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, errors.Annotate(err, "Failed to read config file")
	}

	err = json.Unmarshal(file, rv)

	if err != nil {
		return nil, errors.Annotate(err, "Failed to unmarshal config")
	}

	return rv, nil
}

type Configuration map[string]string