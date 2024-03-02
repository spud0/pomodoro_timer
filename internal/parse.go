package main

import (
	"flag"
	"fmt"
)

func parse () {
	
	// Unsiged integers for time
	timeHrs := flag.int64("hrs",   0, "Amount of Hours");
	timeMins := flag.int64("mins", 5, "Amount of Minutes");
	timeSecs := flag.int64("secs", 0, "Amount of Seconds");

	color := flag.String("color", "white", "Color for timer");
	flag.Parse();	

	if   {

	}

	if {
	}

	fmt.Printf("Time: %s\n", *time); 
	fmt.Printf("Color: %s\n", *color); 
}

func main() {
	parse();
}
