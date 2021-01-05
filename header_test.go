package simplecdxj

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	expected := []byte("World")
	hdr := Header{"Hello": expected}
	actual := hdr.Get("Hello")
	require.Equal(t, expected, actual)
}

func TestSet(t *testing.T) {
	hdr := make(Header)
	key := "Hello"
	expected := []byte("World")
	hdr.Set(key, expected)
	actual := hdr.Get(key)
	require.Equal(t, expected, actual)
}

func TestDelete(t *testing.T) {
	hdr := make(Header)
	key := "Hello"
	hdr.Set(key, []byte("World"))
	hdr.Delete(key)
	require.Nil(t, hdr.Get(key))
}

func TestHas(t *testing.T) {
	hdr := make(Header)
	key := "Hello"
	require.False(t, hdr.Has(key))
	hdr.Set(key, []byte("World"))
	require.True(t, hdr.Has(key))
}

func TestFormat(t *testing.T) {
	hdr := make(Header)
	hdr.Set("d", []byte("four"))
	hdr.Set("a", []byte("one"))
	hdr.Set("c", []byte("three"))
	hdr.Set("b", []byte("two"))
	expected := "!a one\n!b two\n!c three\n!d four\n"
	actual := hdr.Format()
	require.Equal(t, expected, actual)
}

func TestParseHeader(t *testing.T) {
	expectedKey := "hello"
	expectedVal := []byte("world")
	actualKey, actualVal := parseHeader("!hello world")
	require.Equal(t, expectedKey, actualKey)
	require.Equal(t, expectedVal, actualVal)
}
