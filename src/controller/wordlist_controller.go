package controller

import (
	"log"

	"github.com/jpradass/PasswordGenerator/iohandler"
)

// WordlistCtrl is a struct to maintain control on
// wordlists
type WordlistCtrl struct {
	logger *log.Logger
}

// NewWordlist creates a new instance of Wordlist controller
func NewWordlist(l *log.Logger) *WordlistCtrl {
	l.Println("Instantiating a new Wordlist controller")
	return &WordlistCtrl{logger: l}
}

// GetWordlist is a function to get wordlist
func (wctrl *WordlistCtrl) GetWordlist() ([]byte, error) {
	wctrl.logger.Println("Getting wordlist")
	wio := iohandler.NewWordlistIO(wctrl.logger)

	if !wio.CheckWordlistFile() {
		return wio.DownloadWordlist()
	}

	return wio.GetWordlist()
}

// func persistWordlistFile(file []byte) error {

// }
