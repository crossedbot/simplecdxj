package simplecdxj

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	HEADER_PREFIX = "!"
)

// Header maps the CDXJ field-name and field-value
type Header map[string][]byte

// Get returns the field-value for a given field-name
func (h Header) Get(key string) []byte {
	if h == nil {
		return nil
	}
	return h[key]
}

// Set sets the field-value for a given field-name
func (h Header) Set(key string, val []byte) {
	h[key] = val
}

// Delete removes the field-name
func (h Header) Delete(key string) {
	delete(h, key)
}

// Has retuns true if the key exists
func (h Header) Has(key string) bool {
	_, ok := h[key]
	return ok
}

// Format returns the Header as a formatted string
func (h Header) Format() string {
	f := ""
	if h != nil {
		for k, v := range h {
			f = fmt.Sprintf("%s%s%s %s\n", f, HEADER_PREFIX, k, v)
		}
		indexed, _ := Index(bytes.NewReader([]byte(f)))
		b, _ := ioutil.ReadAll(indexed)
		f = string(b)
	}
	return f
}

// parseHeader parses the given line and returns the field-name and field-value
func parseHeader(line string) (key string, val []byte) {
	line = strings.TrimPrefix(line, HEADER_PREFIX)
	parts := strings.Split(line, " ")
	key = strings.TrimSpace(parts[0])
	if len(parts) > 1 {
		val = []byte(
			strings.TrimSpace(
				strings.Join(parts[1:], " "),
			),
		)
	}
	return
}
