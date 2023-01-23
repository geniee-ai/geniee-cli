package command

import (
	"bytes"
	"encoding/json"
	"fmt"

	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/geniee-ai/geniee-cli/internal/config"
	"github.com/geniee-ai/geniee-cli/internal/rgb"
	"github.com/urfave/cli/v2"
)

const (
	GenieeAPI = "http://localhost:9090"
	Endpoint  = "/ask"
)

type Request struct {
	Question string `json:"question"`
	UID      string `json:"uid"`
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

	// fmt.Println(len(cCtx.NArg().Get))

	var (
		question string
	)

	if cCtx.NArg() > 0 {
		question = cCtx.Args().Get(0)
	} else {
		rgb.White.Print(askBanner)
		os.Exit(1)

	}

	uid := config.Cfg.UID

	fmt.Print("\n")
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond) // Build our new spinner
	s.Suffix = " We are processing your request. Please wait..."
	s.Color("red", "bold")
	s.Start()                   // Start the spinner
	time.Sleep(4 * time.Second) // Run for some time to simulate work

	go callAPI(question, uid)
	s.Stop()

	// colors.White
	rgb.White.Print(success)
	rgb.Cyan.Print(note)
	fmt.Print("\n```")
	rgb.Green.Print(<-c1)
	fmt.Println("\n\n```")

	return nil
}

func callAPI(question, uid string) {

	request := &Request{
		Question: question,
		UID:      uid,
	}

	jsonByte, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", GenieeAPI+Endpoint, bytes.NewBuffer(jsonByte))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	response := HTTPResponse{}

	err = json.Unmarshal(body, &response)
	// Check your errors!
	if err != nil {
		log.Fatal(err.Error())
	}

	data := response.Result["data"].(string)
	c1 <- data

}
