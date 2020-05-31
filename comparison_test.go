package carbon

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCarbon_Between_Equal_True(t *testing.T) {
	c1 := CreateFromDate(2000, time.January, 15, getLocation("UTC"))

	min1 := CreateFromDate(2000, time.January, 1, getLocation("UTC"))
	max1 := CreateFromDate(2000, time.January, 31, getLocation("UTC"))
	assert.True(t, c1.Between(min1, max1, true))

	min2 := time.Date(2000, time.January, 1, 0, 0, 0, 0, getLocation("UTC"))
	max2 := time.Date(2000, time.January, 31, 0, 0, 0, 0, getLocation("UTC"))
	assert.True(t, c1.Between(min2, max2, true))
}

func TestCarbon_Between_Not_Equal_True(t *testing.T) {
	c1 := CreateFromDate(2000, time.January, 15, getLocation("UTC"))

	min1 := CreateFromDate(2000, time.January, 1, getLocation("UTC"))
	max1 := CreateFromDate(2000, time.January, 31, getLocation("UTC"))
	assert.True(t, c1.Between(min1, max1, false))

	min2 := time.Date(2000, time.January, 1, 0, 0, 0, 0, getLocation("UTC"))
	max2 := time.Date(2000, time.January, 31, 0, 0, 0, 0, getLocation("UTC"))
	assert.True(t, c1.Between(min2, max2, false))
}

func TestCarbon_EqualTo_True(t *testing.T) {
	c1 := CreateFromTime(10, 0, 0, 0, getLocation("UTC"))
	c2 := CreateFromTime(10, 0, 0, 0, getLocation("UTC"))
	assert.True(t, c1.EqualTo(c2))
}

func TestCarbon_EqualTo_False(t *testing.T) {
	c1 := CreateFromTime(10, 0, 0, 0, getLocation("UTC"))
	c2 := CreateFromTime(11, 0, 0, 0, getLocation("UTC"))
	assert.False(t, c1.EqualTo(c2))
}

func TestCarbon_GreaterThan_True(t *testing.T) {
	c1 := CreateFromTime(11, 0, 0, 0, getLocation("UTC"))
	c2 := CreateFromTime(10, 0, 0, 0, getLocation("UTC"))
	assert.True(t, c1.GreaterThan(c2))
}

func TestCarbon_GreaterThan_False(t *testing.T) {
	c1 := CreateFromTime(9, 0, 0, 0, getLocation("UTC"))
	c2 := CreateFromTime(10, 0, 0, 0, getLocation("UTC"))
	assert.False(t, c1.GreaterThan(c2))
}

func TestCarbon_GreaterThanOrEqualTo_True(t *testing.T) {
	c1 := CreateFromTime(11, 0, 0, 0, getLocation("UTC"))
	c2 := CreateFromTime(10, 0, 0, 0, getLocation("UTC"))
	assert.True(t, c1.GreaterThanOrEqualTo(c2))

	c1 = CreateFromTime(11, 0, 0, 0, getLocation("UTC"))
	c2 = CreateFromTime(11, 0, 0, 0, getLocation("UTC"))
	assert.True(t, c1.GreaterThanOrEqualTo(c2))
}

func TestCarbon_GreaterThanOrEqualTo_False(t *testing.T) {
	c1 := CreateFromTime(3, 0, 0, 0, getLocation("UTC"))
	c2 := CreateFromTime(10, 0, 0, 0, getLocation("UTC"))
	assert.False(t, c1.GreaterThanOrEqualTo(c2))
}

func TestCarbon_IsLeadYear_True(t *testing.T) {
	c1 := CreateFromDate(2000, time.January, 1, getLocation("UTC"))
	assert.True(t, c1.IsLeadYear())
}

func TestCarbon_IsLeadYear_False(t *testing.T) {
	c1 := CreateFromDate(2001, time.January, 1, getLocation("UTC"))
	assert.False(t, c1.IsLeadYear())
}

