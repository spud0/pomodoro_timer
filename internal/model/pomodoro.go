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
func NewPom (dur, stopAfter time.Duratino) *Pomodoro {
	
	return &Pomodoro {
		Active : false, 
		Duration: dur, 
		StopAfter : stopAfter, 
	} 
}

// Starts the timer
func (pom *Pomodoro) Start() {

	pom.Active = true
	pom.StartTime = time.Now()
	pom.EndTime = pom.StartTime.Add(pom.Duration)
}
