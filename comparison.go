package carbon

import (
	"time"
)

// Determines if the instance is equal to another
func (c *Carbon) EqualTo(v interface{}) bool {
	c2 := c.resolveCarbon(v)
	return c.Equal(c2.Time)
}

// Determines if the instance is not equal to another
func (c *Carbon) NotEqualTo(v interface{}) bool {
	c2 := c.resolveCarbon(v)
    return !c.EqualTo(c2.Time)
}

// Determines if the instance is greater (after) than another
func (c *Carbon) GreaterThan(v interface{}) bool {
	c2 := c.resolveCarbon(v)
    return c.After(c2.Time)
}

// Determines if the instance is greater (after) than or equal to another
func (c *Carbon) GreaterThanOrEqualTo(v interface{}) bool {
	c2 := c.resolveCarbon(v)
    return c.After(c2.Time) || c.Equal(c2.Time)
}

// Determines if the instance is less (before) than another
func (c *Carbon) LessThan(v interface{}) bool {
	c2 := c.resolveCarbon(v)
	return c.Before(c2.Time)
}

// Determines if the instance is less (before) or equal to another
func (c *Carbon) LessThanOrEqualTo(v interface{}) bool {
	c2 := c.resolveCarbon(v)
    return c.Before(c2.Time) || c.Equal(c2.Time)
}

// Determines if the instance is between two others.
func (c *Carbon) Between(min interface{}, max interface{}, equal bool) bool {
	cmin := c.resolveCarbon(min)
	if cmin.GreaterThan(max) {
		min, max = max, min
	}
	if equal {
		return c.GreaterThanOrEqualTo(min) && c.LessThanOrEqualTo(max)
	}

	return c.GreaterThan(min) && c.LessThan(max)
}

// Determines if the instance is between two others, bounds included.
func (c *Carbon) BetweenIncluded(min interface{}, max interface{}) bool {
    return c.Between(min, max, true)
}

// Determines if the instance is between two others, bounds excluded.
func (c *Carbon) BetweenExcluded(min interface{}, max interface{}) bool {
	return c.Between(min, max, false)
}

// Determines if the instance is a weekend day.
func (c *Carbon) IsWeekend() bool {
    return c.Weekday() == time.Saturday || c.Weekday() == time.Sunday
}

// Determines if the instance is yesterday.
func (c *Carbon) IsYesterday() bool {
	yesterday, err := Yesterday(c.Timezone())
	if err != nil {
		return false
	}

    return c.ToDateString() == yesterday.ToDateString()
}

// Determines if the instance is today.
func (c *Carbon) IsToday() bool {
	today, err := Today(c.Timezone())
	if err != nil {
		return false
	}

	return c.ToDateString() == today.ToDateString()
}

// Determines if the instance is tomorrow.
func (c *Carbon) IsTomorrow() bool {
	tomorrow, err := Tomorrow(c.Timezone())
	if err != nil {
		return false
	}

    return c.ToDateString() == tomorrow.ToDateString()
}

