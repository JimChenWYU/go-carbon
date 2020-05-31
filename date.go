package carbon

import "time"

// Set the instance's timezone from a string or object.
func (c *Carbon) SetTimezone(loc string) (*Carbon, error) {
	l, err := time.LoadLocation(loc)
	if err != nil {
		return nil, err
	}

	c.Time = time.Date(c.Year(), c.Month(), c.Day(), c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), l)
	return c, nil
}

// Get the TimeZone associated with the Carbon instance (as CarbonTimeZone).
func (c *Carbon) Timezone() string {
	return c.Location().String()
}

// Add milliseconds to the instance.
func (c *Carbon) AddMilli(ms int) *Carbon {
	return NewCarbon(c.Add(time.Duration(ms) * time.Millisecond))
}

// Add microseconds to the instance.
func (c *Carbon) AddMicro(ms int) *Carbon {
	return NewCarbon(c.Add(time.Duration(ms) * time.Microsecond))
}

// Add seconds to the instance.
func (c *Carbon) AddSeconds(s int) *Carbon {
	return NewCarbon(c.Add(time.Duration(s) * time.Second))
}

// Add minutes to the instance.
func (c *Carbon) AddMinutes(m int) *Carbon {
	return NewCarbon(c.Add(time.Duration(m) * time.Minute))
}

// Add hours to the instance.
func (c *Carbon) AddHours(h int) *Carbon {
	return NewCarbon(c.Add(time.Duration(h) * time.Hour))
}

// Add days to the instance.
func (c *Carbon) AddDays(d int) *Carbon {
	return NewCarbon(c.AddDate(0, 0, d))
}

// Add weeks to the instance.
func (c *Carbon) AddWeeks(w int) *Carbon {
	return c.AddDays(DaysPerWeek * w)
}

// Add months to the instance.
func (c *Carbon) AddMonths(m int) *Carbon {
	return NewCarbon(c.AddDate(0, m, 0))
}

// Add years to the instance.
func (c *Carbon) AddYears(y int) *Carbon {
	return NewCarbon(c.AddDate(y, 0, 0))
}

// Sub milliseconds to the instance.
func (c *Carbon) SubMilli(ms int) *Carbon {
	return c.AddMilli(-ms)
}

// Sub microseconds to the instance.
func (c *Carbon) SubMicro(ms int) *Carbon {
	return c.AddMicro(-ms)
}

// Sub seconds to the instance.
func (c *Carbon) SubSeconds(s int) *Carbon {
	return c.AddSeconds(-s)
}

// Sub minutes to the instance.
func (c *Carbon) SubMinutes(m int) *Carbon {
	return c.AddMinutes(-m)
}

// Sub hours to the instance.
func (c *Carbon) SubHours(h int) *Carbon {
	return c.AddHours(-h)
}

// Sub days to the instance.
func (c *Carbon) SubDays(d int) *Carbon {
	return c.AddDays(-d)
}

// Sub weeks to the instance.
func (c *Carbon) SubWeeks(w int) *Carbon {
	return c.AddWeeks(-w)
}

// Sub months to the instance.
func (c *Carbon) SubMonths(m int) *Carbon {
	return c.AddMonths(-m)
}

// Sub years to the instance.
func (c *Carbon) SubYears(y int) *Carbon {
	return c.AddYears(-y)
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
