package simplecdxj

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWriterWrite(t *testing.T) {
	expected := []byte("!context [\"http://oduwsdl.github.io/contexts/arhiveprofiles\"]\n" +
		"!hermes-cdxj 1.0\n" +
		"!id {\"uri\": \"http://archive.org/\"}\n" +
		"!keys [\"surt_uri\", \"year\"]\n" +
		"!meta {\"name\": \"Internet Archive\", \"year\": 1996}\n" +
		"com,abc)/ 2020-01-02T03:04:05-01:00 response {'uri':'abc.com/hello', 'sha':'<SHA-1 digest>', 'ref':'warcfile:IAH-20070824123353-00393-abc.com.arc.gz#25523382'}\n" +
		"com,cnn)/ 2020-01-02T03:04:05-01:00 response {'uri':'cnn.com/world', 'sha':'<SHA-1 digest>', 'ref':'warcfile:IAH-20070824123353-00393-cnn.com.arc.gz#25523382'}\n")
	source := bytes.NewReader(expected)
	reader, err := NewReader(source)
	require.Nil(t, err)
	cdxj, err := reader.Read()
	require.Nil(t, err)
	var b bytes.Buffer
	bw := bufio.NewWriter(&b)
	writer := NewWriter(bw)
	n, err := writer.Write(cdxj)
	bw.Flush()
	require.Nil(t, err)
	require.Equal(t, len(expected), n)
	require.Equal(t, expected, b.Bytes())
}
