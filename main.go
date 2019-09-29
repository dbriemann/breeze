package main

import (
	"github.com/dbriemann/breeze/winctrl"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/kr/pretty"
)

const (
	breezeWindowName = "breeze - the window organizer"
)

func main() {
	winController := &winctrl.WMCtrl{}
	desktops, err := winController.ListDesktops()
	if err != nil {
		panic(err)
	}
	if len(desktops) <= 0 {
		panic("no desktops found")
	}

	// Find active desktop and save work area dimensions.
	var area winctrl.Area
	for _, d := range desktops {
		if d.Active {
			area = d.WorkArea
		}
	}

	rl.SetConfigFlags(rl.FlagWindowUndecorated | rl.FlagWindowTransparent)
	rl.InitWindow(area.Width*4/5, area.Height*4/5, breezeWindowName)

	rl.SetTargetFPS(60)

	windows, err := winController.ListWindows()
	if err != nil {
		panic(err)
	}

	// Find the breeze window and make it sticky.
	for i, w := range windows {
		if w.Name == breezeWindowName {
			err := winController.SetWindowProps(&windows[i], winctrl.ActionAdd, winctrl.PropertyAbove, winctrl.PropertySticky)
			windows[i].Desktop = -1 // Update manually.
			if err != nil {
				panic(err)
			}
		}
	}

	pretty.Println(windows)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.NewColor(0, 0, 0, 128))

		for i, win := range windows {
			// Skip sticky windows, such as desktop or panels.
			if win.Desktop < 0 {
				i--
				continue
			}

			// rl.DrawText(win.Name, 100, 36*int32(i)+100, 24, rl.White)
			rl.DrawRectangleLines(win.XOffset*4/5, win.YOffset*4/5, win.Width*4/5, win.Height*4/5, rl.Red)
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
