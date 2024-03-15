package internal

import (
	"fmt"
	"time"
	"log"
	"github.com/jroimartin/gocui"
)



func Run () {

	screen, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}

	defer screen.Close()
	
	var timerView *gocui.View
	if err := layout(screen, timerView); err != nil {
		log.Panicln(err) 
	} 

	canceledScreen := screen.SetKeybinding("", gocui.KeyCtrlC,  gocui.ModNone, quit); 
	if canceledScreen != nil {
		log.Panicln(canceledScreen)
	}

	// Main loop
	go func () {
		if err := screen.MainLoop(); err != nil && err != gocui.ErrQuit {
			log.Panicln(err) 
		} 
	}()

	// Main Countdown
	timeLeft, _ := ParseInputs()	
	timeRem := &timeLeft
	timeEnd := time.Date(int(timeLeft.Year()), timeLeft.Month(), int(timeLeft.Day()), int(0), int(0), int(0), int(0), timeLeft.Location())

	
	// Decrement Time provided and update the TUI	
	go func () {
		for !timeLeft.Equal(timeEnd) {		
			timeLeft = UpdateTime(timeRem)
			time.Sleep(1 * time.Second)
			timeStr := UpdateTui(timeRem)
			
			screen.Update ( func (gui *gocui.Gui) error {
				if timerView != nil {
					timerView.Clear()
					fmt.Fprintln(timerView, timeStr)
				}
				return nil
			}) 
		}
	}()	

	select{}
}

func layout(screen *gocui.Gui, timerView *gocui.View) error {
    maxX, maxY := screen.Size()
    if timerView, err := screen.SetView("hello", (maxX/2) - 15 , (maxY/2) - 5 , (maxX/2) + 15,  (maxY/2) + 5); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }

		timerAmount, _ := ParseInputs()
		timeStr := timerAmount.Format("00:00:00")
        fmt.Fprintln(timerView, timeStr) 

        timerView.BgColor = gocui.ColorBlack // Changing background color
        timerView.FgColor = gocui.ColorWhite // Changing font color
        timerView.SelBgColor = gocui.ColorDefault // Changing selected background color
        timerView.SelFgColor = gocui.ColorWhite // Changing selected font color
        timerView.Wrap = true // Enabling text wrapping
        timerView.Frame = true // Enable a frame around the view
        timerView.Title = "timer"// Setting a title for the view
        timerView.Autoscroll = true // Automatically scroll to the bottom when new content is added
        timerView.Highlight = true // Enable highlighting
		
    }
    return nil
}


func quit (gui *gocui.Gui, view *gocui.View) error {
	return gocui.ErrQuit
}
