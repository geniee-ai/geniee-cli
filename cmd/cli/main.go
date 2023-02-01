package main

import (
	"fmt"

	"os"

	"github.com/geniee-ai/geniee-cli/internal/command"
	"github.com/geniee-ai/geniee-cli/version"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

const (
	AppName = "Geniee"
)

func init() {
	var logLevel string
	logLevel = os.Getenv("GENIEE_LOG_LEVEL")
	// LOG_LEVEL not set, let's default to debug
	switch logLevel {
	case "debug":
		logLevel = "debug"
	case "error":
		logLevel = "error"
	default:
		logLevel = "info"
	}
	// parse string, this is built-in feature of logrus
	ll, err := logrus.ParseLevel(logLevel)
	if err != nil {
		ll = logrus.DebugLevel
	}
	// set global log level
	logrus.SetLevel(ll)
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
}

func main() {

	cli.AppHelpTemplate = fmt.Sprintf(`%s
WEBSITE: http://geniee.io

SUPPORT: support@geniee.io

	`, cli.AppHelpTemplate)

	app := &cli.App{
		Name:    AppName,
		Version: version.Version,
		Usage:   "ask any questions directly from terminal",
		Action:  command.RootCmd,
		Commands: []*cli.Command{
			{
				Name:    "ask",
				Aliases: []string{"a"},
				Usage:   "Ask questions",
				Action:  command.AskCmd,
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "message", Aliases: []string{"m"}},
				},
			},
			{
				Name:    "login",
				Aliases: []string{"l"},
				Usage:   "Obtain and save credentials from cheesy web.",
				Action:  command.LoginCmd,
			},
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "Show version",
				Action:  command.VersionCmd,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
