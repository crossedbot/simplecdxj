package simplecdxj

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCDXJFormat(t *testing.T) {
	header := Header{
		"meta":        []byte(`{"name": "Internet Archive", "year": 1996}`),
		"Hermes-CDXJ": []byte(`1.0`),
		"id":          []byte(`{"uri": "http://archive.org/"}`),
		"context":     []byte(`["http://oduwsdl.github.io/contexts/arhiveprofiles"]`),
		"keys":        []byte(`["surt_uri", "year"]`),
	}
	ts, err := time.Parse(time.RFC3339, "2020-01-02T03:04:05-01:00")
	require.Nil(t, err)
	ty, err := ParseRecordType("response")
	require.Nil(t, err)
	records := []*Record{
		&Record{
			SURT:      "com,cnn)/",
			Timestamp: ts,
			Type:      ty,
			Content: []byte(`{` +
				`'uri':'cnn.com/world', ` +
				`'sha':'<SHA-1 digest>', ` +
				`'ref':'warcfile:IAH-20070824123353-00393-cnn.com.arc.gz#25523382'` +
				`}`),
		}, &Record{
			SURT:      "com,abc)/",
			Timestamp: ts,
			Type:      ty,
			Content: []byte(`{` +
				`'uri':'abc.com/hello', ` +
				`'sha':'<SHA-1 digest>', ` +
				`'ref':'warcfile:IAH-20070824123353-00393-abc.com.arc.gz#25523382'` +
				`}`),
		},
	}
	cdxj := CDXJ{
		Header:  header,
		Records: records,
	}
	expected := "!Hermes-CDXJ 1.0\n" +
		"!context [\"http://oduwsdl.github.io/contexts/arhiveprofiles\"]\n" +
		"!id {\"uri\": \"http://archive.org/\"}\n" +
		"!keys [\"surt_uri\", \"year\"]\n" +
		"!meta {\"name\": \"Internet Archive\", \"year\": 1996}\n" +
		"com,abc)/ 2020-01-02T03:04:05-01:00 response {'uri':'abc.com/hello', 'sha':'<SHA-1 digest>', 'ref':'warcfile:IAH-20070824123353-00393-abc.com.arc.gz#25523382'}\n" +
		"com,cnn)/ 2020-01-02T03:04:05-01:00 response {'uri':'cnn.com/world', 'sha':'<SHA-1 digest>', 'ref':'warcfile:IAH-20070824123353-00393-cnn.com.arc.gz#25523382'}\n"
	actual := cdxj.Format()
	require.Equal(t, expected, actual)
}

func TestReaderRead(t *testing.T) {
	expected := "!context [\"http://oduwsdl.github.io/contexts/arhiveprofiles\"]\n" +
		"!hermes-cdxj 1.0\n" +
		"!id {\"uri\": \"http://archive.org/\"}\n" +
		"!keys [\"surt_uri\", \"year\"]\n" +
		"!meta {\"name\": \"Internet Archive\", \"year\": 1996}\n" +
		"com,abc)/ 2020-01-02T03:04:05-01:00 response {'uri':'abc.com/hello', 'sha':'<SHA-1 digest>', 'ref':'warcfile:IAH-20070824123353-00393-abc.com.arc.gz#25523382'}\n" +
		"com,cnn)/ 2020-01-02T03:04:05-01:00 response {'uri':'cnn.com/world', 'sha':'<SHA-1 digest>', 'ref':'warcfile:IAH-20070824123353-00393-cnn.com.arc.gz#25523382'}\n"
	source := bytes.NewReader([]byte(expected))
	reader, err := NewReader(source)
	require.Nil(t, err)
	cdxj, err := reader.Read()
	require.Nil(t, err)
	require.Equal(t, expected, cdxj.Format())
}

func TestReaderReadLine(t *testing.T) {
	v := "!context [\"http://oduwsdl.github.io/contexts/arhiveprofiles\"]\n" +
		"!hermes-cdxj 1.0\n" +
		"!id {\"uri\": \"http://archive.org/\"}\n" +
		"!keys [\"surt_uri\", \"year\"]\n" +
		"!meta {\"name\": \"Internet Archive\", \"year\": 1996}\n" +
		"com,abc)/ 2020-01-02T03:04:05-01:00 response {'uri':'abc.com/hello', 'sha':'<SHA-1 digest>', 'ref':'warcfile:IAH-20070824123353-00393-abc.com.arc.gz#25523382'}\n" +
		"com,cnn)/ 2020-01-02T03:04:05-01:00 response {'uri':'cnn.com/world', 'sha':'<SHA-1 digest>', 'ref':'warcfile:IAH-20070824123353-00393-cnn.com.arc.gz#25523382'}\n"
	source := bytes.NewReader([]byte(v))
	reader, err := NewReader(source)
	require.Nil(t, err)
	expected := "!context [\"http://oduwsdl.github.io/contexts/arhiveprofiles\"]"
	actual, err := reader.ReadLine()
	require.Nil(t, err)
	require.Equal(t, expected, actual)
}
