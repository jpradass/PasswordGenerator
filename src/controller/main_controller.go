package controller

import "log"

// Controller is a struct to keep track of program flow
type Controller struct {
	logger  *log.Logger
	wctrler *WordlistCtrl
}

// New creates a new instance of main controller
func New(l *log.Logger) *Controller {
	l.Println("Instantiating a new main controller")
	return &Controller{logger: l}
}

// NewWordlistCtrller creates a new instace of wordlist controller
func (ctrler *Controller) NewWordlistCtrller() {
	if ctrler.wctrler == nil {
		ctrler.wctrler = NewWordlist(ctrler.logger)
	}
}
