package gui

import (
	"log"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
)

// DrawWindow creates a windows, experimental stuff...
func DrawWindow() {
	driver.Main(func(s screen.Screen) {
		w, err := s.NewWindow(&screen.NewWindowOptions{
			Title: "Wrap",
		})
		if err != nil {
			log.Fatal(err)
		}

		defer w.Release()

		w.Publish()

		for {
			// Force the window to stay
		}
	})
}
