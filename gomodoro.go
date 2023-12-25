package main

import (
	"fmt"
	"time"

	"github.com/gen2brain/beeep"
)

var (
	worktime = 1 * time.Minute

	// channels for notifications within the system
	countdownChan = make(chan string)
	spinnerChan   = make(chan string)
)

func main() {

	end := time.Now().Add(worktime)

	go countdown(end, countdownChan)
	go spinner(spinnerChan)

	printer()
}

func printer() {
	var strings = [2]string{}
	for {
		select {
		case spinnerString := <-spinnerChan:
			backtrack(len(strings[0]+strings[1]) + 1)

			strings[1] = spinnerString
			fmt.Print(strings[0] + " " + strings[1])

		case countdownString := <-countdownChan:
			if countdownString == "done" {
				backtrack(len(strings[0]+strings[1]) + 1)
				printDone(worktime)
				notifyUser()

				return
			}

			backtrack(len(strings[0]+strings[1]) + 1)

			strings[0] = countdownString
			fmt.Print(strings[0] + " " + strings[1])
		}
	}
}

func spinner(spinnerChan chan<- string) {
	delay := 200 * time.Millisecond
	for {
		spinnerChan <- "|"
		time.Sleep(delay)

		spinnerChan <- "/"
		time.Sleep(delay)

		spinnerChan <- "-"
		time.Sleep(delay)

		spinnerChan <- "\\"
		time.Sleep(delay)
	}
}

func countdown(end time.Time, countdownChan chan<- string) {
	for time.Until(end).Milliseconds() > 0 {
		countdownChan <- time.Until(end).Round(time.Second).String()
		time.Sleep(time.Second)
	}

	countdownChan <- "done"
}

func backtrack(length int) {
	var backString string
	for i := 0; i < length; i++ {
		backString = backString + "\b"
	}
	fmt.Print(backString)
}

func printDone(worktime time.Duration) {
	fmt.Println("Congratulations! 🎉")
	fmt.Println("You did", worktime.String(), "of work! 🥳")
}

func notifyUser() {
	notifyErr := beeep.Notify("Gomodoro", "Time to relax!", "info.png")
	if notifyErr != nil {
		panic(notifyErr)
	}

	beepErr := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	if beepErr != nil {
		panic(beepErr)
	}
}
