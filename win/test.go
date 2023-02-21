package win

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func Test(app fyne.App) (w fyne.Window) {

	w = Base(app, "test window")
	w.Resize(fyne.NewSize(800, 600))

	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	widgets := make([]fyne.CanvasObject, 0)

	for i := 0; i < 15; i++ {
		widgets = append(widgets, canvas.NewText("Hello", green))
	}

	// content := container.NewWithoutLayout(text1, text2)
	content := container.New(layout.NewGridLayout(5), widgets...)
	w.SetContent(content)

	return
}
