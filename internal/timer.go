package internal

import (
	"time"
	"fmt"
)

// Decrement the time object passed 	
func updateTime (timeLeft *time.Time) time.Time {	
	if !timeLeft.IsZero() {
		*timeLeft = timeLeft.Add(-1 * time.Second); 	
	} else {
		// Return the 'Zero Time'
		return timeLeft	
	} 

}

func updateTui(timeLeft) string {	
	tuiTime := updateTime()
	tuiTime.String()
}	
