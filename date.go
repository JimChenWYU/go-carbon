package carbon

import (
	"errors"
	"time"
)

// Set the instance's timezone from a string or object.
func (c *Carbon) SetTimezone(l *time.Location) *Carbon {
	c.Time = time.Date(c.Year(), c.Month(), c.Day(), c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), l)
	return c
}

// Get the TimeZone associated with the Carbon instance.
func (c *Carbon) Timezone() *time.Location {
	return c.Location()
}

// Add microseconds to the instance.
func (c *Carbon) AddMicro(value int) *Carbon {
	return c.AddUnit("micro", value)
}

// Add milliseconds to the instance.
func (c *Carbon) AddMilli(value int) *Carbon {
	//return c.AddMicro(value * int(time.Millisecond))
	return c.AddUnit("milli", value)
}

// Add seconds to the instance.
func (c *Carbon) AddSeconds(value int) *Carbon {
	//return c.AddMicro(value * int(time.Second))
	return c.AddUnit("second", value)
}

// Add minutes to the instance.
func (c *Carbon) AddMinutes(value int) *Carbon {
	//return c.AddMicro(value * int(time.Minute))
	return c.AddUnit("minute", value)
}

// Add hours to the instance.
func (c *Carbon) AddHours(value int) *Carbon {
	//return c.AddMicro(value * int(time.Hour))
	return c.AddUnit("hour", value)
}

// Add days to the instance.
func (c *Carbon) AddDays(value int) *Carbon {
	return c.AddUnit("day", value)
}

// Add weeks to the instance.
func (c *Carbon) AddWeeks(value int) *Carbon {
	return c.AddDays(value * DaysPerWeek)
}

// Add months to the instance.
func (c *Carbon) AddMonths(value int) *Carbon {
	return c.AddUnit("month", value)
}

// Add years to the instance.
func (c *Carbon) AddYears(value int) *Carbon {
	return c.AddUnit("year", value)
}

// Sub milliseconds to the instance.
func (c *Carbon) SubMilli(value int) *Carbon {
	return c.SubUnit("milli", value)
}

// Sub microseconds to the instance.
func (c *Carbon) SubMicro(value int) *Carbon {
	return c.SubUnit("micro", value)
}

// Sub seconds to the instance.
func (c *Carbon) SubSeconds(value int) *Carbon {
	return c.SubUnit("second", value)
}

// Sub minutes to the instance.
func (c *Carbon) SubMinutes(value int) *Carbon {
	return c.SubUnit("minute", value)
}

// Sub hours to the instance.
func (c *Carbon) SubHours(value int) *Carbon {
	return c.SubUnit("hour", value)
}

// Sub days to the instance.
func (c *Carbon) SubDays(value int) *Carbon {
	return c.SubUnit("day", value)
}

// Sub weeks to the instance.
func (c *Carbon) SubWeeks(value int) *Carbon {
	return c.SubDays(value * DaysPerWeek)
}

// Sub months to the instance.
func (c *Carbon) SubMonths(value int) *Carbon {
	return c.SubUnit("month", value)
}

// Sub years to the instance.
func (c *Carbon) SubYears(value int) *Carbon {
	return c.SubUnit("year", value)
}

// Set the date with gregorian year, month and day numbers.
func (c *Carbon) SetDate(year int, month time.Month, day int) *Carbon {
	c.Time = time.Date(year, month, day, c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Location())
	return c
}

// Resets the current time of the DateTime object to a different time.
func (c *Carbon) SetTime(hour, minute, second, nanosecond int) *Carbon {
	c.Time = time.Date(c.Year(), c.Month(), c.Day(), hour, minute, second, nanosecond, c.Location())
	return c
}

// Sets the date and time based on a Unix timestamp.
func (c *Carbon) SetTimestamp(t int64) *Carbon {
	c.Time = time.Unix(t, 0)
	return c
}

// Set current instance nanoseconds to the given value.
func (c *Carbon) SetNanoSeconds(ns int) *Carbon {
	c.Time = time.Date(c.Year(), c.Month(), c.Day(), c.Hour(), c.Minute(), c.Second(), ns, c.Location())
	return c
}

// Set current instance seconds to the given value.
func (c *Carbon) SetSeconds(s int) *Carbon {
	c.Time = time.Date(c.Year(), c.Month(), c.Day(), c.Hour(), c.Minute(), s, c.Nanosecond(), c.Location())
	return c
}

// Set current instance minutes to the given value.
func (c *Carbon) SetMinutes(m int) *Carbon {
	c.Time = time.Date(c.Year(), c.Month(), c.Day(), c.Hour(), m, c.Second(), c.Nanosecond(), c.Location())
	return c
}

// Set current instance hours to the given value.
func (c *Carbon) SetHours(h int) *Carbon {
	c.Time = time.Date(c.Year(), c.Month(), c.Day(), h, c.Minute(), c.Second(), c.Nanosecond(), c.Location())
	return c
}

// Set current instance days to the given value.
func (c *Carbon) SetDays(d int) *Carbon {
	c.Time = time.Date(c.Year(), c.Month(), d, c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Location())
	return c
}

// Set current instance months to the given value.
func (c *Carbon) SetMonths(m time.Month) *Carbon {
	c.Time = time.Date(c.Year(), m, c.Day(), c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Location())
	return c
}

// Set current instance years to the given value.
func (c *Carbon) SetYears(y int) *Carbon {
	c.Time = time.Date(y, c.Month(), c.Day(), c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Location())
	return c
}

// Add given units to the current instance.
func (c *Carbon) AddUnit(which string, value int) *Carbon {
	switch which {
	case "micro":
		c.Time = c.Add(time.Duration(value) * time.Microsecond)
	case "milli":
		c.Time = c.Add(time.Duration(value) * time.Millisecond)
	case "second":
		c.Time = c.Add(time.Duration(value) * time.Second)
	case "minute":
		c.Time = c.Add(time.Duration(value) * time.Minute)
	case "hour":
		c.Time = c.Add(time.Duration(value) * time.Hour)
	case "day":
		c.Time = c.AddDate(0, 0, value)
	case "month":
		c.Time = c.AddDate(0, value, 0)
	case "quarter":
		c.Time = c.AddDate(0, value*MonthsPerQuarter, 0)
	case "year":
		c.Time = c.AddDate(value, 0, 0)
	case "decade":
		c.Time = c.AddDate(value*YearsPerDecade, 0, 0)
	case "century":
		c.Time = c.AddDate(value*YearsPerCentury, 0, 0)
	case "millennium":
		c.Time = c.AddDate(value*YearsPerMillennium, 0, 0)
	default:
		panic(errors.New("unknown unit to add"))
	}
	return c
}

// Sub given units to the current instance.
func (c *Carbon) SubUnit(which string, value int) *Carbon {
	return c.AddUnit(which, -value)
}
