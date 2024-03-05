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
func  ... (timeLeft time.Time) bool {
	
}


// 'alert' 3/4 of the way
func  ... (timeLeft time.Time) bool {
	
}


func UpdateTui(timeLeft *time.Time) string {	
	// Format the time into a string	
	return timeLeft.Format("00:00:00")
}

