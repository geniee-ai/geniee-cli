package token

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
)

func ValidateToken(accessToken string) bool {
	fmt.Print("\n\n")
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond) // Build our new spinner
	s.Suffix = " Validating access token"
	s.Color("red", "bold")
	s.Start()                   // Start the spinner
	time.Sleep(4 * time.Second) // Run for some time to simulate work
	s.Stop()
	return true
}

func Initialising() bool {
	fmt.Print("\n")
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond) // Build our new spinner
	s.Suffix = " Archiving repository"
	s.Color("red", "bold")
	s.Start()                   // Start the spinner
	time.Sleep(4 * time.Second) // Run for some time to simulate work
	s.Stop()
	return true
}
