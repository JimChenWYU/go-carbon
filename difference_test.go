package carbon

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCarbon_DiffInRealNanoseconds_Absolute(t *testing.T) {
	c1 := Create(2020, time.May, 30, 0, 0, 0, 10, getLocation("UTC"))
	c2 := Create(2020, time.May, 30, 0, 0, 0, 0, getLocation("UTC"))

	assert.EqualValues(t, 10, c1.DiffInRealNanoseconds(c2, true))
	assert.EqualValues(t, 10, c2.DiffInRealNanoseconds(c1, true))
}

func TestCarbon_DiffInRealNanoseconds_NotAbsolute(t *testing.T) {
	c1 := Create(2020, time.May, 30, 0, 0, 0, 10, getLocation("UTC"))
	c2 := Create(2020, time.May, 30, 0, 0, 0, 0, getLocation("UTC"))

	assert.EqualValues(t, -10, c1.DiffInRealNanoseconds(c2, false))
	assert.EqualValues(t, 10, c2.DiffInRealNanoseconds(c1, false))
}

func TestCarbon_DiffInRealMicroseconds_Absolute(t *testing.T) {
	c1 := Create(2020, time.May, 30, 0, 0, 0, 10 * int(time.Microsecond), getLocation("UTC"))
	c2 := Create(2020, time.May, 30, 0, 0, 0, 0, getLocation("UTC"))

	assert.EqualValues(t, 10, c1.DiffInRealMicroseconds(c2, true))
	assert.EqualValues(t, 10, c2.DiffInRealMicroseconds(c1, true))
}

func TestCarbon_DiffInRealMicroseconds_NotAbsolute(t *testing.T) {
	c1 := Create(2020, time.May, 30, 0, 0, 0, 10 * int(time.Microsecond), getLocation("UTC"))
	c2 := Create(2020, time.May, 30, 0, 0, 0, 0, getLocation("UTC"))

	assert.EqualValues(t, -10, c1.DiffInRealMicroseconds(c2, false))
	assert.EqualValues(t, 10, c2.DiffInRealMicroseconds(c1, false))
}

func TestCarbon_DiffInRealMilliseconds_Absolute(t *testing.T) {
	c1 := Create(2020, time.May, 30, 0, 0, 0, 10 * int(time.Millisecond), getLocation("UTC"))
	c2 := Create(2020, time.May, 30, 0, 0, 0, 0, getLocation("UTC"))

	assert.EqualValues(t, 10, c1.DiffInRealMilliseconds(c2, true))
	assert.EqualValues(t, 10, c2.DiffInRealMilliseconds(c1, true))
}

func TestCarbon_DiffInRealMilliseconds_NotAbsolute(t *testing.T) {
	c1 := Create(2020, time.May, 30, 0, 0, 0, 10 * int(time.Millisecond), getLocation("UTC"))
	c2 := Create(2020, time.May, 30, 0, 0, 0, 0, getLocation("UTC"))

	assert.EqualValues(t, -10, c1.DiffInRealMilliseconds(c2, false))
	assert.EqualValues(t, 10, c2.DiffInRealMilliseconds(c1, false))
}

func TestCarbon_DiffInRealSeconds_Absolute(t *testing.T) {
	c1 := Create(2020, time.May, 30, 0, 0, 10, 0, getLocation("UTC"))
	c2 := Create(2020, time.May, 30, 0, 0, 0, 0, getLocation("UTC"))

	assert.EqualValues(t, 10, c1.DiffInRealSeconds(c2, true))
	assert.EqualValues(t, 10, c2.DiffInRealSeconds(c1, true))
}

func TestCarbon_DiffInRealSeconds_NotAbsolute(t *testing.T) {
	c1 := Create(2020, time.May, 30, 0, 0, 10, 0, getLocation("UTC"))
	c2 := Create(2020, time.May, 30, 0, 0, 0, 0, getLocation("UTC"))

	assert.EqualValues(t, -10, c1.DiffInRealSeconds(c2, false))
	assert.EqualValues(t, 10, c2.DiffInRealSeconds(c1, false))
}

