package carbon

import (
	"errors"
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

func TestRawParseUnknown(t *testing.T) {
	_, err := RawParse("sssss", getLocation("Africa/Cairo"))

	assert.NotNil(t, err)
	assert.Equal(t, errors.New("parse syntax error"), err)
}

func TestTodayEET(t *testing.T) {
	today, _ := Today(getLocation("Africa/Cairo"))

	assert.Equal(t, "Africa/Cairo", today.Timezone().String())
}

//func TestTodayUnknown(t *testing.T) {
//	_, err := Today(getLocation("Jupiter/Wondertown"))
//
//	assert.NotNil(t, err)
//}

func TestTomorrowEET(t *testing.T) {
	tomorrow, _ := Tomorrow(getLocation("Africa/Cairo"))

	assert.Equal(t, "Africa/Cairo", tomorrow.Timezone().String())
}

//func TestTomorrowUnknown(t *testing.T) {
//	_, err := Tomorrow(getLocation("Jupiter/Wondertown"))
//
//	assert.NotNil(t, err)
//}

func TestYesterdayEET(t *testing.T) {
	yesterday, _ := Yesterday(getLocation("Africa/Cairo"))

	assert.Equal(t, "Africa/Cairo", yesterday.Timezone().String())
}

//func TestYesterdayUnknown(t *testing.T) {
//	_, err := Tomorrow(getLocation("Jupiter/Wondertown"))
//
//	assert.NotNil(t, err)
//}

func getLocation(loc string) *time.Location {
	l, _ := time.LoadLocation(loc)
	return l
}
