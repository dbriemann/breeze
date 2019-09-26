// Package winctrl provides a general interface for controlling
// a window manager / its windows. It also provides a basic
// implementation that invokes 'wmctrl' command and parses its
// output.
package winctrl

type Area struct {
	X      int32
	Y      int32
	Width  int32
	Height int32
}

type Window struct {
	ID      uint32
	Desktop int32
	PID     uint32
	XOffset int32
	YOffset int32
	Width   int32
	Height  int32
	Host    string
	Name    string
}

type Screen struct {
}

type Desktop struct {
	Num      uint32
	Active   bool
	DeskArea Area
	WorkArea Area
	Name     string
}

// Controller bundles all functions needed by breeze
// for the manipulation of windows. Future implementations
// of this interface could talk to the X server directly
// or use xlib/libxcb.
type Controller interface {
	// ListWindows returns a list of all windows on all screens.
	ListWindows() ([]Window, error)
	// ShowWindow switches to the desktop containing the window,
	// raises the window, and gives it focus.
	ShowWindow(w *Window) error
	// ListDesktops returns a list of all desktops.
	ListDesktops() ([]Desktop, error)
}