func TestCarbon_IsLongYear_True(t *testing.T) {
	c1 := CreateFromDate(2015, time.January, 1, getLocation("UTC"))
	assert.True(t, c1.IsLongYear())
}

func TestCarbon_IsLongYear_False(t *testing.T) {
	c1 := CreateFromDate(2016, time.January, 1, getLocation("UTC"))
	assert.False(t, c1.IsLongYear())
}

func TestCarbon_IsSameAs_True(t *testing.T) {
	c1 := CreateFromDate(2016, time.January, 17, getLocation("UTC"))
	c2 := CreateFromDate(2016, time.January, 17, getLocation("UTC"))
	assert.True(t, c1.IsSameAs(DateFormat, c2))

	c1 = CreateFromDate(2016, time.January, 17, getLocation("UTC"))
	c2 = CreateFromDate(2016, time.January, 30, getLocation("UTC"))
	assert.True(t, c1.IsSameAs("2006-01", c2))
}

func TestCarbon_IsSameAs_False(t *testing.T) {
	c1 := CreateFromDate(2016, time.January, 17, getLocation("UTC"))
	c2 := CreateFromDate(2016, time.January, 18, getLocation("UTC"))
	assert.False(t, c1.IsSameAs(DateFormat, c2))
}

func TestCarbon_IsWeekend_True(t *testing.T) {
	c1 := CreateFromDate(2020, time.May, 31, getLocation("UTC"))
	assert.True(t, c1.IsWeekend())
}

func TestCarbon_IsWeekend_False(t *testing.T) {
	c1 := CreateFromDate(2020, time.May, 28, getLocation("UTC"))
	assert.False(t, c1.IsWeekend())
}

func TestCarbon_LessThan_True(t *testing.T) {
	c1 := CreateFromDate(2016, time.January, 17, getLocation("UTC"))
	c2 := CreateFromDate(2016, time.January, 18, getLocation("UTC"))

	assert.True(t, c1.LessThan(c2))
}


func TestCarbon_LessThan_False(t *testing.T) {
	c1 := CreateFromDate(2016, time.January, 19, getLocation("UTC"))
	c2 := CreateFromDate(2016, time.January, 18, getLocation("UTC"))

	assert.False(t, c1.LessThan(c2))
}

func TestCarbon_LessThanOrEqualTo_True(t *testing.T) {
	c1 := CreateFromDate(2016, time.January, 17, getLocation("UTC"))
	c2 := CreateFromDate(2016, time.January, 18, getLocation("UTC"))

	assert.True(t, c1.LessThanOrEqualTo(c2))
}

func TestCarbon_LessThanOrEqualTo_False(t *testing.T) {
	c1 := CreateFromDate(2016, time.January, 19, getLocation("UTC"))
	c2 := CreateFromDate(2016, time.January, 18, getLocation("UTC"))

	assert.False(t, c1.LessThanOrEqualTo(c2))
}

func TestCarbon_NotEqualTo_True(t *testing.T) {
	c1 := CreateFromTime(10, 1, 0, 0, getLocation("UTC"))
	c2 := CreateFromTime(10, 0, 0, 0, getLocation("UTC"))
	assert.True(t, c1.NotEqualTo(c2))
}

func TestCarbon_NotEqualTo_False(t *testing.T) {
	c1 := CreateFromTime(10, 0, 0, 0, getLocation("UTC"))
	c2 := CreateFromTime(10, 0, 0, 0, getLocation("UTC"))
	assert.False(t, c1.NotEqualTo(c2))
}

func TestCarbon_IsSameWeek_True(t *testing.T) {
	c1 := CreateFromTime(10, 0, 0, 0, getLocation("UTC"))
	c2 := CreateFromTime(11, 0, 0, 0, getLocation("UTC"))
	assert.True(t, c1.IsSameWeek(c2))

	c1 = CreateFromDate(2020, time.May, 30, getLocation("UTC"))
	c2 = CreateFromDate(2020, time.May, 28, getLocation("UTC"))
	assert.True(t, c1.IsSameWeek(c2))
}

