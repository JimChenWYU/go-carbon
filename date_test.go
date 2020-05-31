package carbon

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCarbon_SetTimezone(t *testing.T) {
	c := Create(2020, time.May, 30, 11, 20, 30, 130130130, getLocation("UTC"))
	assert.EqualValues(t, "UTC", c.Timezone().String())
	c.SetTimezone(getLocation("Africa/Cairo"))
	assert.EqualValues(t, "Africa/Cairo", c.Timezone().String())
}

func TestCarbon_SetYears(t *testing.T) {
	c1 := Create(2019, time.May, 22, 0, 0, 0, 0, getLocation("UTC"))
	d := c1.SetYears(2020).Year()
	assert.EqualValues(t, 2020, d)
}

func TestCarbon_SetDays(t *testing.T) {
	c1 := Create(2019, time.May, 22, 0, 0, 0, 0, getLocation("UTC"))
	d := c1.SetDays(23).Day()
	assert.EqualValues(t, 23, d)
}

func TestCarbon_SetHours(t *testing.T) {
	c1 := Create(2019, time.May, 22, 0, 0, 0, 0, getLocation("UTC"))
	d := c1.SetHours(1).Hour()
	assert.EqualValues(t, 1, d)
}

func TestCarbon_SetMinutes(t *testing.T) {
	c1 := Create(2019, time.May, 22, 0, 0, 0, 0, getLocation("UTC"))
	d := c1.SetMinutes(1).Minute()
	assert.EqualValues(t, 1, d)
}

func TestCarbon_SetMonths(t *testing.T) {
	c1 := Create(2019, time.May, 22, 0, 0, 0, 0, getLocation("UTC"))
	d := c1.SetMonths(time.June).Month()
	assert.EqualValues(t, time.June, d)
}

func TestCarbon_SetNanoSeconds(t *testing.T) {
	c1 := Create(2019, time.May, 22, 0, 0, 0, 0, getLocation("UTC"))
	d := c1.SetNanoSeconds(102102102).Nanosecond()
	assert.EqualValues(t, 102102102, d)
}

func TestCarbon_SetSeconds(t *testing.T) {
	c1 := Create(2019, time.May, 22, 0, 0, 0, 0, getLocation("UTC"))
	d := c1.SetNanoSeconds(102102102).Nanosecond()
	assert.EqualValues(t, 102102102, d)
}

func TestCarbon_SetTimestamp(t *testing.T) {
	c1 := Create(2019, time.May, 22, 0, 0, 0, 0, getLocation("UTC"))
	c1.SetTimestamp(1577808000)
	assert.EqualValues(t, "2019-12-31 16:00:00", c1.In(getLocation("UTC")).Format(DefaultFormat))
}

func TestCarbon_SetTime(t *testing.T) {
	c1 := Create(2019, time.May, 22, 0, 0, 0, 0, getLocation("UTC"))
	c1.SetTime(1, 1, 1, 1)
	assert.EqualValues(t, 1, c1.Hour())
	assert.EqualValues(t, 1, c1.Minute())
	assert.EqualValues(t, 1, c1.Second())
	assert.EqualValues(t, 1, c1.Nanosecond())
}

func TestCarbon_SetDate(t *testing.T) {
	c1 := Create(2019, time.May, 22, 0, 0, 0, 0, getLocation("UTC"))
	c1.SetDate(2020, time.April, 1)
	assert.EqualValues(t, 2020, c1.Year())
	assert.EqualValues(t, time.April, c1.Month())
	assert.EqualValues(t, 1, c1.Day())
}

func TestCarbon_AddMilli(t *testing.T) {
	c1 := Create(2019, time.May, 22, 0, 0, 0, 0, getLocation("UTC"))
	c2 := c1.AddMilli(1)
	assert.EqualValues(t, time.Duration(1)*time.Millisecond, c2.Nanosecond())
}

func TestCarbon_AddMicro(t *testing.T) {
	c1 := Create(2019, time.May, 22, 0, 0, 0, 0, getLocation("UTC"))
	c2 := c1.AddMicro(1)
	assert.EqualValues(t, time.Duration(1)*time.Microsecond, c2.Nanosecond())
}

