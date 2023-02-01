package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"io/ioutil"
	"net/http"

	"github.com/briandowns/spinner"
	"github.com/geniee-ai/geniee-cli/internal/config"
	"github.com/geniee-ai/geniee-cli/internal/rgb"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

const (
	GenieeAPI   = "https://api.geniee.io"
	AskEndpoint = "/ask"
)

type Request struct {
	Question string `json:"question"`
	Email    string `json:"email"`
}

type Result map[string]interface{}

type HTTPResponse struct {
	Error      string `json:"error"`
	StatusCode int    `json:"status_code"`
	Result     Result `json:"result"`
}

func NewHTTPResponse() *HTTPResponse {
	return &HTTPResponse{}
}

var c1 = make(chan string)

var (
	askBanner = fmt.Sprint(`
Ask requires at least one argument. May be try "geniee ask --help"

For example:

	$ geniee ask "How can i find the age of the universe?"
`)

	success = fmt.Sprint(`
You request has been processed.

Here is your response.
`)

	note = fmt.Sprint(`
Note: It might take a while to parse response in the terminal. Please have patience.

`)
)

func AskCmd(cCtx *cli.Context) error {

	var (
		question string
	)

	if cCtx.NArg() > 0 {
		question = cCtx.Args().Get(0)
	} else {
		rgb.White.Print(askBanner)
		os.Exit(1)

	}

	cfg, err := config.LoadConfig()
	if err != nil {
		rgb.Red.Println("\nCould not load config file. Please validate config.json file in ~/.geniee/config.json")
		fmt.Println("")
		rgb.White.Println("If not present, May be try \"geniee login\" to generate one\n")
		log.Debug(err.Error())
		os.Exit(1)
	}

	// uid := cfg.UID
	email := cfg.Email
	token := cfg.Token

	fmt.Print("\n")
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond) // Build our new spinner
	s.Suffix = " We are processing your request. Please wait..."
	s.Color("red", "bold")
	s.Start()                   // Start the spinner
	time.Sleep(4 * time.Second) // Run for some time to simulate work

	go callAPI(question, email, token)
	s.Stop()

	// colors.White
	rgb.White.Print(success)
	rgb.Cyan.Print(note)
	fmt.Print("\n```\n")
	rgb.Green.Print(<-c1)
	fmt.Println("\n\n```")

	return nil
}

func callAPI(question, email, token string) {

	request := &Request{
		Question: question,
		Email:    email,
	}

	jsonByte, err := json.Marshal(request)
	if err != nil {
		fmt.Println("Could not process the request.")
		log.Debug(err.Error())
		os.Exit(1)
	}
	req, err := http.NewRequest("POST", GenieeAPI+AskEndpoint, bytes.NewBuffer(jsonByte))
	if err != nil {
		fmt.Println("Could not process the request.")
		log.Debug(err.Error())
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/json")

	queryParams := req.URL.Query()
	queryParams.Add("token", token)
	req.URL.RawQuery = queryParams.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Could not process the request. Please check your token and email in ~/.geniee/config.json")
		log.Debug(err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Could not process the request.")
		log.Debug(err.Error())
		os.Exit(1)
	}
	response := HTTPResponse{}

	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Could not process the request")
		log.Debug(err.Error())
		os.Exit(1)
	}

	switch response.StatusCode {
	case http.StatusUnauthorized:
		rgb.Red.Print("Invalid token or email\n")
		fmt.Println("")
		rgb.White.Print("Please verify your token and email value in ~/.geniee/config.json\n\n")
		os.Exit(1)
	case http.StatusOK:
		data := response.Result["data"].(string)
		c1 <- data
	default:
		rgb.White.Print("Could not process the request\n")
		log.Debug(response.Error)
		os.Exit(1)
	}

}
