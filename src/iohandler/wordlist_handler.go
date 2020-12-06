package iohandler

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

// NewWordlistIO returns a new wordlistHTTP handler with the given logger
func NewWordlistIO(logger *log.Logger) *WordlistIO {
	logger.Println("Instantiating WordlistIO")
	return &WordlistIO{logger}
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

// GetWordlist returns []byte to the file
func (wio *WordlistIO) GetWordlist() ([]byte, error) {
	wio.logger.Println("Reading wordlist locally")
	return ioutil.ReadFile(wordlistPATH)
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