func TestCarbon_AddUnit(t *testing.T) {
	c1 := Create(2019, time.May, 22, 0, 0, 0, 0, getLocation("UTC"))
	d := c1.AddUnit("micro", 1)
	expected := Create(2019, time.May, 22, 0, 0, 0, 1*int(time.Microsecond), getLocation("UTC"))
	assert.Equal(t, expected, d, "The microsecond should be equal to 1")

	c1 = Create(2019, time.May, 22, 0, 0, 0, 0, getLocation("UTC"))
	d = c1.AddUnit("milli", 1)
	expected = Create(2019, time.May, 22, 0, 0, 0, 1*int(time.Millisecond), getLocation("UTC"))
	assert.Equal(t, expected, d, "The millisecond should be equal to 1")

	c1 = Create(2019, time.May, 29, 0, 0, 0, 0, getLocation("UTC"))
	d = c1.AddUnit("second", 1)
	expected = Create(2019, time.May, 29, 0, 0, 1, 0, getLocation("UTC"))
	assert.Equal(t, expected, d, "The second should be equal to 1")

	c1 = Create(2019, time.May, 29, 0, 0, 0, 0, getLocation("UTC"))
	d = c1.AddUnit("minute", 1)
	expected = Create(2019, time.May, 29, 0, 1, 0, 0, getLocation("UTC"))
	assert.Equal(t, expected, d, "The minute should be equal to 1")

	c1 = Create(2019, time.May, 29, 0, 0, 0, 0, getLocation("UTC"))
	d = c1.AddUnit("hour", 1)
	expected = Create(2019, time.May, 29, 1, 0, 0, 0, getLocation("UTC"))
	assert.Equal(t, expected, d, "The hour should be equal to 1")

	c1 = Create(2019, time.May, 29, 0, 0, 0, 0, getLocation("UTC"))
	d = c1.AddUnit("day", 1)
	expected = Create(2019, time.May, 30, 0, 0, 0, 0, getLocation("UTC"))
	assert.Equal(t, expected, d, "The day should be equal to 30")

	c1 = Create(2019, time.May, 29, 0, 0, 0, 0, getLocation("UTC"))
	d = c1.AddUnit("month", 1)
	expected = Create(2019, time.June, 29, 0, 0, 0, 0, getLocation("UTC"))
	assert.Equal(t, expected, d, "The month should be equal to June")

	c1 = Create(2019, time.May, 29, 0, 0, 0, 0, getLocation("UTC"))
	d = c1.AddUnit("quarter", 1)
	expected = Create(2019, time.August, 29, 0, 0, 0, 0, getLocation("UTC"))
	assert.Equal(t, expected, d, "The month should be equal to August")

	c1 = Create(2019, time.May, 29, 0, 0, 0, 0, getLocation("UTC"))
	d = c1.AddUnit("year", 1)
	expected = Create(2020, time.May, 29, 0, 0, 0, 0, getLocation("UTC"))
	assert.Equal(t, expected, d, "The year should be equal to 2020")

	c1 = Create(2019, time.May, 29, 0, 0, 0, 0, getLocation("UTC"))
	d = c1.AddUnit("decade", 1)
	expected = Create(2029, time.May, 29, 0, 0, 0, 0, getLocation("UTC"))
	assert.Equal(t, expected, d, "The year should be equal to 2029")

	c1 = Create(2019, time.May, 29, 0, 0, 0, 0, getLocation("UTC"))
	d = c1.AddUnit("century", 1)
	expected = Create(2119, time.May, 29, 0, 0, 0, 0, getLocation("UTC"))
	assert.Equal(t, expected, d, "The year should be equal to 2119")

	c1 = Create(2019, time.May, 29, 0, 0, 0, 0, getLocation("UTC"))
	d = c1.AddUnit("millennium", 1)
	expected = Create(3019, time.May, 29, 0, 0, 0, 0, getLocation("UTC"))
	assert.Equal(t, expected, d, "The year should be equal to 3019")

	assert.PanicsWithError(t, "unknown unit to add", func() {
		c1 = Create(2019, time.May, 29, 0, 0, 0, 0, getLocation("UTC"))
		c1.AddUnit("unknown unit", 1)
	})
}