// Determines if the instance is a leap year.
func (c *Carbon) IsLeadYear() bool {
	year := c.Year()
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// Determines if the instance is a long year
func (c *Carbon) IsLongYear() bool {
	c2 := Create(c.Year(), time.December, 31, 28, 0, 0, 0, c.Timezone())
	return c2.WeekOfYear() == 53
}

// Compares the formatted values of the two dates.
func (c *Carbon) IsSameAs(format string, v interface{}) bool {
    return c.Format(format) == c.resolveCarbon(v).Format(format)
}

func (c *Carbon) IsSameUnit(unit string, v interface{}) bool {
	var format string
	switch unit {
	case "year":
		format = "2006"
	case "day":
		format = DateFormat
	case "hour":
		format = strjoin(DateFormat, " ", HourFormat)
	case "minute":
		format = strjoin(DateFormat, " ", HourMinuteFormat)
	case "second":
		format = strjoin(DateFormat, " ", TimeFormat)
	case "micro":
		fallthrough
	case "microsecond":
		format = strjoin(DateFormat, " ", MicrosecondTimeFormat)
	default:
		return false
	}

	return c.IsSameAs(format, v)
}

// Checks if the passed in date is in the same quarter as the instance quarter (and year if needed).
func (c *Carbon) IsSameQuarter(v interface{}, ofSameYear bool) bool {
   t := c.resolveCarbon(v)
   return c.Quarter() == t.Quarter() && (!ofSameYear || c.IsSameYear(t))
}

// Checks if the given date is in the same year as the instance. If null passed, compare to now (with the same timezone).
func (c *Carbon) IsSameYear(v interface{}) bool {
    return c.IsSameUnit("year", v)
}

// Checks if the passed in date is in the same month as the instance´s month.
func (c *Carbon) IsSameMonth(v interface{}, ofSameYear bool) bool {
	var format string
	if ofSameYear {
		format = "2006-01"
	} else {
		format = "01"
	}
    return c.IsSameAs(format, v)
}

// Checks if the given date is in the same week as the instance. If null passed, compare to now (with the same timezone).
func (c *Carbon) IsSameWeek(v interface{}) bool {
	resolve := c.resolveCarbon(v)
	cs := strjoin(c.Format("2006"), "-", string(c.WeekOfYear()))
	vs := strjoin(resolve.Format("2006"), "-", string(resolve.WeekOfYear()))
	return cs == vs
}

// Checks if the given date is in the same day as the instance. If null passed, compare to now (with the same timezone).
func (c *Carbon) IsSameDay(v interface{}) bool {
	return c.IsSameUnit("day", v)
}

// Checks if the given date is in the same hour as the instance. If null passed, compare to now (with the same timezone).
func (c *Carbon) IsSameHour(v interface{}) bool {
	return c.IsSameUnit("hour", v)
}

// Checks if the given date is in the same minute as the instance. If null passed, compare to now (with the same timezone).
func (c *Carbon) IsSameMinute(v interface{}) bool {
	return c.IsSameUnit("minute", v)
}

// Checks if the given date is in the same second as the instance. If null passed, compare to now (with the same timezone).
func (c *Carbon) IsSameSecond(v interface{}) bool {
	return c.IsSameUnit("second", v)
}

// Checks if the given date is in the same microsecond as the instance. If null passed, compare to now (with the same timezone).
func (c *Carbon) IsSameMicro(v interface{}) bool {
	return c.IsSameUnit("micro", v)
}

// Checks if the given date is in the same microsecond as the instance. If null passed, compare to now (with the same timezone).
func (c *Carbon) IsSameMicrosecond(v interface{}) bool {
	return c.IsSameUnit("microsecond", v)
}

// Checks if this day is a specific day of the week.
func (c *Carbon) IsDayOfWeek(dayOfWeek time.Weekday) bool {
	return c.Weekday() == dayOfWeek
}

// Check if its the birthday. Compares the date/month values of the two dates.
func (c *Carbon) IsBirthday(v interface{}) bool {
	return c.IsSameAs("0102", v)
}

// Check if today is the last day of the Month
func (c *Carbon) IsLastOfMonth() bool {
	return c.Day() == c.DaysInMonth()
}

// Check if the instance is start of day / midnight.
func (c *Carbon) IsStartOfDay(checkMicroseconds bool) bool {
	if checkMicroseconds {
		return c.Format(MicrosecondTimeFormat) == "00:00:00.000000"
	} else {
		return c.Format(TimeFormat) == "00:00:00"
	}
}

// Check if the instance is end of day.
func (c *Carbon) IsEndOfDay(checkMicroseconds bool) bool {
	if checkMicroseconds {
		return c.Format(MicrosecondTimeFormat) == "23:59:59.999999"
	} else {
		return c.Format(TimeFormat) == "23:59:59"
	}
}

// Check if the instance is midday.
func (c *Carbon) isMidday() bool {
	return c.Format("3:04:05") == "12:00:00"
}
