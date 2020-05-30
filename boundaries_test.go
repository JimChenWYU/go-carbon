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
