package controller

import (
	"log"
	"strings"

	"github.com/jpradass/PasswordGenerator/wordlist"

	"github.com/jpradass/PasswordGenerator/iohandler"
)

// WordlistCtrl is a struct to maintain control on
// wordlists
type WordlistCtrl struct {
	logger   *log.Logger
	io       *iohandler.WordlistIO
	wordlist *[]wordlist.FileJSON
}

// New creates a new instance of Wordlist controller
func New(l *log.Logger) *WordlistCtrl {
	l.Println("Instantiating a new wordlist controller")
	return &WordlistCtrl{
		logger: l,
		io:     iohandler.New(l),
	}
}

// GetWordlist is a function that gets the wordlist
func (wctrl *WordlistCtrl) GetWordlist() error {
	wctrl.logger.Println("Getting wordlist")
	var err error

	if !wctrl.io.CheckWordlistFile() {
		file, err := wctrl.io.DownloadWordlist()
		if err != nil {
			wctrl.logger.Printf("There was an error downloading the wordlist: %s\n", err.Error())
			return err
		}

		wctrl.wordlist, err = parseWordlist(file)
		if err != nil {
			wctrl.logger.Printf("There was an error parsing the wordlist: %s\n", err.Error())
			return err
		}

		if err = wctrl.SaveWordlist(); err != nil {
			wctrl.logger.Printf("There was an error saving the wordlist: %s\n", err.Error())
			return err
		}

	} else {
		wctrl.wordlist, err = wctrl.io.GetWordlist()
		if err != nil {
			wctrl.logger.Printf("There was an error getting the wordlist: %s\n", err.Error())
			return err
		}
	}

	return nil
}

// SaveWordlist is a function that persists the wordlist
func (wctrl *WordlistCtrl) SaveWordlist() error {
	wctrl.logger.Println("Saving wordlist")

	err := wctrl.io.SaveWordlist(wctrl.wordlist)
	if err != nil {
		wctrl.logger.Printf("There was an error saving the wordlist: %s\n", err.Error())
		return err
	}

	return nil
}

func parseWordlist(wl []byte) (*[]wordlist.FileJSON, error) {
	wlstruct := []wordlist.FileJSON{}
	lines := strings.Split(string(wl), "\n")

	for key, value := range lines {
		wlstruct = append(wlstruct, wordlist.FileJSON{
			Key:   10000 + key,
			Value: value,
		})
	}
	return &wlstruct, nil
}
