package model

import (
	"time"
)

type Pomodoro struct {
	Active		bool 			// State for if the timer is still active. 
	StartTime 	time.Time
	EndTime 	time.Time	
	Duration	time.Duration
	StopAfter 	time.Duration	// Stop the timer for a break after X (time units). 
}

// Creates a new Pomodoro object 
func NewPom (dur, stopAfter time.Duration) *Pomodoro {
	
	return &Pomodoro {
		Active : false, 
		Duration: dur, 
		StopAfter : stopAfter, 
	} 
}

// Starts the timer
func (pom *Pomodoro) Start() {
	pom.Active = true
	pom.StartTime = time.Now(
	
	) // Add the duration to the start to get the EndTime
	pom.EndTime = pom.StartTime.Add(pom.Duration) 
}

// Stops the timer
func (pom *Pomodoro) Stop() {
	pom.Active = false
}


// Determines the remaining time
func (pom *Pomodoro) RemainingTime() time.Duration {
	if pom.Active {
		return pom.EndTime.Sub(time.Now())
	}

	return 0 
}





