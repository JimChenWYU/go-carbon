package carbon

import (
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
	RFC822Format          = time.RFC822
	RFC850Format          = time.RFC850
	RFC1036Format         = "Mon, 02 Jan 06 15:04:05 -0700"
	RFC1123Format         = time.RFC1123
	RFC2822Format         = "Mon, 02 Jan 2006 15:04:05 -0700"
	RFC3339Format         = time.RFC3339
	RSSFormat             = "Mon, 02 Jan 2006 15:04:05 -0700"
	W3CFormat             = "2006-01-02T15:04:05-07:00"
	AtomFormat            = "2006-01-02T15:04:05-07:00"

	DaysPerWeek          = 7
	HoursPerDay          = 24
	MinutesPerHour       = 60
	SecondsPerMinute     = 60
	NanosecondsPerSecond = 1000000000
	MonthsPerQuarter     = 3
	MonthsPerYear     = 12
	YearsPerDecade       = 10
	YearsPerCentury      = 100
	YearsPerMillennium   = 1000
)

type Carbon struct {
	time.Time
}
