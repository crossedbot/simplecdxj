package simplecdxj

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

// CDXJ represents a CDXJ formatted file
type CDXJ struct {
	Header  Header
	Records []*Record
}

// Format returns the CDXJ as formatted string
func (cdxj *CDXJ) Format() string {
	hdr := cdxj.Header.Format()
	records := ""
	for _, r := range cdxj.Records {
		records = fmt.Sprintf("%s%s\n", records, r.Format())
	}
	indexed, _ := Index(bytes.NewReader([]byte(records)))
	b, _ := ioutil.ReadAll(indexed)
	records = string(b)
	return fmt.Sprintf("%s%s", hdr, records)
}

// Reader is an interface to a CDXJ reader
type Reader interface {
	// Read reads in, parses, and returns the CDXJ
	Read() (*CDXJ, error)

	// ReadLine returns the next line of CDXJ
	ReadLine() (string, error)
}

// reader implements a CDXJ reader
type reader struct {
	Reader *bufio.Reader
}

// NewReader wraps the given reader and returns a new CDXJ reader
func NewReader(r io.Reader) (Reader, error) {
	indexed, err := Index(r)
	if err != nil {
		return nil, err
	}
	return &reader{
		Reader: bufio.NewReader(indexed),
	}, nil
}

// Read returns a CDXJ
func (r *reader) Read() (*CDXJ, error) {
	cdxj := new(CDXJ)
	cdxj.Header = make(Header)
	line, err := r.ReadLine()
	if err != nil {
		return nil, err
	}
	for line != "" {
		if strings.HasPrefix(line, "!") {
			if key, val := parseHeader(line); key != "" {
				key = strings.ToLower(key)
				cdxj.Header.Set(key, val)
			}
		} else {
			if record, err := parseRecord(line); err == nil {
				cdxj.Records = append(cdxj.Records, record)
			}
		}
		line, err = r.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
	}
	return cdxj, nil
}

// ReadLine returns the next line of the CDXJ
func (r *reader) ReadLine() (string, error) {
	line, isPrefix, err := r.Reader.ReadLine()
	if err != nil {
		return "", err
	}
	str := string(line)
	if isPrefix {
		buffer := bytes.NewBuffer(line)
		for isPrefix {
			line, isPrefix, err = r.Reader.ReadLine()
			if err != nil {
				return "", err
			}
			buffer.Write(line)
		}
		str = buffer.String()
	}
	return str, nil
}