func TestCarbon_DiffInNanoseconds_Absolute(t *testing.T) {
	var c1, c2 *Carbon
	c1 = Create(2020, time.May, 30, 0, 0, 0, 10, getLocation("UTC"))
	c2 = Create(2020, time.May, 30, 0, 0, 0, 0, getLocation("UTC"))
	assert.EqualValues(t, 10, c1.DiffInNanoseconds(c2, true))
	assert.EqualValues(t, 10, c2.DiffInNanoseconds(c1, true))

	c1 = Create(2020, time.May, 30, 0, 0, 10, 0, getLocation("UTC"))
	c2 = Create(2020, time.May, 30, 0, 0, 0, 0, getLocation("UTC"))
	assert.EqualValues(t, 10, c1.DiffInSeconds(c2, true))
	assert.EqualValues(t, 10, c2.DiffInSeconds(c1, true))


	c1 = Create(2020, time.May, 30, 0, 1, 0, 0, getLocation("UTC"))
	c2 = Create(2020, time.May, 30, 0, 0, 0, 0, getLocation("UTC"))
	assert.EqualValues(t, 1, c1.DiffInMinutes(c2, true))
	assert.EqualValues(t, 1, c2.DiffInMinutes(c1, true))

	c1 = Create(2020, time.May, 30, 1, 0, 0, 0, getLocation("UTC"))
	c2 = Create(2020, time.May, 30, 0, 0, 0, 0, getLocation("UTC"))
	assert.EqualValues(t, 1, c1.DiffInHours(c2, true))
	assert.EqualValues(t, 1, c2.DiffInHours(c1, true))
}

func TestCarbon_DiffInNanoseconds_NotAbsolute(t *testing.T) {
	c1 := Create(2020, time.May, 30, 0, 0, 10 * int(time.Nanosecond), 0, getLocation("UTC"))
	c2 := Create(2020, time.May, 30, 0, 0, 0, 0, getLocation("UTC"))

	assert.EqualValues(t, -10, c1.DiffInSeconds(c2, false))
	assert.EqualValues(t, 10, c2.DiffInSeconds(c1, false))
}

func TestCarbon_DiffInMicroseconds_Absolute(t *testing.T) {
	c1 := Create(2020, time.May, 30, 0, 0, 0, 10 * int(time.Microsecond), getLocation("UTC"))
	c2 := Create(2020, time.May, 30, 0, 0, 0, 0, getLocation("UTC"))

	assert.EqualValues(t, 10, c1.DiffInMicroseconds(c2, true))
	assert.EqualValues(t, 10 , c2.DiffInMicroseconds(c1, true))
}

func TestCarbon_DiffInMicroseconds_NotAbsolute(t *testing.T) {
	c1 := Create(2020, time.May, 30, 0, 0, 0, 10 * int(time.Microsecond), getLocation("UTC"))
	c2 := Create(2020, time.May, 30, 0, 0, 0, 0, getLocation("UTC"))

	assert.EqualValues(t, -10, c1.DiffInMicroseconds(c2, false))
	assert.EqualValues(t, 10, c2.DiffInMicroseconds(c1, false))
}

func TestCarbon_DiffInMilliseconds_Absolute(t *testing.T) {
	c1 := Create(2020, time.May, 30, 0, 0, 0, 10 * int(time.Millisecond), getLocation("UTC"))
	c2 := Create(2020, time.May, 30, 0, 0, 0, 0, getLocation("UTC"))

	assert.EqualValues(t, 10, c1.DiffInMilliseconds(c2, true))
	assert.EqualValues(t, 10 , c2.DiffInMilliseconds(c1, true))
}

func TestCarbon_DiffInMilliseconds_NotAbsolute(t *testing.T) {
	c1 := Create(2020, time.May, 30, 0, 0, 0, 10 * int(time.Millisecond), getLocation("UTC"))
	c2 := Create(2020, time.May, 30, 0, 0, 0, 0, getLocation("UTC"))

	assert.EqualValues(t, -10, c1.DiffInMilliseconds(c2, false))
	assert.EqualValues(t, 10 , c2.DiffInMilliseconds(c1, false))
}

