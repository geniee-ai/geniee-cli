package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"os"
	"syscall"

	"github.com/fatih/color"
	"github.com/geniee-ai/geniee-cli/internal/helpers"
	"github.com/geniee-ai/geniee-cli/internal/rgb"
	"github.com/geniee-ai/geniee-cli/internal/token"

	"github.com/urfave/cli/v2"
	"golang.org/x/term"
)

type Credentials struct {
	Token string `json:"token"`
	Email string `json:"email"`
}

const (
	genieeCredsFolderName = ".geniee"
	genieeCredsFileName   = "config"
	genieeCredsFileType   = "json"
)

var (
	banner1 = fmt.Sprint(`
Geniee will store the token in plain text in the following file for use by subsequent commands:
	
Note: Access token generated on web UI will only visible once.
`)

	proceedAsk = fmt.Sprint(`
Do you want to proceed?
Only 'yes' will be accepted to confirm.
`)

	separator = fmt.Sprint(`

---------------------------------------------------------------------------------

`)

	banner2 = fmt.Sprint(`
Please login to geniee web ( https://geniee.io ) and generate API token.

Once you log into web, please generate access token.
	
Once token is generated, please enter the token generated below.
`)
)

var (
	cyan  = color.New(color.FgCyan, color.Bold)
	red   = color.New(color.FgRed, color.Bold)
	green = color.New(color.FgGreen, color.Bold)
	white = color.New(color.FgHiWhite, color.Bold)
)

func LoginCmd(*cli.Context) error {
	configPath := rgb.White.Sprintf("\t" + os.Getenv("HOME") + "/.geniee/" + genieeCredsFileName + "." + genieeCredsFileType)

	fmt.Println(banner1)
	fmt.Println(configPath)
	fmt.Println(proceedAsk)
	white.Printf("\tEnter a value: ")
	response := helpers.ReadInputString()
	fmt.Println(separator)

	switch response {
	case "yes":
		fmt.Println(banner2)

		white.Print("\tEnter your geniee web email: ")
		userEmail := helpers.ReadInputString()

		white.Print("\tEnter the access token: ")
		accessTokenByte, _ := term.ReadPassword(int(syscall.Stdin))

		cred := Credentials{
			Token: string(accessTokenByte),
			Email: userEmail,
		}

		pathToHome, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("could not find current user home directory: %v", err)
		}

		credentialsFile := fmt.Sprintf(genieeCredsFileName + "." + genieeCredsFileType)
		confFolderpath := fmt.Sprintf(pathToHome + "/" + genieeCredsFolderName)

		if !helpers.IsExists(confFolderpath) {
			err = helpers.CreateDir(pathToHome, genieeCredsFolderName)
			if err != nil {
				return fmt.Errorf("could not create cheesy conf directory: %v", err)
			}
		}

		if !helpers.IsExists(confFolderpath + "/" + credentialsFile) {
			err = helpers.CreateFile(confFolderpath, credentialsFile)
			if err != nil {
				return fmt.Errorf("could not create cheesy creds file: %v", err)
			}
		}

		file, _ := json.MarshalIndent(cred, "", " ")
		if err != nil {
			return fmt.Errorf("could not marshan indent: %s", err)
		}

		_ = ioutil.WriteFile(confFolderpath+"/"+credentialsFile, file, 0644)
		// err = helpers.WriteTofile(confFolderpath+"/"+credentialsFile, file, 0755)
		if err != nil {
			return fmt.Errorf("could not write to file file: %v", err)
		}

		// // TODO ( validate token )
		ok := token.ValidateToken(string(accessTokenByte))
		if !ok {
			red := color.New(color.FgRed, color.Bold).SprintFunc()
			return fmt.Errorf("%s. Please check if you entered correct access token", red("Invalid Token"))
		}

		green := color.New(color.FgGreen, color.Bold).SprintFunc()
		fmt.Printf("\nLogged in %s\n", green("successfully"))

	default:
		red := color.New(color.FgRed, color.Bold).SprintFunc()
		white := color.New(color.FgHiWhite, color.Bold).SprintFunc()
		fmt.Printf("%s, %s \n", red("ERROR"), white("Login Cancelled"))
		os.Exit(1)
	}

	return nil
}
