package financier

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDateJSONUnmarshalling(t *testing.T) {
	expected := Date(
		time.Date(2014, time.January, 23, 0, 0, 0, 0, time.UTC),
	)
	actual := Date{}
	assert.NoError(t, actual.UnmarshalJSON([]byte(`"2014-01-23"`)))
	assert.Equal(t, expected, actual)
	actual = Date{}
	assert.Error(t, actual.UnmarshalJSON([]byte(`null`)))
	assert.Equal(t, Date{}, actual)
}

func TestDateJSONMarshalling(t *testing.T) {
	date := Date(
		time.Date(2014, time.January, 23, 0, 0, 0, 0, time.UTC),
	)
	json, err := date.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, `"2014-01-23"`, string(json))
	date = Date{}
	json, err = date.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, `null`, string(json))
}
