package gui

import (
	"log"
	"os"

	"golang.org/x/mobile/event/lifecycle"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
)

func manageGui(scrn screen.Screen) {
	wndw, err := scrn.NewWindow(&screen.NewWindowOptions{
		Title:  "Wrap",
		Width:  800,
		Height: 500,
	})

	if err != nil {
		log.Fatal(err)
	}

	defer wndw.Release()

	wndw.Publish()

	// Event loop
	for {
		evnt := wndw.NextEvent()

		switch evnt := evnt.(type) {
		case lifecycle.Event:
			// User pressed close button
			if evnt.To == lifecycle.StageDead {
				os.Exit(0)
			}
		}
	}
}

// DrawWindow starts the GUI
func DrawWindow() {
	driver.Main(manageGui)
}
