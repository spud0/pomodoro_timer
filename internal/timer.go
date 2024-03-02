package internal

import (
	"time"
	"fmt"
)

// Decrement the time object passed 	
func updateTime (t *time.Time) {	
	if !t.IsZero() {
		*t = t.Add(-1 * time.Second); 	
	} else {
		// Do something else
	} 

}

// TODO: Implement this function... 
func updateScreen() {
	
	screenTime := updateTime()
}	
