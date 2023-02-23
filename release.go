//go:build !debug
// +build !debug

package main

import "embed"

//go:embed quran-images quran-images/dark
var images embed.FS
