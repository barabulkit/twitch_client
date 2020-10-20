package ui

import (
	"github.com/goki/gi/gi"
)

func RenderUi() {
	width := 1024
	height := 768

	win := gi.NewMainWindow("gogi-basic", "Twitch client", width, height)
	win.SetCloseCleanFunc(func(w *gi.Window) {
		go gi.Quit()
	})
	win.StartEventLoop()
}
