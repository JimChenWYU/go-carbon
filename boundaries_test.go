package carbon

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const (
	MaxNS = 999999999
)

func TestCarbon_startOfDay(t *testing.T) {
	c1, _ := Create(2019, time.May, 22, 1, 1, 1, 1, "UTC")
	c2 := c1.StartOfDay()

	assert.EqualValues(t, 0, c2.Hour())
	assert.EqualValues(t, 0, c2.Minute())
	assert.EqualValues(t, 0, c2.Second())
	assert.EqualValues(t, 0, c2.Nanosecond())
}

func TestCarbon_EndOfDay(t *testing.T) {
	c1, _ := Create(2019, time.May, 22, 1, 1, 1, 1, "UTC")
	c2 := c1.EndOfDay()

	assert.EqualValues(t, 23, c2.Hour())
	assert.EqualValues(t, 59, c2.Minute())
	assert.EqualValues(t, 59, c2.Second())
	assert.EqualValues(t, MaxNS, c2.Nanosecond())
}

func TestCarbon_StartOfMonth(t *testing.T) {
	c1, _ := Create(2019, time.May, 22, 1, 1, 1, 1, "UTC")
	c2 := c1.StartOfMonth()

	assert.EqualValues(t, 1, c2.Day())
	assert.EqualValues(t, 0, c2.Hour())
	assert.EqualValues(t, 0, c2.Minute())
	assert.EqualValues(t, 0, c2.Second())
	assert.EqualValues(t, 0, c2.Nanosecond())
}

func TestCarbon_EndOfMonth(t *testing.T) {
	c1, _ := Create(2019, time.May, 22, 1, 1, 1, 1, "UTC")
	c2 := c1.EndOfMonth()
	assert.EqualValues(t, 31, c2.Day())
	assert.EqualValues(t, 23, c2.Hour())
	assert.EqualValues(t, 59, c2.Minute())
	assert.EqualValues(t, 59, c2.Second())
	assert.EqualValues(t, MaxNS, c2.Nanosecond())

	c1, _ = Create(2019, time.April, 22, 1, 1, 1, 1, "UTC")
	c2 = c1.EndOfMonth()
	assert.EqualValues(t, 30, c2.Day())
	assert.EqualValues(t, 23, c2.Hour())
	assert.EqualValues(t, 59, c2.Minute())
	assert.EqualValues(t, 59, c2.Second())
	assert.EqualValues(t, MaxNS, c2.Nanosecond())

	c1, _ = Create(2019, time.July, 22, 1, 1, 1, 1, "UTC")
	c2 = c1.EndOfMonth()
	assert.EqualValues(t, 31, c2.Day())
	assert.EqualValues(t, 23, c2.Hour())
	assert.EqualValues(t, 59, c2.Minute())
	assert.EqualValues(t, 59, c2.Second())
	assert.EqualValues(t, MaxNS, c2.Nanosecond())

	c1, _ = Create(2019, time.August, 22, 1, 1, 1, 1, "UTC")
	c2 = c1.EndOfMonth()
	assert.EqualValues(t, 31, c2.Day())
	assert.EqualValues(t, 23, c2.Hour())
	assert.EqualValues(t, 59, c2.Minute())
	assert.EqualValues(t, 59, c2.Second())
	assert.EqualValues(t, MaxNS, c2.Nanosecond())

	c1, _ = Create(2019, time.February, 10, 1, 1, 1, 1, "UTC")
	c2 = c1.EndOfMonth()
	assert.EqualValues(t, 28, c2.Day())
	assert.EqualValues(t, 23, c2.Hour())
	assert.EqualValues(t, 59, c2.Minute())
	assert.EqualValues(t, 59, c2.Second())
	assert.EqualValues(t, MaxNS, c2.Nanosecond())

	c1, _ = Create(2008, time.February, 10, 1, 1, 1, 1, "UTC")
	c2 = c1.EndOfMonth()
	assert.EqualValues(t, 29, c2.Day())
	assert.EqualValues(t, 23, c2.Hour())
	assert.EqualValues(t, 59, c2.Minute())
	assert.EqualValues(t, 59, c2.Second())
	assert.EqualValues(t, MaxNS, c2.Nanosecond())
}

func TestCarbon_StartOfQuarter(t *testing.T) {
	c1, _ := Create(2008, time.February, 10, 1, 1, 1, 1, "UTC")
	c2 := c1.StartOfQuarter()
	assert.EqualValues(t, time.January, c2.Month())
	assert.EqualValues(t, 1, c2.Day())
	assert.EqualValues(t, 0, c2.Hour())
	assert.EqualValues(t, 0, c2.Minute())
	assert.EqualValues(t, 0, c2.Second())
	assert.EqualValues(t, 0, c2.Nanosecond())
}

func TestCarbon_EndOfQuarter(t *testing.T) {
	c1, _ := Create(2008, time.February, 10, 1, 1, 1, 1, "UTC")
	c2 := c1.EndOfQuarter()
	assert.EqualValues(t, time.March, c2.Month())
	assert.EqualValues(t, 31, c2.Day())
	assert.EqualValues(t, 23, c2.Hour())
	assert.EqualValues(t, 59, c2.Minute())
	assert.EqualValues(t, 59, c2.Second())
	assert.EqualValues(t, MaxNS, c2.Nanosecond())
}

func TestCarbon_StartOfYear(t *testing.T) {
	c1, _ := Create(2008, time.February, 10, 1, 1, 1, 1, "UTC")
	c2 := c1.StartOfYear()
	assert.EqualValues(t, time.January, c2.Month())
	assert.EqualValues(t, 1, c2.Day())
	assert.EqualValues(t, 0, c2.Hour())
	assert.EqualValues(t, 0, c2.Minute())
	assert.EqualValues(t, 0, c2.Second())
	assert.EqualValues(t, 0, c2.Nanosecond())
}

