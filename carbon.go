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
	HourFormat            = "15"
	DayDateTimeFormat     = "Mon, Jan 2, 2006 3:04 PM"
	CookieFormat          = "Monday, 02-Jan-2006 15:04:05 MST"
	RFC822Format          = "Mon, 02 Jan 06 15:04:05 -0700"
	RFC850Format          = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1036Format         = "Mon, 02 Jan 06 15:04:05 -0700"
	RFC1123Format         = "Mon, 02 Jan 2006 15:04:05 -0700"
	RFC2822Format         = "Mon, 02 Jan 2006 15:04:05 -0700"
	RFC3339Format         = "2006-01-02T15:04:05-07:00"
	RSSFormat             = "Mon, 02 Jan 2006 15:04:05 -0700"
	W3CFormat             = "2006-01-02T15:04:05-07:00"
	AtomFormat            = "2006-01-02T15:04:05-07:00"

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

var (
	// Format to use for __toString method when type juggling occurs.
	toStringFormat = DefaultFormat
)

type Carbon struct {
	time.Time
}

// Set the default format used when type juggling a Carbon instance to a string
func SetToStringFormat(format string) {
	toStringFormat = format
}

// Reset the format used to the default when type juggling a Carbon instance to a string
func ResetToStringFormat() {
	toStringFormat = DefaultFormat
}

// Format the instance as a string using the set format
func (c *Carbon) String() string {
	return c.Format(toStringFormat)
}

// Format the instance as date and time
func (c *Carbon) ToDateTimeString(unitPrecision string) string {
	if unitPrecision == "" {
		return c.Format(DefaultFormat)
	}

	format, err := getTimeFormatByPrecision(unitPrecision)
	if err != nil {
		panic(err)
	}

	//var sb strings.Builder
	//sb.WriteString(DateFormat)
	//sb.WriteString(" ")
	//sb.WriteString(format)
	return c.Format(strjoin(DateFormat, " ", format))
}

// Format the instance as date
func (c *Carbon) ToDateString() string {
	return c.Format(DateFormat)
}

// Format the instance as a readable date
func (c *Carbon) ToFormattedDateString() string {
	return c.Format(FormattedDateFormat)
}

// Format the instance as time
func (c *Carbon) ToTimeString() string {
	return c.Format(TimeFormat)
}

// Format the instance as date and time T-separated with no timezone
func (c *Carbon) ToDateTimeLocalString(unitPercision string) string {
	format, err := getTimeFormatByPrecision(unitPercision)
	if err != nil {
		panic(err)
	}
	return c.Format(strjoin(DateFormat, "T", format))
}

// Format the instance with day, date and time
func (c *Carbon) ToDayDateTimeString() string {
	return c.Format(DayDateTimeFormat)
}

// Format the instance as ATOM
func (c *Carbon) ToAtomString() string {
	return c.Format(AtomFormat)
}

// Format the instance as ISO8601
func (c *Carbon) ToIso8601String() string {
	return c.ToAtomString()
}

// Format the instance as COOKIE
func (c *Carbon) ToCookieString() string {
	return c.Format(CookieFormat)
}

// Format the instance as RFC822
func (c *Carbon) ToRfc822String() string {
	return c.Format(RFC822Format)
}

// Format the instance as RFC850
func (c *Carbon) ToRfc850String() string {
	return c.Format(RFC850Format)
}

// Format the instance as RFC1036
func (c *Carbon) ToRfc1036String() string {
	return c.Format(RFC1036Format)
}

// Format the instance as RFC1123
func (c *Carbon) ToRfc1123String() string {
	return c.Format(RFC1123Format)
}

// Format the instance as RFC2822
func (c *Carbon) ToRfc2822String() string {
	return c.Format(RFC2822Format)
}

// Format the instance as RFC3339
func (c *Carbon) ToRfc3339String() string {
	return c.Format(RFC3339Format)
}

// Format the instance as Rss
func (c *Carbon) ToRssString() string {
	return c.Format(RSSFormat)
}

func (c *Carbon) ToW3cString() string {
	return c.Format(W3CFormat)
}

// Return a format from H:i to H:i:s.u according to given unit precision.
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