func TestCarbon_IsSameWeek_False(t *testing.T) {
	c1 := CreateFromDate(2020, time.May, 30, getLocation("UTC"))
	c2 := CreateFromDate(2020, time.May, 20, getLocation("UTC"))
	assert.False(t, c1.IsSameWeek(c2))
}

func TestCarbon_IsSameMonthWithSameYear_True(t *testing.T) {
	c1 := CreateFromDate(2020, time.May, 30, getLocation("UTC"))
	c2 := CreateFromDate(2020, time.May, 20, getLocation("UTC"))
	assert.True(t, c1.IsSameMonth(c2, true))
}

func TestCarbon_IsSameMonthWithSameYear_False(t *testing.T) {
	c1 := CreateFromDate(2020, time.May, 30, getLocation("UTC"))
	c2 := CreateFromDate(2020, time.April, 20, getLocation("UTC"))
	assert.False(t, c1.IsSameMonth(c2, true))
}

func TestCarbon_IsSameMonthWithNotSameYear_True(t *testing.T) {
	c1 := CreateFromDate(2020, time.May, 30, getLocation("UTC"))
	c2 := CreateFromDate(2019, time.May, 20, getLocation("UTC"))
	assert.True(t, c1.IsSameMonth(c2, false))
}

func TestCarbon_IsSameMonthWithNotSameYear_False(t *testing.T) {
	c1 := CreateFromDate(2020, time.May, 30, getLocation("UTC"))
	c2 := CreateFromDate(2019, time.April, 20, getLocation("UTC"))
	assert.False(t, c1.IsSameMonth(c2, false))
}

func TestCarbon_IsSameUnit_True(t *testing.T) {
	var c1, c2 *Carbon
	c1 = CreateFromDate(2020, time.May, 30, getLocation("UTC"))
	c2 = CreateFromDate(2020, time.May, 30, getLocation("UTC"))
	assert.True(t, c1.IsSameUnit("year", c2))

	c1 = CreateFromDate(2020, time.May, 30, getLocation("UTC"))
	c2 = CreateFromDate(2020, time.May, 30, getLocation("UTC"))
	assert.True(t, c1.IsSameUnit("day", c2))

	c1 = CreateFromTime(1, 1, 1, 1*int(time.Microsecond), getLocation("UTC"))
	c2 = CreateFromTime(1, 1, 1, 1*int(time.Microsecond), getLocation("UTC"))
	assert.True(t, c1.IsSameUnit("hour", c2))

	c1 = CreateFromTime(1, 1, 1, 1*int(time.Microsecond), getLocation("UTC"))
	c2 = CreateFromTime(1, 1, 1, 1*int(time.Microsecond), getLocation("UTC"))
	assert.True(t, c1.IsSameUnit("minute", c2))

	c1 = CreateFromTime(1, 1, 1, 1*int(time.Microsecond), getLocation("UTC"))
	c2 = CreateFromTime(1, 1, 1, 1*int(time.Microsecond), getLocation("UTC"))
	assert.True(t, c1.IsSameUnit("second", c2))

	c1 = CreateFromTime(1, 1, 1, 1*int(time.Microsecond), getLocation("UTC"))
	c2 = CreateFromTime(1, 1, 1, 1*int(time.Microsecond), getLocation("UTC"))
	assert.True(t, c1.IsSameUnit("micro", c2))

	c1 = CreateFromTime(1, 1, 1, 1*int(time.Microsecond), getLocation("UTC"))
	c2 = CreateFromTime(1, 1, 1, 1*int(time.Microsecond), getLocation("UTC"))
	assert.True(t, c1.IsSameUnit("microsecond", c2))
}

