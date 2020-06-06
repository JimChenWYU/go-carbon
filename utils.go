package carbon

import (
	"strings"
	"time"
)

// Join some strings to create a single string
func joinStrings(str ...string) string {
	return strings.Join(str, "")
}

// calculate the input value to absolute value when absolute is true
func absValueIf(v int64, absolute bool) int64 {
	if absolute && v < 0 {
		v = -v
	}
	return v
}

// Return time.location according to input string
func getLocation(loc string) *time.Location {
	l, _ := time.LoadLocation(loc)
	return l
}
