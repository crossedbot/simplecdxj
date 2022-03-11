package simplecdxj

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIndexerLen(t *testing.T) {
	in := indexer{
		[]byte{0x01},
		[]byte{0x02},
		[]byte{0x03},
	}
	require.Equal(t, 3, in.Len())
}

func TestIndexerLess(t *testing.T) {
	in := indexer{
		[]byte{0x01},
		[]byte{0x02},
		[]byte{0x03},
	}
	isLess := in.Less(1, 2)
	require.True(t, isLess)
	isGreater := !in.Less(2, 1)
	require.True(t, isGreater)
}

func TestIndexerSwap(t *testing.T) {
	in := indexer{
		[]byte{0x01},
		[]byte{0x02},
		[]byte{0x03},
	}
	in.Swap(1, 2)
	isGreater := !in.Less(1, 2)
	require.True(t, isGreater)
}

func TestIndex(t *testing.T) {
	v := []byte("d\na\nc\nb\n")
	indexed, err := Index(bytes.NewReader(v))
	require.Nil(t, err)
	expected := []byte("a\nb\nc\nd\n")
	actual, err := ioutil.ReadAll(indexed)
	require.Nil(t, err)
	require.Equal(t, expected, actual)
}
