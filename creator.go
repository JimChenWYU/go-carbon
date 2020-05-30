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

func Today(loc string) (*Carbon, error) {
	return RawParse("today", loc)

}

func Tomorrow(loc string) (*Carbon, error) {
	return RawParse("tomorrow", loc)
}

func Yesterday(loc string) (*Carbon, error) {
	return RawParse("yesterday", loc)
}

func Create(year int, month time.Month, day int, hour int, minute int, second int, ns int, loc string) (*Carbon, error) {
	l, err := time.LoadLocation(loc)
	if err != nil {
		return nil, err
	}

	return NewCarbon(time.Date(year, month, day, hour, minute, second, ns, l)), nil
}

func CreateFromDate(year int, month time.Month, day int, loc string) (*Carbon, error) {
	h, m, s := Now().Clock()
	ns := Now().Nanosecond()
	return Create(year, month, day, h, m, s, ns, loc)
}

func CreateFromTime(h, m, s, ns int, loc string) (*Carbon, error) {
	year, month, day := Now().Date()
	return Create(year, month, day, h, m, s, ns, loc)
}

func RawParse(t string, loc string) (*Carbon, error) {
	l, err := time.LoadLocation(loc)
	if err != nil {
		return nil, err
	}

	switch t {
	case "today":
		return NewCarbon(Now().In(l)), nil
	case "tomorrow":
		today := NewCarbon(Now().In(l))
		return today.AddDays(1), nil
	case "yesterday":
		today := NewCarbon(Now().In(l))
		return today.SubDays(1), nil
	}

	return nil, errors.New("parse syntax error")
}
