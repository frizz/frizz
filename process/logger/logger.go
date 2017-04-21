package logger // import "frizz.io/process/logger"

// frizz: {"package": {"complete": true}}

import (
	"bytes"
	"io"
	"os"
)

func Logger(log bool) (combined *bytes.Buffer, stdout io.Writer, stderr io.Writer) {
	combined = &bytes.Buffer{}
	if log {
		stderr = MultiWriter(os.Stderr, combined)
		stdout = MultiWriter(os.Stdout, combined)
	} else {
		stderr = combined
		stdout = combined
	}
	return
}

// MultiWriter creates a writer that duplicates its writes to all the
// provided writers, similar to the Unix tee(1) command.
func MultiWriter(primary io.Writer, writers ...io.Writer) io.Writer {
	w := make([]io.Writer, len(writers))
	copy(w, writers)
	return &multiWriter{primary: primary, writers: w}
}

type multiWriter struct {
	primary io.Writer
	writers []io.Writer
}

func (t *multiWriter) Write(p []byte) (n int, err error) {
	for _, w := range t.writers {
		w.Write(p)
	}
	return t.primary.Write(p)
}
