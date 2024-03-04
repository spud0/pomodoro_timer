package internal

import (
	"time"
	"fmt"
)

// Decrement the time object passed 	
func UpdateTime (timeLeft *time.Time) time.Time {	
	// Return the time regardless
	return timeLeft.Add(-1 * time.Second)
}

func UpdateTui(timeLeft *time.Time) string {	
	// Format the time into a string	
	return tuiTime.Format("00:00:00")
}

