package apps

import (
	"embed"
	"fmt"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"github.com/ahmedsat/quran/win"
)

var keyEventHandlers map[fyne.KeyName]func(q *Quran) error = map[fyne.KeyName]func(q *Quran) error{
	fyne.KeyEscape: quit,

	fyne.KeyL:        next,
	fyne.KeyRight:    next,
	fyne.KeyH:        previous,
	fyne.KeyLeft:     previous,
	fyne.KeyK:        up,
	fyne.KeyUp:       up,
	fyne.KeyPageUp:   up,
	fyne.KeyJ:        down,
	fyne.KeyDown:     down,
	fyne.KeyPageDown: down,
}

type Quran struct {
	Images     embed.FS
	App        *fyne.App
	master     fyne.Window
	imageBytes []byte
	imageCount int
	DarkTheme  bool
}

func (q *Quran) Run() (err error) {

	q.imageCount = 1

	title := "Quran"
	q.master = win.Base(*q.App, title)

	q.master.SetFixedSize(false)
	q.master.Canvas().SetOnTypedKey(func(ke *fyne.KeyEvent) {
		if _, ok := keyEventHandlers[ke.Name]; ok {
			err = keyEventHandlers[ke.Name](q)
		}
	})
	if err != nil {
		return
	}

	q.master.SetMaster()

	err = q.refresh()
	if err != nil {
		return
	}

	q.master.Show()

	return
}

func quit(q *Quran) (err error) {
	os.Exit(0)

	return
}

func (q *Quran) refresh() (err error) {

	if q.DarkTheme {
		q.imageBytes, err = q.Images.ReadFile(fmt.Sprintf("quran-images/dark/%v.png", q.imageCount))
	} else {
		q.imageBytes, err = q.Images.ReadFile(fmt.Sprintf("quran-images/%v.png", q.imageCount))
	}

	if err != nil {
		return
	}

	image := canvas.NewImageFromResource(fyne.NewStaticResource("sora", q.imageBytes))
	image.FillMode = canvas.ImageFillContain

	q.master.SetContent(image)

	return
}

func next(q *Quran) (err error) {

	if q.imageCount >= 604 {
		q.imageCount = 1
	} else {
		q.imageCount++
	}

	err = q.refresh()
	if err != nil {
		return
	}

	return
}

func previous(q *Quran) (err error) {

	if q.imageCount <= 1 {
		q.imageCount = 604
	} else {
		q.imageCount--
	}

	err = q.refresh()
	if err != nil {
		return
	}

	return
}

func up(q *Quran) (err error) {
	log.Println("up")

	return
}

func down(q *Quran) (err error) {
	log.Println("down")

	return
}
