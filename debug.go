//go:build debug
// +build debug

package main

import "embed"

//go:embed quran-images/1.png  quran-images/2.png  quran-images/3.png  quran-images/4.png  quran-images/5.png
//go:embed quran-images/dark/1.png  quran-images/dark/2.png  quran-images/dark/3.png  quran-images/dark/4.png  quran-images/dark/5.png
var images embed.FS
