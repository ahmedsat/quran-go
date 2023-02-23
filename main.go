package main

import (
	"os"

	"fyne.io/fyne/v2/app"
	"github.com/ahmedsat/quran/apps"
)

// todo : make the app remember last page
// todo : search for better images
// todo : add GUI controls
// todo// : fix dark theme
// todo : add jump to page future
// todo : add search ability

// todo : make the app configurable

func main() {

	a := app.NewWithID("com.ahmedsat.quran")

	// a.Settings().SetTheme(theme.LightTheme()) // ! deprecated: to be removed

	// darkTheme := true

	quran := apps.Quran{
		DarkTheme: true,
		App:       &a,
		Images:    images,
	}

	err := quran.Run()
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

}
