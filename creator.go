package carbon

import (
	"errors"
	"time"
)

func Now() *Carbon {
	return NewCarbon(time.Now())
}

func NewCarbon(t time.Time) *Carbon {
	return &Carbon{
		Time: t,
	}
}

func Today(loc *time.Location) (*Carbon, error) {
	return RawParse("today", loc)

}

func Tomorrow(loc *time.Location) (*Carbon, error) {
	return RawParse("tomorrow", loc)
}

func Yesterday(loc *time.Location) (*Carbon, error) {
	return RawParse("yesterday", loc)
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

func RawParse(t string, loc *time.Location) (*Carbon, error) {
	switch t {
	case "today":
		return NewCarbon(Now().In(loc)), nil
	case "tomorrow":
		today := NewCarbon(Now().In(loc))
		return today.AddDays(1), nil
	case "yesterday":
		today := NewCarbon(Now().In(loc))
		return today.SubDays(1), nil
	}

	return nil, errors.New("parse syntax error")
}
