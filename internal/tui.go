package internal

import (
	"fmt"
	"time"
	"pomodoro_timer/internal"
	// "github.com/jroimartin/gocui"
)



func Run () {

	// Main Countdown
	timeLeft, color := internal.ParseInputs()	
	timeEnd := time.Date(timeLeft.Year(), timeLeft.Month(), timeLeft.Day(), 0, 0, 0)

	// Decrement Time provided and update the TUI	
	for !timeLeft == timeEnd {
		timeLeft = internal.UpdateTime(timeLeft)
		time.sleep(1 * time.Second)
		timeStr := internal.UpdateTui(timeLeft)
	} 

	
}
