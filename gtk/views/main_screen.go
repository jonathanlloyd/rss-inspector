package views

import (
	"fmt"
	"unsafe"

	"github.com/jonathanlloyd/rss-inspector/entities"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

type MainScreenViewModel struct {
	Entries []entities.RssEntry
}

type MainScreen struct {
	vbox           *gtk.VBox
	loadFeedDialog *gtk.Dialog

	loadFeedDialogOpen bool
	loadFeedDialogURL  string
}

func NewMainScreen() MainScreen {
	containerBox := gtk.NewVBox(false, 1)
	innerBox := gtk.NewVBox(false, 1)

	label := gtk.NewLabel("No feed selected. Load a new feed to get started.")
	labelAlign := gtk.NewAlignment(0.5, 0.5, 0, 0)
	labelAlign.Add(label)
	innerBox.PackStart(labelAlign, false, false, 10)

	button := gtk.NewButtonWithLabel("Load Feed")
	buttonAlign := gtk.NewAlignment(0.5, 0, 0, 0)
	buttonAlign.Add(button)
	innerBox.PackStart(buttonAlign, false, false, 0)

	containerBox.PackStart(innerBox, true, false, 0)

	ms := MainScreen{
		vbox: containerBox,
	}

	button.Connect("clicked", func() {
		ms.loadFeedDialogOpen = true
		ms.Render(MainScreenViewModel{})
	})

	return ms
}

func (ms *MainScreen) showDialog() {
	if ms.loadFeedDialog == nil {
		dialog := gtk.NewDialog()
		dialog.SetSizeRequest(420, 70)
		dialog.SetTitle("Load Feed")

		dialogContainerBox := gtk.NewHBox(false, 5)
		dialogAlign := gtk.NewAlignment(0.5, 0.5, 0, 0)
		dialogAlign.Add(dialogContainerBox)

		dialog.GetVBox().Add(dialogAlign)

		entry := gtk.NewEntry()
		entry.SetWidthChars(35)
		entry.SetText("Enter feed URL")
		entry.Connect("insert-text", func(ctx *glib.CallbackContext) {
			a := (*[2000]uint8)(unsafe.Pointer(ctx.Args(0)))
			s := string(a[0])
			ms.loadFeedDialogURL = entry.GetText() + s
			ms.Render(MainScreenViewModel{})
		})
		dialogContainerBox.PackStart(entry, false, false, 0)

		dialogOKButton := gtk.NewButtonWithLabel("Load Feed")
		dialogContainerBox.PackStart(dialogOKButton, false, true, 0)

		dialog.Connect("destroy", func() {
			ms.loadFeedDialogOpen = false
			ms.Render(MainScreenViewModel{})
		})

		dialog.ShowAll()
		ms.loadFeedDialog = dialog
	}
}

func (ms *MainScreen) hideDialog() {
	if ms.loadFeedDialog != nil {
		ms.loadFeedDialog.Destroy()
		ms.loadFeedDialog = nil
	}
}

func (ms *MainScreen) Mount(w *gtk.Window) {
	w.Add(ms.vbox)
	ms.Render(MainScreenViewModel{})
}

func (ms *MainScreen) Unmount(w *gtk.Window) {
	w.Remove(ms.vbox)
}

func (ms *MainScreen) Render(vm MainScreenViewModel) {
	if ms.loadFeedDialogOpen {
		ms.showDialog()
	} else {
		ms.hideDialog()
	}

	fmt.Println(ms.loadFeedDialogURL)
}
