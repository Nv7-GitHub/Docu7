package main

// The idea is to use bild's shear X & Y + having the user mark the 4 points to make an epic doc scanner

import (
	"image"
	"os"

	"image/color"
	_ "image/jpeg"
	_ "image/png"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
)

var win fyne.Window

func handle(err error) {
	if err != nil {
		dialog.ShowError(err, win)
		panic(err)
	}
}

func main() {
	app := app.New()
	win = app.NewWindow("Docu7")

	// Load file from file dialog instead
	file, err := os.Open("input.png")
	handle(err)
	img, _, err := image.Decode(file)
	handle(err)

	image := canvas.NewImageFromImage(img)
	image.FillMode = canvas.ImageFillContain

	circle := canvas.NewCircle(color.RGBA{R: 255, G: 0, B: 0, A: 255})

	box := container.NewMax(image, circle)
	win.SetContent(box)
	win.Resize(fyne.NewSize(850, 500))
	win.ShowAndRun()
}
