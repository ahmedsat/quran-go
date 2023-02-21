package utils

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

var (
	Esc  = &desktop.CustomShortcut{KeyName: fyne.Key0, Modifier: fyne.KeyModifierControl}
	Next = &desktop.CustomShortcut{KeyName: fyne.KeyN, Modifier: 0}
)
