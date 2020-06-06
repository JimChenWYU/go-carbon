package carbon

import (
	"time"
)

func Now() *Carbon {
	var t time.Time
	if HasTestNow() {
		t = TestNow().Time
	} else {
		t = time.Now()
	}
	return NewCarbon(t)
}

func NowWithTz(loc *time.Location) *Carbon {
	now := Now()
	now.Time = now.Time.In(loc)
    return now
}

func NewCarbon(t time.Time) *Carbon {
	return &Carbon{
		Time: t,
	}
}

func Today(loc *time.Location) *Carbon {
	return NewCarbon(time.Now().In(loc))
}

func Tomorrow(loc *time.Location) *Carbon {
	return NewCarbon(time.Now().In(loc)).AddDays(1)
}

func Yesterday(loc *time.Location) *Carbon {
	return NewCarbon(time.Now().In(loc)).SubDays(1)
}

func Create(year int, month time.Month, day int, hour int, minute int, second int, ns int, loc *time.Location) *Carbon {
	return NewCarbon(time.Date(year, month, day, hour, minute, second, ns, loc))
}

func CreateFromDate(year int, month time.Month, day int, loc *time.Location) *Carbon {
	h, m, s := Now().Clock()
	ns := Now().Nanosecond()
	return Create(year, month, day, h, m, s, ns, loc)
}

func CreateFromTime(h, m, s, ns int, loc *time.Location) *Carbon {
	year, month, day := Now().Date()
	return Create(year, month, day, h, m, s, ns, loc)
}

func RawParse(layout, value string, loc *time.Location) *Carbon {
	t, err := time.ParseInLocation(layout, value, loc)
	if err != nil {
		panic(err)
	}
	return NewCarbon(t)
}

func RawParseWithDefaultTz(layout, value string) *Carbon {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return NewCarbon(t)
}
