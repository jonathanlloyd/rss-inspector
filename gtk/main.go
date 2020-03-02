package main

import (
	"os"

	"github.com/jonathanlloyd/rss-inspector/gtk/views"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

var APPLICATION_TITLE = "RSS Inspector"

func main() {
	gtk.Init(&os.Args)

	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetTitle(APPLICATION_TITLE)
	window.SetIconName("gtk-dialog-info")
	window.Connect("destroy", func(ctx *glib.CallbackContext) {
		gtk.MainQuit()
	}, "foo")
	window.SetSizeRequest(600, 900)

	mainScreen := views.NewMainScreen()
	mainScreen.Mount(window)

	window.ShowAll()
	gtk.Main()
}