func TestCarbon_IsSameUnit_False(t *testing.T) {
	var c1, c2 *Carbon
	c1 = CreateFromDate(2019, time.May, 30, getLocation("UTC"))
	c2 = CreateFromDate(2020, time.May, 30, getLocation("UTC"))
	assert.False(t, c1.IsSameUnit("year", c2))

	c1 = CreateFromDate(2020, time.May, 30, getLocation("UTC"))
	c2 = CreateFromDate(2020, time.May, 31, getLocation("UTC"))
	assert.False(t, c1.IsSameUnit("day", c2))

	c1 = CreateFromTime(1, 1, 1, 1*int(time.Microsecond), getLocation("UTC"))
	c2 = CreateFromTime(2, 1, 1, 1*int(time.Microsecond), getLocation("UTC"))
	assert.False(t, c1.IsSameUnit("hour", c2))

	c1 = CreateFromTime(1, 1, 1, 1*int(time.Microsecond), getLocation("UTC"))
	c2 = CreateFromTime(1, 2, 1, 1*int(time.Microsecond), getLocation("UTC"))
	assert.False(t, c1.IsSameUnit("minute", c2))

	c1 = CreateFromTime(1, 1, 1, 1*int(time.Microsecond), getLocation("UTC"))
	c2 = CreateFromTime(1, 1, 2, 1*int(time.Microsecond), getLocation("UTC"))
	assert.False(t, c1.IsSameUnit("second", c2))

	c1 = CreateFromTime(1, 1, 1, 1*int(time.Microsecond), getLocation("UTC"))
	c2 = CreateFromTime(1, 1, 1, 2*int(time.Microsecond), getLocation("UTC"))
	assert.False(t, c1.IsSameUnit("micro", c2))

	c1 = CreateFromTime(1, 1, 1, 1*int(time.Microsecond), getLocation("UTC"))
	c2 = CreateFromTime(1, 1, 1, 2*int(time.Microsecond), getLocation("UTC"))
	assert.False(t, c1.IsSameUnit("microsecond", c2))
}

func TestCarbon_IsDayOfWeek_True(t *testing.T) {
	c1 := CreateFromDate(2020, time.May, 30, getLocation("UTC"))
	assert.True(t, c1.IsDayOfWeek(time.Saturday))
}

func TestCarbon_IsDayOfWeek_False(t *testing.T) {
	c1 := CreateFromDate(2020, time.May, 30, getLocation("UTC"))
	assert.False(t, c1.IsDayOfWeek(time.Sunday))
}

func TestCarbon_IsBirthday_True(t *testing.T) {
	c1 := CreateFromDate(1994, time.May, 30, getLocation("UTC"))
	birthday := CreateFromDate(1994, time.May, 30, getLocation("UTC"))
	assert.True(t, c1.IsBirthday(birthday))

	c1 = CreateFromDate(1994, time.May, 30, getLocation("UTC"))
	birthday = CreateFromDate(2020, time.May, 30, getLocation("UTC"))
	assert.True(t, c1.IsBirthday(birthday))
}

func TestCarbon_IsBirthday_False(t *testing.T) {
	c1 := CreateFromDate(1994, time.May, 30, getLocation("UTC"))
	birthday := CreateFromDate(1994, time.May, 31, getLocation("UTC"))
	assert.False(t, c1.IsBirthday(birthday))

	c1 = CreateFromDate(1994, time.May, 30, getLocation("UTC"))
	birthday = CreateFromDate(2020, time.May, 31, getLocation("UTC"))
	assert.False(t, c1.IsBirthday(birthday))
}

func TestCarbon_IsLastOfMonth_True(t *testing.T) {
	var c1 *Carbon
	c1 = CreateFromDate(1994, time.January, 31, getLocation("UTC"))
	assert.True(t, c1.IsLastOfMonth())

	c1 = CreateFromDate(1994, time.February, 28, getLocation("UTC"))
	assert.True(t, c1.IsLastOfMonth())

	c1 = CreateFromDate(2000, time.February, 29, getLocation("UTC"))
	assert.True(t, c1.IsLastOfMonth())

	c1 = CreateFromDate(2000, time.April, 30, getLocation("UTC"))
	assert.True(t, c1.IsLastOfMonth())
}

