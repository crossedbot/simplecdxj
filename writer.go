package simplecdxj

import (
	"io"
)

// Writer is an interface for writing CDXJs
type Writer interface {
	// Write writes the CDXJ using the wrapped writer
	Write(cdxj *CDXJ) (int, error)
}

// writer implements a writer for CDJXs
type writer struct {
	Writer io.Writer
}

// NewWriter wraps the given writer and returns a CDXJ writer
func NewWriter(w io.Writer) Writer {
	return &writer{Writer: w}
}

// Writer writes the CDXJ
func (w *writer) Write(cdxj *CDXJ) (int, error) {
	b := []byte(cdxj.Format())
	return w.Writer.Write(b)
}
