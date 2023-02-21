package main

// todo : make the app remember last page
// todo : search for better images
// todo : add GUI controls
// todo : fix dark theme
// todo : add jump to page future
// todo : add search ability

// todo : make the app configurable

import (
	"embed"
	"os"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"github.com/ahmedsat/quran/apps"
)

//go:embed quran-images
var images embed.FS

func main() {

	a := app.New()

	a.Settings().SetTheme(theme.LightTheme()) // ! deprecated: to be removed

	quran := apps.Quran{
		App:    &a,
		Images: images,
	}

	err := quran.Run()
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}
