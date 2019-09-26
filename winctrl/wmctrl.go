package winctrl

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var (
	ErrParseFailed = fmt.Errorf("error during parsing")
)

// WMCtrl implements the Controller interface by invoking the
// command 'wmctrl' to retrieve and set window data.
type WMCtrl struct {
}

func (c *WMCtrl) ListDesktops() ([]Desktop, error) {
	ds := []Desktop{}
	resp, err := execWMCtrl("wmctrl", "-d")
	if err != nil {
		return ds, err
	}

	lines := parseWMCtrlOutput(resp)
	for _, line := range lines {
		d := Desktop{}

		id, err := strconv.ParseUint(line[0], 10, 32)
		if err != nil {
			return ds, fmt.Errorf("%w (Desktop.Num): %s", ErrParseFailed, err)
		}
		d.Num = uint32(id)
		line = line[1:]

		if line[0] == "*" {
			d.Active = true
		}
		line = line[1:]

		line = line[1:] // skip "DG:"
		if _, err = fmt.Sscanf(line[0], "%dx%d", &d.DeskArea.Width, &d.DeskArea.Height); err != nil {
			return ds, fmt.Errorf("%w (Desktop.DeskArea.Size): %s", ErrParseFailed, err)
		}
		line = line[1:]

		line = line[1:] // skip "VP:"
		if _, err = fmt.Sscanf(line[0], "%d,%d", &d.DeskArea.X, &d.DeskArea.Y); err != nil {
			return ds, fmt.Errorf("%w (Desktop.DeskArea.Pos): %s", ErrParseFailed, err)
		}
		line = line[1:]

		line = line[1:] // skip "WA:"
		if _, err = fmt.Sscanf(line[0], "%d,%d", &d.WorkArea.X, &d.WorkArea.Y); err != nil {
			return ds, fmt.Errorf("%w (Desktop.WorkArea.Pos): %s", ErrParseFailed, err)
		}
		line = line[1:]
		if _, err = fmt.Sscanf(line[0], "%dx%d", &d.WorkArea.Width, &d.WorkArea.Height); err != nil {
			return ds, fmt.Errorf("%w (Desktop.WorkArea.Size): %s", ErrParseFailed, err)
		}
		line = line[1:]

		d.Name = strings.Join(line, " ")

		ds = append(ds, d)
	}

	fmt.Println("DS", ds)

	return ds, nil
}

func (c *WMCtrl) ShowWindow(w *Window) error {
	id := fmt.Sprintf("0x%x", w.ID)
	_, err := execWMCtrl("wmctrl", "-a", id, "-i")

	return err
}

func (c *WMCtrl) ListWindows() ([]Window, error) {
	ws := []Window{}
	resp, err := execWMCtrl("wmctrl", "-l", "-p", "-G")
	if err != nil {
		return ws, err
	}

	lines := parseWMCtrlOutput(resp)
	for _, line := range lines {
		w := Window{}

		// ID is hex as string.. strip away "0x".
		id, err := strconv.ParseUint(line[0][2:], 16, 32)
		if err != nil {
			return ws, fmt.Errorf("%w (Window.ID): %s", ErrParseFailed, err)
		}
		w.ID = uint32(id)
		line = line[1:]

		desktop, err := strconv.ParseInt(line[0], 10, 32)
		if err != nil {
			return ws, fmt.Errorf("%w (Window.Desktop): %s", ErrParseFailed, err)
		}
		w.Desktop = int32(desktop)
		line = line[1:]

		pid, err := strconv.ParseUint(line[0], 10, 32)
		if err != nil {
			return ws, fmt.Errorf("%w (Window.PID): %s", ErrParseFailed, err)
		}
		w.PID = uint32(pid)
		line = line[1:]

		xoff, err := strconv.ParseInt(line[0], 10, 32)
		if err != nil {
			return ws, fmt.Errorf("%w (Window.XOffset): %s", ErrParseFailed, err)
		}
		w.XOffset = int32(xoff)
		line = line[1:]

		yoff, err := strconv.ParseInt(line[0], 10, 32)
		if err != nil {
			return ws, fmt.Errorf("%w (Window.YOffset): %s", ErrParseFailed, err)
		}
		w.YOffset = int32(yoff)
		line = line[1:]

		width, err := strconv.ParseInt(line[0], 10, 32)
		if err != nil {
			return ws, fmt.Errorf("%w (Window.Width): %s", ErrParseFailed, err)
		}
		w.Width = int32(width)
		line = line[1:]

		height, err := strconv.ParseInt(line[0], 10, 32)
		if err != nil {
			return ws, fmt.Errorf("%w (Window.Height): %s", ErrParseFailed, err)
		}
		w.Height = int32(height)
		line = line[1:]

		w.Host = line[0]
		line = line[1:]

		// Rest of the fields is the complete window name and must be joined again.
		w.Name = strings.Join(line, " ")

		ws = append(ws, w)
	}

	return ws, nil
}

func parseWMCtrlOutput(output []byte) [][]string {
	reader := bytes.NewReader(output)
	bufReader := bufio.NewReader(reader)

	result := [][]string{}
	for {
		linestr, _, err := bufReader.ReadLine()
		if err == io.EOF {
			break
		}

		line := strings.Fields(string(linestr))
		result = append(result, line)
	}
	return result
}

func execWMCtrl(name string, args ...string) ([]byte, error) {
	cmd := exec.Command(name, args...)
	response, err := cmd.CombinedOutput()

	if err == exec.ErrNotFound {
		fmt.Println("'wmctrl' could not be found\nmake sure it is installed and resides inside path")
		os.Exit(1)
	}
	return response, err
}
