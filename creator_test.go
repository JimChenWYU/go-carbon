package carbon

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	c, _ := Create(2011, time.April, 20, 11, 20, 10, 103, "UTC")

	assert.EqualValues(t, 2011, c.Year())
	assert.EqualValues(t, time.April, c.Month())
	assert.EqualValues(t, 20, c.Day())
	assert.EqualValues(t, 11, c.Hour())
	assert.EqualValues(t, 20, c.Minute())
	assert.EqualValues(t, 10, c.Second())
	assert.EqualValues(t, 103, c.Nanosecond())
	assert.EqualValues(t, "UTC", c.Location().String())
}

func TestCreateFromDate(t *testing.T) {
	c, _ := CreateFromDate(2011, time.April, 20, "UTC")

	assert.EqualValues(t, 2011, c.Year())
	assert.EqualValues(t, time.April, c.Month())
	assert.EqualValues(t, 20, c.Day())
}

func TestCreateFromTime(t *testing.T) {
	c, _ := CreateFromTime(11, 20, 10, 103, "UTC")

	assert.EqualValues(t, 11, c.Hour())
	assert.EqualValues(t, 20, c.Minute())
	assert.EqualValues(t, 10, c.Second())
	assert.EqualValues(t, 103, c.Nanosecond())
}

func TestRawParseUnknown(t *testing.T) {
	_, err := RawParse("sssss", "Africa/Cairo")

	assert.NotNil(t, err)
}

func TestTodayEET(t *testing.T) {
	today, _ := Today("Africa/Cairo")

	assert.Equal(t, "Africa/Cairo", today.Timezone())
}

func TestTodayUnknown(t *testing.T) {
	_, err := Today("Jupiter/Wondertown")

	assert.NotNil(t, err)
}

func TestTomorrowEET(t *testing.T) {
	tomorrow, _ := Tomorrow("Africa/Cairo")

	assert.Equal(t, "Africa/Cairo", tomorrow.Timezone())
}

func TestTomorrowUnknown(t *testing.T) {
	_, err := Tomorrow("Jupiter/Wondertown")

	assert.NotNil(t, err)
}

func TestYesterdayEET(t *testing.T) {
	yesterday, _ := Yesterday("Africa/Cairo")

	assert.Equal(t, "Africa/Cairo", yesterday.Timezone())
}

func TestYesterdayUnknown(t *testing.T) {
	_, err := Tomorrow("Jupiter/Wondertown")

	assert.NotNil(t, err)
}
