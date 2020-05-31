package carbon

import (
	"strings"
	"time"
)

// Join some strings to create a single string
func strjoin(str ...string) string {
	return strings.Join(str, "")
}

// Return time.location according to input string
func getLocation(loc string) *time.Location {
	l, _ := time.LoadLocation(loc)
	return l
}
