package iohandler

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/jpradass/PasswordGenerator/wordlist"
)

var (
	wordlistURL  string = "https://theworld.com/~reinhold/diceware.wordlist.asc"
	wordlistPATH string = "wordlist/diceware.wordlist.json"
)

// WordlistIO struct to control connection to handle wordlist resource
type WordlistIO struct {
	logger *log.Logger
}

// New returns a new wordlist IO handler
func New(l *log.Logger) *WordlistIO {
	l.Println("Instantiating a new wordlist IO handler")
	return &WordlistIO{logger: l}
}

// DownloadWordlist function to download wordlist locally
func (wio *WordlistIO) DownloadWordlist() ([]byte, error) {
	wio.logger.Println("Downloading wordlist remotely")

	client := getClient()
	res, err := client.Get(wordlistURL)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// GetWordlist returns the wordlist in JSON format
func (wio *WordlistIO) GetWordlist() (*[]wordlist.FileJSON, error) {
	wio.logger.Println("Reading wordlist locally")
	wl := []wordlist.FileJSON{}

	file, err := os.Open(wordlistPATH)
	if err != nil {
		wio.logger.Printf("There was an error opening the wordlist file: %s\n", err.Error())
		return nil, err
	}

	decd := json.NewDecoder(file)
	err = decd.Decode(&wl)
	if err != nil {
		wio.logger.Printf("There was an error decoding the file: %s\n", err.Error())
		return nil, err
	}

	return &wl, nil
}

// SaveWordlist handles the wordlist file locally
func (wio *WordlistIO) SaveWordlist(wl *[]wordlist.FileJSON) error {
	wio.logger.Println("Saving the wordlist locally")

	file, err := os.OpenFile(wordlistPATH, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		wio.logger.Printf("There was an error opening the wordlist file: %s\n", err.Error())
		return err
	}

	enc := json.NewEncoder(file)
	err = enc.Encode(&wl)
	if err != nil {
		wio.logger.Printf("There was an error encoding the file: %s\n", err.Error())
		return err
	}

	return nil
}

// CheckWordlistFile checks if wordlist file exists
func (wio *WordlistIO) CheckWordlistFile() bool {
	if _, err := os.Stat(wordlistPATH); os.IsNotExist(err) {
		return false
	}
	return true
}

func getClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
}
