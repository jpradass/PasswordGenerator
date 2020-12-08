package controller

import (
	"log"
	"strconv"
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
	l.Println("Instantiating a new Wordlist controller")
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
			return err
		}

		wctrl.wordlist, err = parseWordlist(file)
		if err != nil {
			return err
		}

	} else {
		wctrl.wordlist, err = wctrl.io.GetWordlist()
		if err != nil {
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
		return err
	}

	return nil
}

func parseWordlist(wl []byte) (*[]wordlist.FileJSON, error) {
	wlstruct := []wordlist.FileJSON{}
	lines := strings.Split(string(wl), "\n")

	for _, value := range lines[2:7778] {
		result := strings.Split(value, "\t")
		key, _ := strconv.Atoi(result[0])
		wlstruct = append(wlstruct, wordlist.FileJSON{
			Key:   key,
			Value: result[1],
		})
	}
	return &wlstruct, nil
}
