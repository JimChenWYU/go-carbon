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
	return c.SetDate(c.Year(), 1, 1).StartOfDay()
}

/**
 * Resets the date to end of the year and time to 23:59:59.999999
 */
func (c *Carbon) EndOfYear() *Carbon {
	return c.SetDate(c.Year(), 12, 31).EndOfDay()
}
