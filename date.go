package carbon

import "time"

func (c *Carbon) SetTimeZone(loc string) (*Carbon, error) {
	l, err := time.LoadLocation(loc)
	if err != nil {
		return nil, err
	}

	c.Time = time.Date(c.Year(), c.Month(), c.Day(), c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), l);
	return c, nil
}

func (c *Carbon) Timezone() string {
	return c.Location().String()
}

func (c *Carbon) AddSeconds(s int) *Carbon {
	return NewCarbon(c.Add(time.Duration(s) * time.Second))
}

func (c *Carbon) AddMinutes(m int) *Carbon {
	return NewCarbon(c.Add(time.Duration(m) * time.Minute))
}

func (c *Carbon) AddHours(h int) *Carbon {
	return NewCarbon(c.Add(time.Duration(h) * time.Hour))
}

func (c *Carbon) AddDays(d int) *Carbon {
	return NewCarbon(c.AddDate(0, 0, d))
}

func (c *Carbon) AddWeeks(w int) *Carbon {
	return c.AddDays(DaysPerWeek * w)
}

func (c *Carbon) AddMonths(m int) *Carbon {
	return NewCarbon(c.AddDate(0, m, 0))
}

func (c *Carbon) AddYears(y int) *Carbon {
	return NewCarbon(c.AddDate(y, 0, 0))
}

func (c *Carbon) SubSeconds(s int) *Carbon {
	return c.AddSeconds(-s)
}

func (c *Carbon) SubMinutes(m int) *Carbon {
	return c.AddMinutes(-m)
}

func (c *Carbon) SubHours(h int) *Carbon {
	return c.AddHours(-h)
}

func (c *Carbon) SubDays(d int) *Carbon {
	return NewCarbon(c.AddDate(0, 0, -d))
}

func (c *Carbon) SubWeeks(w int) *Carbon {
	return c.SubDays(DaysPerWeek * w)
}

func (c *Carbon) SubMonths(m int) *Carbon {
	return NewCarbon(c.AddDate(0, -m, 0))
}

func (c *Carbon) SubYears(y int) *Carbon {
	return NewCarbon(c.AddDate(-y, 0, 0))
}

func (c *Carbon) SetTimestamp(t int64) *Carbon {
	c.Time = time.Unix(t, 0)
	return c
}

func (c *Carbon) SetNanoSeconds(ns int) *Carbon {
	c.Time = time.Date(c.Year(), c.Month(), c.Day(), c.Hour(), c.Minute(), c.Second(), ns, c.Location())
	return c
}

func (c *Carbon) SetSeconds(s int) *Carbon {
	c.Time = time.Date(c.Year(), c.Month(), c.Day(), c.Hour(), c.Minute(), s, c.Nanosecond(), c.Location())
	return c
}

func (c *Carbon) SetMinutes(m int) *Carbon {
	c.Time = time.Date(c.Year(), c.Month(), c.Day(), c.Hour(), m, c.Second(), c.Nanosecond(), c.Location())
	return c
}

func (c *Carbon) SetHours(h int) *Carbon {
	c.Time = time.Date(c.Year(), c.Month(), c.Day(), h, c.Minute(), c.Second(), c.Nanosecond(), c.Location())
	return c
}

func (c *Carbon) SetDays(d int) *Carbon {
	c.Time = time.Date(c.Year(), c.Month(), d, c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Location())
	return c
}

func (c *Carbon) SetMonths(m time.Month) *Carbon {
	c.Time = time.Date(c.Year(), m, c.Day(), c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Location())
	return c
}

func (c *Carbon) SetYears(y int) *Carbon {
	c.Time = time.Date(y, c.Month(), c.Day(), c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Location())
	return c
}