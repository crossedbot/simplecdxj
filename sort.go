package simplecdxj

import (
	"bytes"
	"io"
	"io/ioutil"
	"sort"
)

// indexer is a list of byte strings that implements the Sort interface
type indexer [][]byte

// Len returns the length of the indexer's list
func (in indexer) Len() int { return len(in) }

// Less returns true if the value at index 'i' is less than the value at index
// 'j'
func (in indexer) Less(i, j int) bool { return bytes.Compare(in[i], in[j]) < 0 }

// Swap swaps the values at index 'i' and 'j'
func (in indexer) Swap(i, j int) { in[i], in[j] = in[j], in[i] }

// Index indexes (sorts) the CDXJ read from reader
func Index(r io.Reader) (io.Reader, error) {
	c, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	c = bytes.Trim(c, "\n")
	l := bytes.Split(c, []byte{0x0a})
	sort.Sort(indexer(l))
	c = bytes.Join(l, []byte{0x0a})
	return bytes.NewReader(append(c, 0x0a)), nil
}
