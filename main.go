package main

// The idea is to use bild's shear X & Y + having the user mark the 4 points to make an epic doc scanner

import (
	"image"
	"os"

	_ "image/jpeg"
	_ "image/png"

	r "github.com/Lachee/raylib-goplus/raylib"
)

const fps = 60

var (
	width   = 0
	height  = 0
	isHidpi = true

	div float32 = 1
)

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

func windowResize() {
	width, height = r.GetScreenWidth(), r.GetScreenHeight()
	if isHidpi {
		pos := r.GetWindowPosition()
		r.SetWindowPosition(int(pos.X+1), int(pos.Y+1))
		r.SetWindowPosition(int(pos.X), int(pos.Y))
	}
}

func main() {
	// Load file from file dialog instead
	file, err := os.Open("input.png")
	handle(err)
	img, _, err := image.Decode(file)
	handle(err)
	imw, imh := float32(img.Bounds().Dx()), float32(img.Bounds().Dy())

	// Raylib
	r.SetConfigFlags(r.FlagWindowResizable | r.FlagVsyncHint)

	r.InitWindow(img.Bounds().Dx(), img.Bounds().Dy(), "Docu7")
	r.SetTargetFPS(fps)
	defer r.UnloadAll()

	// Load image onto the GPU
	tex := r.LoadTextureFromGo(img)
	windowResize()

	// Game loop
	for !r.WindowShouldClose() {
		if r.IsWindowResized() {
			if isHidpi && div == 1 {
				div = 0.5
			}

			windowResize()
		}

		r.BeginDrawing()
		r.ClearBackground(r.RayWhite)

		sc := float32(height) / imh
		wsc := float32(width) / imw
		if wsc < sc {
			sc = wsc
		}
		r.DrawTextureEx(tex, r.NewVector2((float32(width)/2-sc*imw/2)*div, (float32(height)/2-sc*imh/2)*div), 0, div*sc, r.White)

		r.EndDrawing()
	}
}
