package carbon

import "time"

// Create a Carbon instance from a timestamp.
func CreateFromTimestamp(timestamp int64, location string) (*Carbon, error) {
	l, err := time.LoadLocation(location)
	if err != nil {
		return nil, err
	}

	t := NewCarbon(Now().In(l))
	t.SetTimestamp(timestamp)
	return t, nil
}

// Returns the UNIX timestamp for the current date.
func (c *Carbon) Timestamp() int64 {
	return c.Unix()
}

// Returns the UNIX millitimestamp for the current date.
func (c *Carbon) MilliTimestamp() int64 {
	return c.UnixNano() / 1e6
}

// Returns the UNIX microtimestamp for the current date.
func (c *Carbon) MicroTimestamp() int64 {
	return c.UnixNano() / 1e3
}
