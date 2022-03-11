package simplecdxj

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRecordTypeString(t *testing.T) {
	expected := "response"
	actual := ResponseRecordType.String()
	require.Equal(t, expected, actual)

	expected = "request"
	actual = RequestRecordType.String()
	require.Equal(t, expected, actual)

	expected = "revisit"
	actual = RevisitRecordType.String()
	require.Equal(t, expected, actual)
}

func TestParseRecordType(t *testing.T) {
	v := "response"
	actual, err := ParseRecordType(v)
	require.Nil(t, err)
	require.Equal(t, ResponseRecordType, actual)

	v = "request"
	actual, err = ParseRecordType(v)
	require.Nil(t, err)
	require.Equal(t, RequestRecordType, actual)

	v = "revisit"
	actual, err = ParseRecordType(v)
	require.Nil(t, err)
	require.Equal(t, RevisitRecordType, actual)
}

func TestRecordFormat(t *testing.T) {
	ts, err := time.Parse(time.RFC3339, "2020-01-02T03:04:05-01:00")
	require.Nil(t, err)
	ty, err := ParseRecordType("response")
	require.Nil(t, err)
	rec := &Record{
		SURT:      "com,cnn)/",
		Timestamp: ts,
		Type:      ty,
		Content:   []byte(`{'uri':'cnn.com/world', 'sha':'<SHA-1 digest>', 'ref':'warcfile:IAH-20070824123353-00393-cnn.com.arc.gz#25523382'}`),
	}
	expected := `com,cnn)/ 2020-01-02T03:04:05-01:00 response {'uri':'cnn.com/world', 'sha':'<SHA-1 digest>', 'ref':'warcfile:IAH-20070824123353-00393-cnn.com.arc.gz#25523382'}`
	actual := rec.Format()
	require.Equal(t, expected, actual)
}

func TestParseRecord(t *testing.T) {
	ts, err := time.Parse(time.RFC3339, "2020-01-02T03:04:05-01:00")
	require.Nil(t, err)
	ty, err := ParseRecordType("response")
	require.Nil(t, err)
	expected := Record{
		SURT:      "com,cnn)/",
		Timestamp: ts,
		Type:      ty,
		Content:   []byte(`{'uri':'cnn.com/world', 'sha':'<SHA-1 digest>', 'ref':'warcfile:IAH-20070824123353-00393-cnn.com.arc.gz#25523382'}`),
	}
	actual, err := parseRecord(expected.Format())
	require.Nil(t, err)
	require.Equal(t, expected.Format(), actual.Format())
}