func TestCarbon_DiffInSeconds_Absolute(t *testing.T) {
	c1 := Create(2020, time.May, 30, 0, 0, 10, 0, getLocation("UTC"))
	c2 := Create(2020, time.May, 30, 0, 0, 0, 0, getLocation("UTC"))

	assert.EqualValues(t, 10, c1.DiffInSeconds(c2, true))
	assert.EqualValues(t, 10, c2.DiffInSeconds(c1, true))
}

func TestCarbon_DiffInSeconds_NotAbsolute(t *testing.T) {
	c1 := Create(2020, time.May, 30, 0, 0, 10, 0, getLocation("UTC"))
	c2 := Create(2020, time.May, 30, 0, 0, 0, 0, getLocation("UTC"))

	assert.EqualValues(t, -10, c1.DiffInSeconds(c2, false))
	assert.EqualValues(t, 10, c2.DiffInSeconds(c1, false))
}

func TestCarbon_DiffInYears_Absolute(t *testing.T) {
	var c1, c2 *Carbon
	c1 = Create(2020, time.May, 30, 0, 0, 0, 0, getLocation("UTC"))
	c2 = Create(2020, time.May, 30, 0, 0, 0, 0, getLocation("UTC"))
	assert.EqualValues(t, 0, c1.DiffInYears(c2, true))
	assert.EqualValues(t, 0, c2.DiffInYears(c1, true))

	c1 = Create(2019, time.May, 30, 0, 0, 0, 0, getLocation("UTC"))
	c2 = Create(2020, time.May, 30, 0, 0, 0, 0, getLocation("UTC"))
	assert.EqualValues(t, 1, c1.DiffInYears(c2, true))
	assert.EqualValues(t, 1, c2.DiffInYears(c1, true))

	c1 = Create(2019, time.May, 30, 0, 0, 0, 0, getLocation("UTC"))
	c2 = Create(2020, time.May, 30, 1, 0, 0, 0, getLocation("UTC"))
	assert.EqualValues(t, 1, c1.DiffInYears(c2, true))
	assert.EqualValues(t, 1, c2.DiffInYears(c1, true))

	c1 = Create(2019, time.May, 30, 1, 0, 0, 0, getLocation("UTC"))
	c2 = Create(2020, time.May, 30, 0, 0, 0, 0, getLocation("UTC"))

	assert.EqualValues(t, 0, c1.DiffInYears(c2, true))
	assert.EqualValues(t, 0, c2.DiffInYears(c1, true))
}

func TestCarbon_DiffInYears_NotAbsolute(t *testing.T) {
	var c1, c2 *Carbon
	c1 = Create(2020, time.May, 30, 0, 0, 0, 0, getLocation("UTC"))
	c2 = Create(2020, time.May, 30, 0, 0, 0, 0, getLocation("UTC"))
	assert.EqualValues(t, 0, c1.DiffInYears(c2, false))
	assert.EqualValues(t, 0, c2.DiffInYears(c1, false))

	c1 = Create(2019, time.May, 30, 0, 0, 0, 0, getLocation("UTC"))
	c2 = Create(2020, time.May, 30, 0, 0, 0, 0, getLocation("UTC"))
	assert.EqualValues(t, 1, c1.DiffInYears(c2, false))
	assert.EqualValues(t, -1, c2.DiffInYears(c1, false))

	c1 = Create(2019, time.May, 30, 0, 0, 0, 0, getLocation("UTC"))
	c2 = Create(2020, time.May, 30, 1, 0, 0, 0, getLocation("UTC"))
	assert.EqualValues(t, 1, c1.DiffInYears(c2, false))
	assert.EqualValues(t, -1, c2.DiffInYears(c1, false))

	c1 = Create(2019, time.May, 30, 1, 0, 0, 0, getLocation("UTC"))
	c2 = Create(2020, time.May, 30, 0, 0, 0, 0, getLocation("UTC"))

	assert.EqualValues(t, 0, c1.DiffInYears(c2, false))
	assert.EqualValues(t, 0, c2.DiffInYears(c1, false))
}

