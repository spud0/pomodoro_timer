package main

import (
	"flag"
	"fmt"
	"os"
	"time"
	// 	"github.com/spud0/pomodoro_timer/internal"
)

func parse () (time.Time, string) {
	
	timeHrs := flag.Int64("hrs",   0, "Amount of Hours");
	timeMins := flag.Int64("mins", 5, "Amount of Minutes");
	timeSecs := flag.Int64("secs", 0, "Amount of Seconds");
	color := flag.String("color", "white", "Color for timer");

	flag.Parse();	

	if *timeHrs >= 0 && *timeMins >= 0 && *timeSecs >= 0  {
		now := time.Now()
		currentTime := time.Date(now.Year(), now.Month(), now.Day(), int(*timeHrs), 
								int(*timeMins), int(*timeSecs), 0, now.Location(),)	
		return currentTime, *color
	
	} else {
		fmt.Printf("Invalid input. Doesn't support negative time values.\n")
		flag.Usage() 
		return time.Time{}, *color
	}
}

func main() {
	timeObj, _ := parse()
	if timeObj.IsZero() {
		fmt.Printf("no time provided\n")	
		os.Exit(1)
	} else {
		fmt.Printf("time provided\n")	
	}	
}
