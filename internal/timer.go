package internal

import (
	"time"
)

// Decrement the time object passed 	
func UpdateTime (timeLeft *time.Time) time.Time {	
	// Return the time regardless
	return timeLeft.Add(-1 * time.Second)
}

// 'alert' halfway
// func  QuarterTimeReached (timeLeft time.Time) bool { }

func UpdateTui(timeLeft *time.Time) string {	
	// Format the time into a string	
	return timeLeft.Format("00:00:00")
}

