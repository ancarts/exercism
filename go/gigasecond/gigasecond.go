// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package gigasecond contains resolution for gigasecond exercise
package gigasecond

// import path for the time package from the standard library
import "time"

const secondsInGigasecond = 1000000000

// AddGigasecond shows the date and time one gigasecond after a certain date
func AddGigasecond(t time.Time) time.Time {
	gigaAge := t.Add(time.Second * secondsInGigasecond)
	return gigaAge
}