func TestCarbon_IsLastOfMonth_False(t *testing.T) {
	var c1 *Carbon
	c1 = CreateFromDate(1994, time.January, 30, getLocation("UTC"))
	assert.False(t, c1.IsLastOfMonth())

	c1 = CreateFromDate(1994, time.February, 27, getLocation("UTC"))
	assert.False(t, c1.IsLastOfMonth())

	c1 = CreateFromDate(2000, time.February, 28, getLocation("UTC"))
	assert.False(t, c1.IsLastOfMonth())

	c1 = CreateFromDate(2000, time.April, 29, getLocation("UTC"))
	assert.False(t, c1.IsLastOfMonth())
}

func TestCarbon_IsStartOfDayWithMicrosecond_True(t *testing.T) {
	var c1 *Carbon
	c1 = CreateFromTime(0, 0, 0, 0, getLocation("UTC"))
	assert.True(t, c1.IsStartOfDay(true))
}

func TestCarbon_IsStartOfDayWithoutMicrosecond_True(t *testing.T) {
	var c1 *Carbon

	c1 = CreateFromTime(0, 0, 0, 0, getLocation("UTC"))
	assert.True(t, c1.IsStartOfDay(false))

	c1 = CreateFromTime(0, 0, 0, 1*int(time.Microsecond), getLocation("UTC"))
	assert.True(t, c1.IsStartOfDay(false))
}

func TestCarbon_IsStartOfDayWithMicrosecond_False(t *testing.T) {
	var c1 *Carbon

	c1 = CreateFromTime(0, 0, 1, 0, getLocation("UTC"))
	assert.False(t, c1.IsStartOfDay(true))

	c1 = CreateFromTime(0, 0, 0, 1*int(time.Microsecond), getLocation("UTC"))
	assert.False(t, c1.IsStartOfDay(true))
}

func TestCarbon_IsStartOfDayWithoutMicrosecond_False(t *testing.T) {
	var c1 *Carbon

	c1 = CreateFromTime(0, 0, 1, 1*int(time.Microsecond), getLocation("UTC"))
	assert.False(t, c1.IsStartOfDay(false))
}

func TestCarbon_IsEndOfDayWithMicrosecond_True(t *testing.T) {
	var c1 *Carbon

	c1 = CreateFromTime(23, 59, 59, MaxNS, getLocation("UTC"))
	assert.True(t, c1.IsEndOfDay(true))
}

func TestCarbon_IsEndOfDayWithoutMicrosecond_True(t *testing.T) {
	var c1 *Carbon

	c1 = CreateFromTime(23, 59, 59, MaxNS, getLocation("UTC"))
	assert.True(t, c1.IsEndOfDay(false))

	c1 = CreateFromTime(23, 59, 59, 0, getLocation("UTC"))
	assert.True(t, c1.IsEndOfDay(false))
}

func TestCarbon_IsEndOfDayWithMicrosecond_False(t *testing.T) {
	var c1 *Carbon

	c1 = CreateFromTime(23, 59, 58, MaxNS, getLocation("UTC"))
	assert.False(t, c1.IsEndOfDay(true))

	c1 = CreateFromTime(23, 59, 59, 0, getLocation("UTC"))
	assert.False(t, c1.IsEndOfDay(true))
}

func TestCarbon_IsEndOfDayWithoutMicrosecond_False(t *testing.T) {
	var c1 *Carbon

	c1 = CreateFromTime(23, 59, 58, 0, getLocation("UTC"))
	assert.False(t, c1.IsEndOfDay(true))
}

var (
	midAt = 12
)

func TestCarbon_isMidday_True(t *testing.T) {
	var c1 *Carbon
	c1 = CreateFromTime(midAt, 0, 0, 0, getLocation("UTC"))
	assert.True(t, c1.isMidday())
}

func TestCarbon_isMidday_False(t *testing.T) {
	var c1 *Carbon
	c1 = CreateFromTime(midAt, 1, 0, 0, getLocation("UTC"))
	assert.False(t, c1.isMidday())
}
