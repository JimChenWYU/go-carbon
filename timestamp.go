package carbon

import "time"

func CreateFromTimestamp(timestamp int64, location string) (*Carbon, error)  {
	l, err := time.LoadLocation(location)
	if err != nil {
		return nil, err
	}

	t := NewCarbon(Now().In(l))
	t.SetTimestamp(timestamp)
	return t, nil
}

func (c *Carbon) Timestamp() int64 {
	return c.Unix()
}

func (c *Carbon) MilliTimestamp() int64 {
	return c.UnixNano() / 1e6
}

func (c *Carbon) MicroTimestamp() int64 {
	return c.UnixNano() / 1e3
}