func TestCarbon_DiffInMonths_Absolute(t *testing.T) {
	var c1, c2 *Carbon

	c1 = Create(2020, time.May, 30, 0, 0, 0, 0, getLocation("UTC"))
	c2 = Create(2020, time.April, 30, 0, 0, 0, 0, getLocation("UTC"))
	assert.EqualValues(t, 1, c1.DiffInMonths(c2, true))
	assert.EqualValues(t, 1, c2.DiffInMonths(c1, true))

	c1 = Create(2020, time.May, 30, 0, 0, 0, 0, getLocation("UTC"))
	c2 = Create(2020, time.May, 20, 0, 0, 0, 0, getLocation("UTC"))
	assert.EqualValues(t, 0, c1.DiffInMonths(c2, true))
	assert.EqualValues(t, 0, c2.DiffInMonths(c1, true))

	c1 = Create(2020, time.May, 20, 0, 0, 0, 0, getLocation("UTC"))
	c2 = Create(2020, time.April, 30, 0, 0, 0, 0, getLocation("UTC"))
	assert.EqualValues(t, 0, c1.DiffInMonths(c2, true))
	assert.EqualValues(t, 0, c2.DiffInMonths(c1, true))

	c1 = Create(2020, time.May, 20, 0, 0, 0, 0, getLocation("UTC"))
	c2 = Create(2023, time.May, 20, 0, 0, 0, 0, getLocation("UTC"))
	assert.EqualValues(t, 3*MonthsPerYear, c1.DiffInMonths(c2, true))
	assert.EqualValues(t, 3*MonthsPerYear, c2.DiffInMonths(c1, true))

	c1 = Create(2020, time.May, 20, 0, 0, 0, 0, getLocation("UTC"))
	c2 = Create(2023, time.April, 20, 0, 0, 0, 0, getLocation("UTC"))
	assert.EqualValues(t, 3*MonthsPerYear-1, c1.DiffInMonths(c2, true))
	assert.EqualValues(t, 3*MonthsPerYear-1, c2.DiffInMonths(c1, true))

	c1 = Create(2020, time.May, 20, 0, 0, 0, 0, getLocation("UTC"))
	c2 = Create(2023, time.May, 10, 0, 0, 0, 0, getLocation("UTC"))
	assert.EqualValues(t, 3*MonthsPerYear-1, c1.DiffInMonths(c2, true))
	assert.EqualValues(t, 3*MonthsPerYear-1, c2.DiffInMonths(c1, true))

	c1 = Create(2020, time.May, 20, 1, 0, 0, 0, getLocation("UTC"))
	c2 = Create(2023, time.May, 20, 0, 0, 0, 0, getLocation("UTC"))
	assert.EqualValues(t, 3*MonthsPerYear-1, c1.DiffInMonths(c2, true))
	assert.EqualValues(t, 3*MonthsPerYear-1, c2.DiffInMonths(c1, true))
}

