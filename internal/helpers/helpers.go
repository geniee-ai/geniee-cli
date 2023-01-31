package helpers

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/briandowns/spinner"
)

func CreateDir(path string, folderName string) error {
	folderPath := fmt.Sprintf(path + "/" + folderName)
	err := os.Mkdir(folderPath, os.ModePerm)
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			return nil
		} else {
			return fmt.Errorf("could not create directory named %v at %v: %v", folderName, path, err)
		}
	}
	return nil
}

func CreateFile(path string, fileName string) error {

	f, err := os.Create(path + "/" + fileName)
	if err != nil {
		return fmt.Errorf("could not create file named %v at %v: %v", fileName, path, err)
	}
	defer f.Close()
	return nil
}

func WriteTofile(path string, data []byte, perm fs.FileMode) error {
	err := ioutil.WriteFile("test.json", data, perm)
	if err != nil {
		return fmt.Errorf("could not write to file: %v", err)
	}
	return nil
}

func ReadInputString() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	inputString := strings.TrimSuffix(input, "\n")
	return inputString
}

func Spinner() {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond) // Build our new spinner
	s.Suffix = " Archiving repository"
	s.Color("red", "bold")
	s.Start()                   // Start the spinner
	time.Sleep(4 * time.Second) // Run for some time to simulate work
	s.Stop()
}

func IsExists(path string) bool {
	_, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}
