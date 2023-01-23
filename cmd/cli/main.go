package main

import (
	"fmt"

	"log"
	"os"

	"github.com/geniee-ai/geniee-cli/internal/command"
	"github.com/geniee-ai/geniee-cli/version"
	"github.com/urfave/cli/v2"
)

const (
	AppName = "Geniee"
)

func main() {
	cli.AppHelpTemplate = fmt.Sprintf(`%s
WEBSITE: http://geniee.io

SUPPORT: support@geniee.io

	`, cli.AppHelpTemplate)

	// cli.VersionFlag = &cli.BoolFlag{
	// 	Name:    "version",
	// 	Aliases: []string{"V"},
	// 	Usage:   "print only the version",
	// }

	// cli.VersionPrinter = func(cCtx *cli.Context) {
	// 	fmt.Printf("version=%s revision=%s\n", cCtx.App.Version, version.Version)
	// }

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
