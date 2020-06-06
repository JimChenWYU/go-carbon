package carbon

import (
	"math"
	"time"
)

// Get the difference in nanoseconds using timestamps.
func (c *Carbon) DiffInRealNanoseconds(v interface{}, absolute bool) int64 {
	diff := c.resolveCarbon(v).UnixNano() - c.UnixNano()
	if absolute && diff < 0 {
		diff = -diff
	}
	return diff
}

// Get the difference in microseconds using timestamps.
func (c *Carbon) DiffInRealMicroseconds(v interface{}, absolute bool) int64 {
	diff := c.DiffInRealNanoseconds(v, absolute)
	return diff / int64(time.Microsecond)
}

// Get the difference in milliseconds rounded down using timestamps.
func (c *Carbon) DiffInRealMilliseconds(v interface{}, absolute bool) int64 {
	diff := c.DiffInRealNanoseconds(v, absolute)
	return diff / int64(time.Millisecond)
}

// Get the difference in seconds using timestamps.
func (c *Carbon) DiffInRealSeconds(v interface{}, absolute bool) int64 {
	diff := c.resolveCarbon(v).Timestamp() - c.Timestamp()
	if absolute && diff < 0 {
		diff = -diff
	}
	return diff
}

// Get the difference in nanoseconds.
func (c *Carbon) DiffInNanoseconds(v interface{}, absolute bool) int64 {
    diff := c.resolveCarbon(v).Sub(c.Time)
	if absolute && diff < 0 {
		diff = -diff
	}

	return int64(diff)
}

// Get the difference in microseconds.
func (c *Carbon) DiffInMicroseconds(v interface{}, absolute bool) int64 {
	return c.DiffInNanoseconds(v, absolute) / int64(time.Microsecond)
}

// Get the difference in milliseconds rounded down.
func (c *Carbon) DiffInMilliseconds(v interface{}, absolute bool) int64 {
	return c.DiffInNanoseconds(v, absolute) / int64(time.Millisecond)
}

// Get the difference in seconds rounded down.
func (c *Carbon) DiffInSeconds(v interface{}, absolute bool) int64 {
	return c.DiffInNanoseconds(v, absolute) / int64(time.Second)
}

// Get the difference in minutes rounded down.
func (c *Carbon) DiffInMinutes(v interface{}, absolute bool) int64 {
    return c.DiffInSeconds(v, absolute) / SecondsPerMinute
}

// Get the difference in minutes rounded down using timestamps.
func (c *Carbon) DiffInRealMinutes(v interface{}, absolute bool) int64 {
    return c.DiffInRealSeconds(v, absolute) / SecondsPerMinute
}

// Get the difference in hours rounded down.
func (c *Carbon) DiffInHours(v interface{}, absolute bool) int64 {
    return c.DiffInMinutes(v, absolute) / MinutesPerHour
}

// Get the difference in hours rounded down using timestamps.
func (c *Carbon) DiffInRealHours(v interface{}, absolute bool) int64 {
	return c.DiffInRealMinutes(v, absolute) / MinutesPerHour
}

// Get the difference in days rounded down.
func (c *Carbon) DiffInDays(v interface{}, absolute bool) int64 {
	return c.DiffInHours(v, absolute) / HoursPerDay
}

// Get the difference in weeks rounded down.
func (c *Carbon) DiffInWeeks(v interface{}, absolute bool) int64 {
	return c.DiffInDays(v, absolute) / DaysPerWeek
}

// DiffInYears returns the difference in years
func (c *Carbon) DiffInYears(v interface{}, absolute bool) int64 {
	resolve := c.resolveCarbon(v)
	if c.Year() == resolve.Year() {
		return 0
	}

	start := NewCarbon(c.Time)
	end := NewCarbon(resolve.Time)
	swap := false
	if 	end.UnixNano() < start.UnixNano() {
		start, end = end, start
		swap = true
	}

	diff := end.Year() - start.Year() - 1
	start.SetYears(end.Year())
	if start.UnixNano() <= end.UnixNano() {
		diff++
	}
	if swap {
		diff = -diff
	}
	if absolute && diff < 0 {
		diff = -diff
	}
	return int64(diff)
}

// DiffInMonths returns the difference in months
// TODO: Add more tests
func (c *Carbon) DiffInMonths(v interface{}, absolute bool) int64 {
	resolve := c.resolveCarbon(v)
	originalCopy := c.Copy()
	resolveCopy := resolve.Copy()
	diffMonths := 0
	if resolveCopy.Year() == originalCopy.Year() && resolveCopy.Month() != originalCopy.Month() {
		remainingTime := int(originalCopy.DiffInHours(resolveCopy, false))
		// less than 0 when resolve < original
		if remainingTime < 0 {
			diffMonths = int(math.Ceil(float64(remainingTime) / float64(resolveCopy.DaysInMonth()*HoursPerDay)))
		} else {
			diffMonths = int(math.Floor(float64(remainingTime) / float64(originalCopy.DaysInMonth()*HoursPerDay)))
		}
	}
	if resolveCopy.Year() < originalCopy.Year() {
		diffMonths = int(MonthsPerYear - resolveCopy.Month() + originalCopy.Month() - 1)
		if originalCopy.hasRemainingOneMonthHour(resolveCopy) {
			diffMonths++
		}
		if resolve.LessThan(c) {
			diffMonths = -diffMonths
		}
	}
	if resolveCopy.Year() > originalCopy.Year() {
		diffMonths = int(MonthsPerYear - originalCopy.Month() + resolveCopy.Month() - 1)
		if resolveCopy.hasRemainingOneMonthHour(originalCopy) {
			diffMonths++
		}
		if resolve.LessThan(c) {
			diffMonths = -diffMonths
		}
	}
	diffMonths = diffMonths % MonthsPerYear
	if absolute && diffMonths < 0 {
		diffMonths = -diffMonths
	}
	return c.DiffInYears(v, absolute) * MonthsPerYear + int64(diffMonths)
}

func (c *Carbon) hasRemainingOneMonthHour(resolve *Carbon) bool {
	totalHour := int64(c.DaysInMonth() * HoursPerDay)
	diff := c.Copy().StartOfMonth().DiffInHours(c, false)
	remainHour := totalHour - diff
	spendHour := resolve.Copy().StartOfMonth().DiffInHours(resolve, false)
	if c.Month() == resolve.Month() {
		return diff == spendHour
	}

	return remainHour + spendHour >= totalHour
}