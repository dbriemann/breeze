package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func run() {
	m := pixelgl.PrimaryMonitor()
	mw, mh := m.Size()
	w, h := mw/2, mh/2

	cfg := pixelgl.WindowConfig{
		Title:       "breeze - the window organizer",
		Bounds:      pixel.R(0, 0, w, h),
		Resizable:   false,
		VSync:       true,
		Undecorated: true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.SetSmooth(false)
	win.SetPos(pixel.V((mw-w)/2, (mh-h)/2))

	for !win.Closed() && !win.Pressed(pixelgl.KeyEscape) {
		win.Clear(pixel.Alpha(0.8))
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
