package winctrl

import (
	"io"
	"log"
	"os"
)

var (
	trace *log.Logger
)

func init() {
	trace = log.New(os.Stderr, "[DEBUG]: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// InitTrace allows to provide another io.Writer for logging purpose.
// Default is stderr.
func InitTrace(w io.Writer) {
	trace = log.New(w, "[DEBUG]: ", log.Ldate|log.Ltime|log.Lshortfile)
}
