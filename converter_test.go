package carbon

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_ToCookieString(t *testing.T) {
	c := Create(2020, time.May, 30, 11, 20, 30, 103, getLocation("UTC"))
	assert.EqualValues(t, "Saturday, 30-May-2020 11:20:30 UTC", c.ToCookieString())
}

func Test_ToDateTimeString_Default(t *testing.T) {
	c := Create(2020, time.May, 30, 11, 20, 30, 103, getLocation("UTC"))
	dts := c.ToDateTimeString("")
	assert.EqualValues(t, "2020-05-30 11:20:30", dts)
}

func Test_ToDateTimeString_Second(t *testing.T) {
	c := Create(2020, time.May, 30, 11, 20, 30, 103, getLocation("UTC"))
	dts := c.ToDateTimeString("second")
	assert.EqualValues(t, "2020-05-30 11:20:30", dts)
}

func Test_ToDateTimeString_Minute(t *testing.T) {
	c := Create(2020, time.May, 30, 11, 20, 30, 103, getLocation("UTC"))
	dts := c.ToDateTimeString("minute")
	assert.EqualValues(t, "2020-05-30 11:20", dts)
}

func Test_ToDateTimeString_Millisecond(t *testing.T) {
	c := Create(2020, time.May, 30, 11, 20, 30, 103103103, getLocation("UTC"))
	dts1 := c.ToDateTimeString("m")
	dts2 := c.ToDateTimeString("millisecond")
	assert.EqualValues(t, "2020-05-30 11:20:30.103", dts1)
	assert.EqualValues(t, "2020-05-30 11:20:30.103", dts2)
}

func Test_ToDateTimeString_Microsecond(t *testing.T) {
	c := Create(2020, time.May, 30, 11, 20, 30, 103103103, getLocation("UTC"))
	dts1 := c.ToDateTimeString("µ")
	dts2 := c.ToDateTimeString("microsecond")
	assert.EqualValues(t, "2020-05-30 11:20:30.103103", dts1)
	assert.EqualValues(t, "2020-05-30 11:20:30.103103", dts2)
}

func Test_ToDateString(t *testing.T) {
	c := Create(2020, time.May, 30, 11, 20, 30, 103103103, getLocation("UTC"))
	assert.EqualValues(t, "2020-05-30", c.ToDateString())
}

func Test_ToDayDateTimeString(t *testing.T) {
	c := Create(2020, time.May, 30, 11, 20, 30, 103103103, getLocation("UTC"))
	assert.EqualValues(t, "Sat, May 30, 2020 11:20 AM", c.ToDayDateTimeString())
}

func Test_ToFormattedDateString(t *testing.T) {
	c := Create(2020, time.May, 30, 11, 20, 30, 103103103, getLocation("UTC"))
	assert.EqualValues(t, "May 30, 2020", c.ToFormattedDateString())
}

func Test_ToRfc1036String(t *testing.T) {
	c := Create(2020, time.May, 30, 11, 20, 30, 103103103, getLocation("UTC"))
	assert.EqualValues(t, "Sat, 30 May 20 11:20:30 +0000", c.ToRfc1036String())
}

func Test_ToRfc2822String(t *testing.T) {
	c := Create(2020, time.May, 30, 11, 20, 30, 103103103, getLocation("UTC"))
	assert.EqualValues(t, "Sat, 30 May 2020 11:20:30 +0000", c.ToRfc2822String())
}

func Test_ToRfc3339String(t *testing.T) {
	c := Create(2020, time.May, 30, 11, 20, 30, 103103103, getLocation("UTC"))
	assert.EqualValues(t, "2020-05-30T11:20:30Z", c.ToRfc3339String())
}

func Test_ToRfc822String(t *testing.T) {
	c := Create(2020, time.May, 30, 11, 20, 30, 103103103, getLocation("UTC"))
	assert.EqualValues(t, "30 May 20 11:20 UTC", c.ToRfc822String())
}

func Test_ToRfc850String(t *testing.T) {
	c := Create(2020, time.May, 30, 11, 20, 30, 103103103, getLocation("UTC"))
	assert.EqualValues(t, "Saturday, 30-May-20 11:20:30 UTC", c.ToRfc850String())
}

func Test_ToRssString(t *testing.T) {
	c := Create(2020, time.May, 30, 11, 20, 30, 103103103, getLocation("UTC"))
	assert.EqualValues(t, "Sat, 30 May 2020 11:20:30 +0000", c.ToRssString())
}

func Test_ToTimeString(t *testing.T) {
	c := Create(2020, time.May, 30, 11, 20, 30, 103103103, getLocation("UTC"))
	assert.EqualValues(t, "11:20:30", c.ToTimeString())
}

func Test_ToAtomString(t *testing.T) {
	c := Create(2020, time.May, 30, 11, 20, 30, 103103103, getLocation("UTC"))
	assert.EqualValues(t, "2020-05-30T11:20:30+00:00", c.ToAtomString())
}

func Test_ToRfc1123String(t *testing.T) {
	c := Create(2020, time.May, 30, 11, 20, 30, 103103103, getLocation("UTC"))
	assert.EqualValues(t, "Sat, 30 May 2020 11:20:30 UTC", c.ToRfc1123String())
}

func Test_ToW3cString(t *testing.T) {
	c := Create(2020, time.May, 30, 11, 20, 30, 103103103, getLocation("UTC"))
	assert.EqualValues(t, "2020-05-30T11:20:30+00:00", c.ToW3cString())
}

func Test_ToDateTimeLocalString(t *testing.T) {
	c := Create(2020, time.May, 30, 11, 20, 30, 103103103, getLocation("UTC"))
	assert.EqualValues(t, "2020-05-30T11:20:30", c.ToDateTimeLocalString("second"))
}
