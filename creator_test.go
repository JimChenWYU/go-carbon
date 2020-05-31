package carbon

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	c := Create(2011, time.April, 20, 11, 20, 10, 103, getLocation("UTC"))

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
	c := CreateFromDate(2011, time.April, 20, getLocation("UTC"))

	assert.EqualValues(t, 2011, c.Year())
	assert.EqualValues(t, time.April, c.Month())
	assert.EqualValues(t, 20, c.Day())
}

func TestCreateFromTime(t *testing.T) {
	c := CreateFromTime(11, 20, 10, 103, getLocation("UTC"))

	assert.EqualValues(t, 11, c.Hour())
	assert.EqualValues(t, 20, c.Minute())
	assert.EqualValues(t, 10, c.Second())
	assert.EqualValues(t, 103, c.Nanosecond())
}

func TestTodayEET(t *testing.T) {
	today := Today(getLocation("Africa/Cairo"))

	assert.Equal(t, "Africa/Cairo", today.Timezone().String())
}

func TestTomorrowEET(t *testing.T) {
	tomorrow := Tomorrow(getLocation("Africa/Cairo"))

	assert.Equal(t, "Africa/Cairo", tomorrow.Timezone().String())
}

func TestYesterdayEET(t *testing.T) {
	yesterday := Yesterday(getLocation("Africa/Cairo"))

	assert.Equal(t, "Africa/Cairo", yesterday.Timezone().String())
}

func TestNowWithTz(t *testing.T) {
	l, _ := time.LoadLocation("Africa/Cairo")
	c1 := NowWithTz(l)

	assert.EqualValues(t, "Africa/Cairo", c1.Location().String())
}

func TestRawParse(t *testing.T) {
	c := RawParse(DefaultFormat, "2010-10-20 10:20:30", getLocation("Africa/Cairo"))

	assert.EqualValues(t, 2010, c.Year())
	assert.EqualValues(t, time.October, c.Month())
	assert.EqualValues(t, 20, c.Day())
	assert.EqualValues(t, 10, c.Hour())
	assert.EqualValues(t, 20, c.Minute())
	assert.EqualValues(t, 30, c.Second())
	assert.EqualValues(t, "Africa/Cairo", c.Location().String())
}

func TestRawParseWithDefaultTz(t *testing.T) {
	c := RawParseWithDefaultTz(DefaultFormat, "2010-10-20 10:20:30")
	assert.EqualValues(t, "UTC", c.Location().String())
}
