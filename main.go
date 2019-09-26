package main

import (
	"errors"

	"github.com/dbriemann/breeze/winctrl"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.SetConfigFlags(rl.FlagWindowUndecorated | rl.FlagWindowTransparent)
	rl.InitWindow(1200, 800, "breeze - the window organizer")

	rl.SetTargetFPS(60)

	winController := &winctrl.WMCtrl{}
	windows, err := winController.ListWindows()
	if errors.Is(err, winctrl.ErrParseFailed) {
		panic(err)
	} else if err != nil {
		panic(err)
	}

	// desktops, err := winController.ListDesktops()
	// if errors.Is(err, winctrl.ErrParseFailed) {
	// 	panic(err)
	// } else if err != nil {
	// 	panic(err)
	// }

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.NewColor(0, 0, 0, 128))

		for i, win := range windows {
			// Skip sticky windows, such as desktop or panels.
			if win.Desktop < 0 {
				i--
				continue
			}
			rl.DrawText(win.Name, 100, 36*int32(i)+100, 24, rl.White)
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