func TestCarbon_DiffInMonths_NotAbsolute(t *testing.T) {
	var c1, c2 *Carbon

	c1 = Create(2020, time.May, 30, 0, 0, 0, 0, getLocation("UTC"))
	c2 = Create(2020, time.April, 30, 0, 0, 0, 0, getLocation("UTC"))
	assert.EqualValues(t, -1, c1.DiffInMonths(c2, false))
	assert.EqualValues(t, 1, c2.DiffInMonths(c1, false))

	c1 = Create(2020, time.May, 30, 0, 0, 0, 0, getLocation("UTC"))
	c2 = Create(2020, time.May, 20, 0, 0, 0, 0, getLocation("UTC"))
	assert.EqualValues(t, 0, c1.DiffInMonths(c2, false))
	assert.EqualValues(t, 0, c2.DiffInMonths(c1, false))

	c1 = Create(2020, time.May, 20, 0, 0, 0, 0, getLocation("UTC"))
	c2 = Create(2020, time.April, 30, 0, 0, 0, 0, getLocation("UTC"))
	assert.EqualValues(t, 0, c1.DiffInMonths(c2, false))
	assert.EqualValues(t, 0, c2.DiffInMonths(c1, false))

	c1 = Create(2020, time.May, 20, 0, 0, 0, 0, getLocation("UTC"))
	c2 = Create(2023, time.May, 20, 0, 0, 0, 0, getLocation("UTC"))
	assert.EqualValues(t, 3*MonthsPerYear, c1.DiffInMonths(c2, false))
	assert.EqualValues(t, -3*MonthsPerYear, c2.DiffInMonths(c1, false))

	c1 = Create(2020, time.May, 20, 0, 0, 0, 0, getLocation("UTC"))
	c2 = Create(2023, time.April, 20, 0, 0, 0, 0, getLocation("UTC"))
	assert.EqualValues(t, 3*MonthsPerYear-1, c1.DiffInMonths(c2, false))
	assert.EqualValues(t, -(3*MonthsPerYear-1), c2.DiffInMonths(c1, false))

	c1 = Create(2020, time.May, 20, 0, 0, 0, 0, getLocation("UTC"))
	c2 = Create(2023, time.May, 10, 0, 0, 0, 0, getLocation("UTC"))
	assert.EqualValues(t, 3*MonthsPerYear-1, c1.DiffInMonths(c2, false))
	assert.EqualValues(t, -(3*MonthsPerYear-1), c2.DiffInMonths(c1, false))
}

func TestCarbon_DiffInMonths_Positive(t *testing.T) {
	c1 := CreateFromDate(2000, time.January, 1, getLocation("UTC"))
	assert.EqualValues(t, 13, c1.DiffInMonths(c1.Copy().AddYears(1).AddMonths(1), true))
}

func TestCarbon_DiffInMonths_NegativeWithSign(t *testing.T) {
	c1 := CreateFromDate(2000, time.January, 1, getLocation("UTC"))
	assert.EqualValues(t, -11, c1.DiffInMonths(c1.Copy().SubYears(1).AddMonths(1), false))
}

func TestCarbon_DiffInMonths_NegativeNoSign(t *testing.T) {
	c1 := CreateFromDate(2000, time.January, 1, getLocation("UTC"))
	assert.EqualValues(t, 11, c1.DiffInMonths(c1.Copy().SubYears(1).AddMonths(1), true))
}

func TestCarbon_DiffInMonths_VsDefaultNow(t *testing.T) {
	wrapWithDefaultNowTestNow(func() {
		assert.EqualValues(t, 12, Now().SubYears(1).DiffInMonths(Now(), true))
	})
}

func TestCarbon_DiffInMonths_EnsureIsTruncated(t *testing.T) {
	c1 := CreateFromDate(2000, time.January, 1, getLocation("UTC"))
	assert.EqualValues(t, 1, c1.DiffInMonths(c1.Copy().AddMonths(1).AddDays(16), true))
}

func TestCarbon_DiffInQuarters_Positive(t *testing.T) {
	c1 := CreateFromDate(2000, time.January, 1, getLocation("UTC"))
	assert.EqualValues(t, 1, c1.DiffInQuarters(c1.Copy().AddQuarters(1).AddDays(1), true))
}

func TestCarbon_DiffInQuarters_NegativeWithSign(t *testing.T) {
	c1 := CreateFromDate(2000, time.January, 1, getLocation("UTC"))
	assert.EqualValues(t, -4, c1.DiffInQuarters(c1.Copy().SubQuarters(4), false))
}

func TestCarbon_DiffInQuarters_NegativeWithNoSign(t *testing.T) {
	c1 := CreateFromDate(2000, time.January, 1, getLocation("UTC"))
	assert.EqualValues(t, 4, c1.DiffInQuarters(c1.Copy().SubQuarters(4), true))
}

func TestCarbon_DiffInQuarters_VsDefaultNow(t *testing.T) {
	wrapWithDefaultNowTestNow(func() {
		assert.EqualValues(t, 4, Now().SubYears(1).DiffInQuarters(Now(), true))
	})
}
