package command

import (
	"fmt"

	"github.com/geniee-ai/geniee-cli/version"
	"github.com/urfave/cli/v2"
)

func VersionCmd(*cli.Context) error {
	fmt.Println(version.Version)
	return nil
}
