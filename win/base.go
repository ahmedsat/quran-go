package win

import (
	"fyne.io/fyne/v2"
)

func Base(app fyne.App, title string) (w fyne.Window) {
	w = app.NewWindow(title)
	w.SetFixedSize(true)
	return
}