func TestCarbon_SubUnit(t *testing.T) {
	c1 := Create(2019, time.May, 29, 0, 0, 0, 1*int(time.Millisecond), getLocation("UTC"))
	d := c1.SubUnit("milli", 1)
	expected := Create(2019, time.May, 29, 0, 0, 0, 0, getLocation("UTC"))
	assert.Equal(t, expected, d, "The millisecond should be equal to 0")

	c1 = Create(2019, time.May, 29, 0, 0, 0, 1*int(time.Microsecond), getLocation("UTC"))
	d = c1.SubUnit("micro", 1)
	expected = Create(2019, time.May, 29, 0, 0, 0, 0, getLocation("UTC"))
	assert.Equal(t, expected, d, "The microsecond should be equal to 0")

	c1 = Create(2019, time.May, 29, 0, 0, 1, 0, getLocation("UTC"))
	d = c1.SubUnit("second", 1)
	expected = Create(2019, time.May, 29, 0, 0, 0, 0, getLocation("UTC"))
	assert.Equal(t, expected, d, "The second should be equal to 0")

	c1 = Create(2019, time.May, 29, 0, 1, 0, 0, getLocation("UTC"))
	d = c1.SubUnit("minute", 1)
	expected = Create(2019, time.May, 29, 0, 0, 0, 0, getLocation("UTC"))
	assert.Equal(t, expected, d, "The minute should be equal to 0")

	c1 = Create(2019, time.May, 29, 1, 0, 0, 0, getLocation("UTC"))
	d = c1.SubUnit("hour", 1)
	expected = Create(2019, time.May, 29, 0, 0, 0, 0, getLocation("UTC"))
	assert.Equal(t, expected, d, "The hour should be equal to 0")

	c1 = Create(2019, time.May, 29, 0, 0, 0, 0, getLocation("UTC"))
	d = c1.SubUnit("day", 1)
	expected = Create(2019, time.May, 28, 0, 0, 0, 0, getLocation("UTC"))
	assert.Equal(t, expected, d, "The day should be equal to 28")

	c1 = Create(2019, time.May, 29, 0, 0, 0, 0, getLocation("UTC"))
	d = c1.SubUnit("month", 1)
	expected = Create(2019, time.April, 29, 0, 0, 0, 0, getLocation("UTC"))
	assert.Equal(t, expected, d, "The month should be equal to April")

	c1 = Create(2019, time.May, 29, 0, 0, 0, 0, getLocation("UTC"))
	d = c1.SubUnit("quarter", 1)
	expected = Create(2019, time.February, 29, 0, 0, 0, 0, getLocation("UTC"))
	assert.Equal(t, expected, d, "The month should be equal to February")

	c1 = Create(2019, time.May, 29, 0, 0, 0, 0, getLocation("UTC"))
	d = c1.SubUnit("year", 1)
	expected = Create(2018, time.May, 29, 0, 0, 0, 0, getLocation("UTC"))
	assert.Equal(t, expected, d, "The year should be equal to 2018")

	c1 = Create(2019, time.May, 29, 0, 0, 0, 0, getLocation("UTC"))
	d = c1.SubUnit("decade", 1)
	expected = Create(2009, time.May, 29, 0, 0, 0, 0, getLocation("UTC"))
	assert.Equal(t, expected, d, "The year should be equal to 2009")

	c1 = Create(2019, time.May, 29, 0, 0, 0, 0, getLocation("UTC"))
	d = c1.SubUnit("century", 1)
	expected = Create(1919, time.May, 29, 0, 0, 0, 0, getLocation("UTC"))
	assert.Equal(t, expected, d, "The year should be equal to 1919")

	c1 = Create(2019, time.May, 29, 0, 0, 0, 0, getLocation("UTC"))
	d = c1.SubUnit("millennium", 1)
	expected = Create(1019, time.May, 29, 0, 0, 0, 0, getLocation("UTC"))
	assert.Equal(t, expected, d, "The year should be equal to 1019")

	assert.PanicsWithError(t, "unknown unit to add", func() {
		c1 = Create(2019, time.May, 29, 0, 0, 0, 0, getLocation("UTC"))
		c1.SubUnit("unknown unit", 1)
	})
}
