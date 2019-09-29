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

type windowAction string
type windowProperty string

const (
	ActionAdd    windowAction = "add"
	ActionToggle windowAction = "toggle"
	ActionRemove windowAction = "remove"

	PropertyModal       windowProperty = "modal"
	PropertySticky      windowProperty = "sticky"
	PropertyMaxVert     windowProperty = "maximized_vert"
	PropertyMaxHorz     windowProperty = "maximized_horz"
	PropertyShaded      windowProperty = "shaded"
	PropertySkipTaskbar windowProperty = "skip_taskbar"
	PropertySkipPager   windowProperty = "skip_pager"
	PropertyHidden      windowProperty = "hidden"
	PropertyFullscreen  windowProperty = "fullscreen"
	PropertyAbove       windowProperty = "above"
	PropertyBelow       windowProperty = "below"
	PropertyNone        windowProperty = ""
)

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
	// SetWindowProps allows to add, toggle or remove an arbitrary
	// amount of window properties.
	SetWindowProps(w *Window, action windowAction, props ...windowProperty) error
}
