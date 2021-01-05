package simplecdxj

import (
	"fmt"
	"strings"
	"time"
)

// RecordType represents the type of CDXJ record
type RecordType int

const (
	ResponseRecordType RecordType = iota + 1
	RequestRecordType
	RevisitRecordType
)

var RecordTypeStrings = []string{
	"response",
	"request",
	"revisit",
}

// String returns the string representation of the RecordType; if a string can
// not be found, an empty string is returned instead
func (rt RecordType) String() string {
	index := int(rt) - 1
	if index >= len(RecordTypeStrings) {
		return ""
	}
	return RecordTypeStrings[index]
}

// ParseRecordType parses the given string value and returns its RecordType
func ParseRecordType(v string) (RecordType, error) {
	for i, s := range RecordTypeStrings {
		if strings.EqualFold(v, s) {
			return RecordType(i + 1), nil
		}
	}
	return -1, fmt.Errorf("unknown record type string: %s", v)
}

// Record represents an CDXJ record
type Record struct {
	SURT      string     // SURT formatted domain name
	Timestamp time.Time  // Record timestamp
	Type      RecordType // Record type
	Content   []byte     // Record contents
}

// Format returns the string representation of the record
func (r *Record) Format() string {
	return fmt.Sprintf(
		"%s %s %s %s",
		r.SURT,
		r.Timestamp.Format(time.RFC3339),
		r.Type.String(),
		string(r.Content),
	)
}

// parseRecord parses the given string and returns it as a Record
func parseRecord(line string) (*Record, error) {
	parts := strings.Split(line, " ")
	if len(parts) < 4 {
		return nil, fmt.Errorf("invalid number of fields")
	}
	var err error
	record := new(Record)
	record.SURT = parts[0]
	record.Timestamp, err = time.Parse(time.RFC3339, parts[1])
	if err != nil {
		return nil, err
	}
	record.Type, err = ParseRecordType(parts[2])
	if err != nil {
		return nil, err
	}
	record.Content = []byte(strings.Join(parts[3:], " "))
	return record, nil
}
