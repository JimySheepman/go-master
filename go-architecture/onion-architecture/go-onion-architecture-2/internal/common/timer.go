// Package common holds useful code to be used across all package and layers.
package common

import "time"

// Timer - timer basic struct
type Timer struct{}

// NewTimer - instantiates a new timer
func NewTimer() *Timer {
	return &Timer{}
}

// Now returns the current time
func (t Timer) Now() time.Time {
	return time.Now()
}
