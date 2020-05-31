package carbon

import (
	"math"
	"time"
)

/**
 * Resets the time to 00:00:00 start of day
 */
func (c *Carbon) StartOfDay() *Carbon {
	return c.SetTime(0, 0, 0, 0)
}

/**
 * Resets the time to 23:59:59.999999 end of day
 */
func (c *Carbon) EndOfDay() *Carbon {
	return c.SetTime(HoursPerDay-1, MinutesPerHour-1, SecondsPerMinute-1, NanosecondsPerSecond-1)
}

/**
 * Resets the date to the first day of the month and the time to 00:00:00
 */
func (c *Carbon) StartOfMonth() *Carbon {
	return c.SetDate(c.Year(), c.Month(), 1).StartOfDay()
}

/**
 * Resets the date to end of the month and time to 23:59:59.999999
 */
func (c *Carbon) EndOfMonth() *Carbon {
	c.StartOfMonth()
	c.Time = c.AddDate(0, 1, -1)
	return c.EndOfDay()
}

/**
 * Resets the date to the first day of the quarter and the time to 00:00:00
 */
func (c *Carbon) StartOfQuarter() *Carbon {
	quarter := math.Ceil(float64(c.Month()) / float64(MonthsPerQuarter))
	month := (quarter-1)*MonthsPerQuarter + 1
	return c.SetDate(c.Year(), time.Month(month), 1).StartOfDay()
}

/**
 * Resets the date to end of the quarter and time to 23:59:59.999999
 */
func (c *Carbon) EndOfQuarter() *Carbon {
	return c.StartOfQuarter().AddMonths(MonthsPerQuarter - 1).EndOfMonth()
}

/**
 * Resets the date to the first day of the year and the time to 00:00:00
 */
func (c *Carbon) StartOfYear() *Carbon {
	return c.SetDate(c.Year(), time.January, 1).StartOfDay()
}

/**
 * Resets the date to end of the year and time to 23:59:59.999999
 */
func (c *Carbon) EndOfYear() *Carbon {
	return c.SetDate(c.Year(), time.December, 31).EndOfDay()
}

// Resets the date to the first day of the decade and the time to 00:00:00
func (c *Carbon) StartOfDecade() *Carbon {
	year := c.Year() - c.Year()%YearsPerDecade
	return c.SetDate(year, time.January, 1).StartOfDay()
}

// Resets the date to end of the decade and time to 23:59:59.999999
func (c *Carbon) EndOfDecade() *Carbon {
	year := c.Year() - c.Year()%YearsPerDecade + YearsPerDecade - 1
	return c.SetDate(year, time.December, 31).EndOfDay()
}

// Resets the date to the first day of the century and the time to 00:00:00
func (c *Carbon) StartOfCentury() *Carbon {
	year := c.Year() - (c.Year()-1)%YearsPerCentury
	return c.SetDate(year, time.January, 1).StartOfDay()
}

// Resets the date to end of the century and time to 23:59:59.999999
func (c *Carbon) EndOfCentury() *Carbon {
	year := c.Year() - 1 - (c.Year()-1)%YearsPerCentury + YearsPerCentury
	return c.SetDate(year, time.December, 31).EndOfDay()
}

// Resets the date to the first day of the millennium and the time to 00:00:00
func (c *Carbon) StartOfMillennium() *Carbon {
	year := c.Year() - (c.Year()-1)%YearsPerMillennium
	return c.SetDate(year, time.January, 1)
}

// Resets the date to end of the millennium and time to 23:59:59.999999
func (c *Carbon) EndOfMillennium() *Carbon {
	year := c.Year() - 1 - (c.Year()-1)%YearsPerMillennium + YearsPerMillennium
	return c.SetDate(year, time.December, 31)
}

// Resets the date to the first day of week (defined in $weekStartsAt) and the time to 00:00:00
func (c *Carbon) StartOfWeek(weekStartAt time.Weekday) *Carbon {
	sub := int((DaysPerWeek + c.Weekday() - weekStartAt) % DaysPerWeek)
	return c.SubDays(sub).StartOfDay()
}

// Resets the date to end of week (defined in $weekEndsAt) and time to 23:59:59.999999
func (c *Carbon) EndOfWeek(weekEndAt time.Weekday) *Carbon {
	add := int((DaysPerWeek - c.Weekday() + weekEndAt) % DaysPerWeek)
	return c.AddDays(add).EndOfDay()
}

// Modify to start of current hour, minutes and seconds become 0
func (c *Carbon) StartOfHour() *Carbon {
	return c.SetTime(c.Hour(), 0, 0, 0)
}

// Modify to end of current hour, minutes and seconds become 59
func (c *Carbon) EndOfHour() *Carbon {
	return c.SetTime(c.Hour(), MinutesPerHour-1, SecondsPerMinute-1, NanosecondsPerSecond-1)
}

// Modify to start of current minute, seconds become 0
func (c *Carbon) StartOfMinute() *Carbon {
	return c.SetTime(c.Hour(), c.Minute(), 0, 0)
}

// Modify to end of current minute, seconds become 59
func (c *Carbon) EndOfMinute() *Carbon {
	return c.SetTime(c.Hour(), c.Minute(), SecondsPerMinute-1, NanosecondsPerSecond-1)
}

// Modify to start of current second, microseconds become 0
func (c *Carbon) StartOfSecond() *Carbon {
	return c.SetTime(c.Hour(), c.Minute(), c.Second(), 0)
}

// Modify to end of current second, microseconds become 999999
func (c *Carbon) EndOfSecond() *Carbon {
	return c.SetTime(c.Hour(), c.Minute(), c.Second(), NanosecondsPerSecond-1)
}
