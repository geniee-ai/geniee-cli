package command

import (
	"encoding/json"
	"fmt"

	"os"
	"syscall"

	"github.com/fatih/color"
	"github.com/geniee-ai/geniee-cli/internal/helpers"
	"github.com/geniee-ai/geniee-cli/internal/rgb"
	"github.com/geniee-ai/geniee-cli/internal/token"

	"github.com/urfave/cli/v2"
	"golang.org/x/term"
)

type credentials struct {
	token string
	email string
}

const (
	cheesyCredsFolderName = ".cheesy"
	cheesyCredsFileName   = "credentials"
	cheesyCredsFileType   = "json"
)

var (
	banner = fmt.Sprint(`
Please login to geniee web ( https://geniee.io ) and generate API token using your browser.

After login is successful, Please go to settings and generate access
token.
	
Geniee will store the token in plain text in the following file for use by subsequent commands:
`)

	proceedAsk = fmt.Sprint(`
	
Do you want to proceed?
Only 'yes' will be accepted to confirm.
`)

	banner2 = fmt.Sprint(`

---------------------------------------------------------------------------------

Login to geniee.io and go to Settings > Create Token to generate a new token.
	
Note: Access token generated on web UI will only visible once.
	
---------------------------------------------------------------------------------

Once Token is generated in browser, copy and paste it into the prompt below.

Geniee will store the token in plain text in
the following file for use by subsequent commands:
`)
)

var (
	cyan  = color.New(color.FgCyan, color.Bold)
	red   = color.New(color.FgRed, color.Bold)
	green = color.New(color.FgGreen, color.Bold)
	white = color.New(color.FgHiWhite, color.Bold)
)

func LoginCmd(*cli.Context) error {
	configPath := rgb.White.Sprintf("\t" + os.Getenv("HOME") + "/.geniee/credentials.json\n")

	fmt.Println(banner)
	fmt.Println(configPath)
	fmt.Println(proceedAsk)
	white.Printf("Enter a value: ")
	response := helpers.ReadInputString()

	switch response {
	case "yes":
		fmt.Println(banner2)
		fmt.Println(configPath)

		white.Print("  Enter your geniee web email: ")
		userEmail := helpers.ReadInputString()

		white.Print("  Enter the access token: ")
		accessTokenByte, _ := term.ReadPassword(int(syscall.Stdin))

		cred := &credentials{
			token: string(accessTokenByte),
			email: userEmail,
		}

		pathToHome, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("could not find current user home directory: %v", err)
		}

		credentialsFile := fmt.Sprintf(cheesyCredsFileName + "." + cheesyCredsFileType)
		confFolderpath := fmt.Sprintf(pathToHome + "/" + cheesyCredsFolderName)

		if !helpers.IsExists(confFolderpath) {
			err = helpers.CreateDir(pathToHome, cheesyCredsFolderName)
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
		err = helpers.WriteTofile(confFolderpath+"/"+credentialsFile, file, 0755)
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