func TestCarbon_EndOfYear(t *testing.T) {
	c1, _ := Create(2008, time.February, 10, 1, 1, 1, 1, "UTC")
	c2 := c1.EndOfYear()
	assert.EqualValues(t, time.December, c2.Month())
	assert.EqualValues(t, 31, c2.Day())
	assert.EqualValues(t, 23, c2.Hour())
	assert.EqualValues(t, 59, c2.Minute())
	assert.EqualValues(t, 59, c2.Second())
	assert.EqualValues(t, MaxNS, c2.Nanosecond())
}

func TestCarbon_StartOfDecade(t *testing.T) {
	c1, _ := Create(2008, time.February, 10, 1, 1, 1, 1, "UTC")
	c2 := c1.StartOfDecade()
	assert.EqualValues(t, 2000, c2.Year())
	assert.EqualValues(t, time.January, c2.Month())
	assert.EqualValues(t, 1, c2.Day())
}

func TestCarbon_EndOfDecade(t *testing.T) {
	c1, _ := Create(2008, time.February, 10, 1, 1, 1, 1, "UTC")
	c2 := c1.EndOfDecade()
	assert.EqualValues(t, 2009, c2.Year())
	assert.EqualValues(t, time.December, c2.Month())
	assert.EqualValues(t, 31, c2.Day())
}

func TestCarbon_StartOfCentury(t *testing.T) {
	c1, _ := Create(2008, time.February, 10, 1, 1, 1, 1, "UTC")
	c2 := c1.StartOfCentury()
	assert.EqualValues(t, 2001, c2.Year())
	assert.EqualValues(t, time.January, c2.Month())
	assert.EqualValues(t, 1, c2.Day())
}

func TestCarbon_EndOfCentury(t *testing.T) {
	c1, _ := Create(2008, time.February, 10, 1, 1, 1, 1, "UTC")
	c2 := c1.EndOfCentury()
	assert.EqualValues(t, 2100, c2.Year())
	assert.EqualValues(t, time.December, c2.Month())
	assert.EqualValues(t, 31, c2.Day())
}

func TestCarbon_StartOfMillennium(t *testing.T) {
	c1, _ := Create(2008, time.February, 10, 1, 1, 1, 1, "UTC")
	c2 := c1.StartOfMillennium()
	assert.EqualValues(t, 2001, c2.Year())
	assert.EqualValues(t, time.January, c2.Month())
	assert.EqualValues(t, 1, c2.Day())
}

func TestCarbon_EndOfMillennium(t *testing.T) {
	c1, _ := Create(2008, time.February, 10, 1, 1, 1, 1, "UTC")
	c2 := c1.EndOfMillennium()
	assert.EqualValues(t, 3000, c2.Year())
	assert.EqualValues(t, time.December, c2.Month())
	assert.EqualValues(t, 31, c2.Day())
}

func TestCarbon_StartOfWeek(t *testing.T) {
	c1, _ := Create(2020, time.May, 28, 1, 1, 1, 1, "UTC")
	c2 := c1.StartOfWeek(time.Sunday)
	assert.EqualValues(t, 24, c2.Day())
}

func TestCarbon_EndOfWeek(t *testing.T) {
	c1, _ := Create(2020, time.May, 28, 1, 1, 1, 1, "UTC")
	c2 := c1.EndOfWeek(time.Saturday)
	assert.EqualValues(t, 30, c2.Day())
}

func TestCarbon_StartOfHour(t *testing.T) {
	c1, _ := Create(2020, time.May, 28, 1, 1, 1, 1, "UTC")
	c2 := c1.StartOfHour()
	assert.EqualValues(t, 0, c2.Minute())
	assert.EqualValues(t, 0, c2.Second())
	assert.EqualValues(t, 0, c2.Nanosecond())
}

func TestCarbon_EndOfHour(t *testing.T) {
	c1, _ := Create(2020, time.May, 28, 1, 1, 1, 1, "UTC")
	c2 := c1.EndOfHour()
	assert.EqualValues(t, 59, c2.Minute())
	assert.EqualValues(t, 59, c2.Second())
	assert.EqualValues(t, MaxNS, c2.Nanosecond())
}

func TestCarbon_StartOfMinute(t *testing.T) {
	c1, _ := Create(2020, time.May, 28, 1, 1, 1, 1, "UTC")
	c2 := c1.StartOfMinute()
	assert.EqualValues(t, 0, c2.Second())
	assert.EqualValues(t, 0, c2.Nanosecond())
}

func TestCarbon_EndOfMinute(t *testing.T) {
	c1, _ := Create(2020, time.May, 28, 1, 1, 1, 1, "UTC")
	c2 := c1.EndOfMinute()
	assert.EqualValues(t, 59, c2.Second())
	assert.EqualValues(t, MaxNS, c2.Nanosecond())
}

func TestCarbon_StartOfSecond(t *testing.T) {
	c1, _ := Create(2020, time.May, 28, 1, 1, 1, 1, "UTC")
	c2 := c1.StartOfSecond()
	assert.EqualValues(t, 0, c2.Nanosecond())
}

func TestCarbon_EndOfSecond(t *testing.T) {
	c1, _ := Create(2020, time.May, 28, 1, 1, 1, 1, "UTC")
	c2 := c1.EndOfSecond()
	assert.EqualValues(t, MaxNS, c2.Nanosecond())
}
