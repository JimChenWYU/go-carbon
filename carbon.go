package carbon

import (
	"errors"
	"time"
)

// Represents the different string formats for dates
const (
	DefaultFormat         = "2006-01-02 15:04:05"
	DateFormat            = "2006-01-02"
	FormattedDateFormat   = "Jan 2, 2006"
	TimeFormat            = "15:04:05"
	MillisecondTimeFormat = "15:04:05.000"
	MicrosecondTimeFormat = "15:04:05.000000"
	HourMinuteFormat      = "15:04"
	DayDateTimeFormat     = "Mon, Jan 2, 2006 3:04 PM"
	CookieFormat          = "Monday, 02-Jan-2006 15:04:05 MST"
	RFC822Format          = "Mon, 02 Jan 06 15:04:05 -0700"
	RFC850Format          = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1036Format         = "Mon, 02 Jan 06 15:04:05 -0700"
	RFC2822Format         = "Mon, 02 Jan 2006 15:04:05 -0700"
	RFC3339Format         = "2006-01-02T15:04:05-07:00"
	RSSFormat             = "Mon, 02 Jan 2006 15:04:05 -0700"

	DaysPerWeek          = 7
	HoursPerDay          = 24
	MinutesPerHour       = 60
	SecondsPerMinute     = 60
	NanosecondsPerSecond = 1000000000
	MonthsPerQuarter     = 3
	YearsPerDecade       = 10
	YearsPerCentury      = 100
	YearsPerMillennium   = 1000
)

type Carbon struct {
	time.Time
}

func (c *Carbon) ToDateTimeString(unitPrecision string) (string, error) {
	if unitPrecision == "" {
		return c.Format(DefaultFormat), nil
	}

	format, err := getTimeFormatByPrecision(unitPrecision)
	if err != nil {
		return "", err
	}

	//var sb strings.Builder
	//sb.WriteString(DateFormat)
	//sb.WriteString(" ")
	//sb.WriteString(format)
	return c.Format(strjoin(DateFormat, " ", format)), nil
}

func (c *Carbon) ToDateString() string {
	return c.Format(DateFormat)
}

func (c *Carbon) ToFormattedDateString() string {
	return c.Format(FormattedDateFormat)
}

func (c *Carbon) ToTimeString() string {
	return c.Format(TimeFormat)
}

func (c *Carbon) ToDayDateTimeString() string {
	return c.Format(DayDateTimeFormat)
}

func (c *Carbon) ToCookieString() string {
	return c.Format(CookieFormat)
}

func (c *Carbon) ToRfc822String() string {
	return c.Format(RFC822Format)
}

func (c *Carbon) ToRfc850String() string {
	return c.Format(RFC850Format)
}

func (c *Carbon) ToRfc1036String() string {
	return c.Format(RFC1036Format)
}

func (c *Carbon) ToRfc2822String() string {
	return c.Format(RFC2822Format)
}

func (c *Carbon) ToRfc3339String() string {
	return c.Format(RFC3339Format)
}

func (c *Carbon) ToRssString() string {
	return c.Format(RSSFormat)
}

func getTimeFormatByPrecision(unitPrecision string) (string, error) {
	switch unitPrecision {
	case "minute":
		return HourMinuteFormat, nil
	case "second":
		return TimeFormat, nil
	case "m":
		fallthrough
	case "millisecond":
		return MillisecondTimeFormat, nil
	case "µ":
		fallthrough
	case "microsecond":
		return MicrosecondTimeFormat, nil
	}

	return "", errors.New("unknown unitPrecision")
}
