package token

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/geniee-ai/geniee-cli/internal/model"
	"github.com/geniee-ai/geniee-cli/internal/rgb"
)

func ValidateToken(email, token string) bool {

	creds := model.Credentials{
		Email: email,
		Token: token,
	}
	fmt.Print("\n\n")
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond) // Build our new spinner
	s.Suffix = " Validating access token"
	s.Color("red", "bold")
	s.Start()                   // Start the spinner
	time.Sleep(4 * time.Second) // Run for some time to simulate work
	s.Stop()

	URL := "http://localhost:9090/login"
	jsonByte, err := json.Marshal(creds)
	if err != nil {
		fmt.Errorf("Could not marshal config data")
		os.Exit(1)
	}
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonByte))
	if err != nil {
		fmt.Errorf("Could not perform HTTP request")
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Errorf("Could not perform HTTP request")
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("Could not process the request. Please try again after some time.")
		os.Exit(1)
	}
	response := model.HTTPResponse{}

	err = json.Unmarshal(body, &response)
	if err != nil {
		rgb.Red.Print("Could not process the request. Please try again after some time.")
		os.Exit(1)
	}

	if response.StatusCode == http.StatusUnauthorized {
		rgb.Red.Print("Invalid token or email\n")
		fmt.Println("")
		rgb.White.Print("Please verify your token and email value in ~/.geniee/config.json\n")
		os.Exit(1)
	}

	return true
}
