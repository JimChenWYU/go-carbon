package carbon

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCarbon_SetTimeZone(t *testing.T) {
	c, _ := Create(2020, time.May, 30, 11, 20, 30, 130130130, "UTC")
	assert.EqualValues(t, "UTC", c.Timezone())
	c.SetTimeZone("Africa/Cairo")
	assert.EqualValues(t, "Africa/Cairo", c.Timezone())
}

func TestCarbon_AddDays(t *testing.T) {
	c1, _ := Create(2019, time.May, 29, 0, 0, 0, 0, "UTC")
	d := c1.AddDays(1)
	expected, _ := Create(2019, time.May, 30, 0, 0, 0, 0, "UTC")

	assert.Equal(t, expected, d, "The day should be equal to 30")
}

func TestCarbon_AddHours(t *testing.T) {
	c1, _ := Create(2019, time.May, 29, 0, 0, 0, 0, "UTC")
	d := c1.AddHours(1)
	expected, _ := Create(2019, time.May, 29, 1, 0, 0, 0, "UTC")

	assert.Equal(t, expected, d, "The hour should be equal to 1")
}

func TestCarbon_AddMinutes(t *testing.T) {
	c1, _ := Create(2019, time.May, 29, 0, 0, 0, 0, "UTC")
	d := c1.AddMinutes(1)
	expected, _ := Create(2019, time.May, 29, 0, 1, 0, 0, "UTC")

	assert.Equal(t, expected, d, "The minute should be equal to 1")
}

func TestCarbon_AddSeconds(t *testing.T) {
	c1, _ := Create(2019, time.May, 29, 0, 0, 0, 0, "UTC")
	d := c1.AddSeconds(1)
	expected, _ := Create(2019, time.May, 29, 0, 0, 1, 0, "UTC")

	assert.Equal(t, expected, d, "The second should be equal to 1")
}

func TestCarbon_AddMonths(t *testing.T) {
	c1, _ := Create(2019, time.May, 29, 0, 0, 0, 0, "UTC")
	d := c1.AddMonths(1)
	expected, _ := Create(2019, time.June, 29, 0, 0, 0, 0, "UTC")

	assert.Equal(t, expected, d, "The month should be equal to June")
}

func TestCarbon_AddWeeks(t *testing.T) {
	c1, _ := Create(2019, time.May, 22, 0, 0, 0, 0, "UTC")
	d := c1.AddWeeks(1)
	expected, _ := Create(2019, time.May, 29, 0, 0, 0, 0, "UTC")

	assert.Equal(t, expected, d, "The day should be equal to 29")
}

func TestCarbon_AddYears(t *testing.T) {
	c1, _ := Create(2019, time.May, 22, 0, 0, 0, 0, "UTC")
	d := c1.AddYears(1)
	expected, _ := Create(2020, time.May, 22, 0, 0, 0, 0, "UTC")

	assert.Equal(t, expected, d, "The year should be equal to 2020")
}

func TestCarbon_SubMonths(t *testing.T) {
	c1, _ := Create(2019, time.May, 22, 0, 0, 0, 0, "UTC")
	d := c1.SubMonths(1)
	expected, _ := Create(2019, time.April, 22, 0, 0, 0, 0, "UTC")

	assert.Equal(t, expected, d, "The month should be equal to April")
}

func TestCarbon_SubDays(t *testing.T) {
	c1, _ := Create(2019, time.May, 22, 0, 0, 0, 0, "UTC")
	d := c1.SubDays(1)
	expected, _ := Create(2019, time.May, 21, 0, 0, 0, 0, "UTC")

	assert.Equal(t, expected, d, "The day should be equal to 21")
}

func TestCarbon_SubHours(t *testing.T) {
	c1, _ := Create(2019, time.May, 22, 1, 0, 0, 0, "UTC")
	d := c1.SubHours(1)
	expected, _ := Create(2019, time.May, 22, 0, 0, 0, 0, "UTC")

	assert.Equal(t, expected, d, "The day should be equal to 0")
}

func TestCarbon_SubMinutes(t *testing.T) {
	c1, _ := Create(2019, time.May, 22, 0, 1, 0, 0, "UTC")
	d := c1.SubMinutes(1)
	expected, _ := Create(2019, time.May, 22, 0, 0, 0, 0, "UTC")

	assert.Equal(t, expected, d, "The minute should be equal to 0")
}

func TestCarbon_SubSeconds(t *testing.T) {
	c1, _ := Create(2019, time.May, 22, 0, 0, 1, 0, "UTC")
	d := c1.SubSeconds(1)
	expected, _ := Create(2019, time.May, 22, 0, 0, 0, 0, "UTC")

	assert.Equal(t, expected, d, "The second should be equal to 0")
}

func TestCarbon_SubWeeks(t *testing.T) {
	c1, _ := Create(2019, time.May, 29, 0, 0, 0, 0, "UTC")
	d := c1.SubWeeks(1)
	expected, _ := Create(2019, time.May, 22, 0, 0, 0, 0, "UTC")

	assert.Equal(t, expected, d, "The day should be equal to 22")
}

func TestCarbon_SubYears(t *testing.T) {
	c1, _ := Create(2019, time.May, 22, 0, 0, 0, 0, "UTC")
	d := c1.SubYears(10)
	expected, _ := Create(2009, time.May, 22, 0, 0, 0, 0, "UTC")

	assert.Equal(t, expected, d, "The year should be equal to 2009")
}

func TestCarbon_SetYears(t *testing.T) {
	c1, _ := Create(2019, time.May, 22, 0, 0, 0, 0, "UTC")
	d := c1.SetYears(2020).Year()
	assert.EqualValues(t, 2020, d)
}

func TestCarbon_SetDays(t *testing.T) {
	c1, _ := Create(2019, time.May, 22, 0, 0, 0, 0, "UTC")
	d := c1.SetDays(23).Day()
	assert.EqualValues(t, 23, d)
}

func TestCarbon_SetHours(t *testing.T) {
	c1, _ := Create(2019, time.May, 22, 0, 0, 0, 0, "UTC")
	d := c1.SetHours(1).Hour()
	assert.EqualValues(t, 1, d)
}

func TestCarbon_SetMinutes(t *testing.T) {
	c1, _ := Create(2019, time.May, 22, 0, 0, 0, 0, "UTC")
	d := c1.SetMinutes(1).Minute()
	assert.EqualValues(t, 1, d)
}

func TestCarbon_SetMonths(t *testing.T) {
	c1, _ := Create(2019, time.May, 22, 0, 0, 0, 0, "UTC")
	d := c1.SetMonths(time.June).Month()
	assert.EqualValues(t, time.June, d)
}

func TestCarbon_SetNanoSeconds(t *testing.T) {
	c1, _ := Create(2019, time.May, 22, 0, 0, 0, 0, "UTC")
	d := c1.SetNanoSeconds(102102102).Nanosecond()
	assert.EqualValues(t, 102102102, d)
}

func TestCarbon_SetSeconds(t *testing.T) {
	c1, _ := Create(2019, time.May, 22, 0, 0, 0, 0, "UTC")
	d := c1.SetNanoSeconds(102102102).Nanosecond()
	assert.EqualValues(t, 102102102, d)
}

func TestCarbon_SetTimestamp(t *testing.T) {
	c1, _ := Create(2019, time.May, 22, 0, 0, 0, 0, "UTC")
	c1.SetTimestamp(1577808000)
	loc, _ := time.LoadLocation("UTC")
	assert.EqualValues(t, "2019-12-31 16:00:00", c1.In(loc).Format(DefaultFormat))
}
