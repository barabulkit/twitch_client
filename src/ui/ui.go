package ui

import (
	"log"
	"github.com/gotk3/gotk3/gtk"
)

func Render() {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window", err)
	}

	win.SetTitle("Twitch client")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	win.SetDefaultSize(800, 600)
	win.ShowAll()

	gtk.Main()
}
