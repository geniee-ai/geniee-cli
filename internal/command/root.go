package command

import (
	"fmt"
	"os"

	"github.com/geniee-ai/geniee-cli/internal/rgb"
	"github.com/urfave/cli/v2"
)

var (
	ErrExpectAtLeastOneArg = rgb.Red.Sprintf("\nError: geniee expects at least one argument. Please check possible commands using \"geniee --help\" or \"cheesy -h\"\n")
)

func RootCmd(c *cli.Context) error {
	if len(os.Args) <= 1 {
		fmt.Println(ErrExpectAtLeastOneArg)
		cli.ShowAppHelp(c)
		// os.Exit(1)
	}
	return nil
}
