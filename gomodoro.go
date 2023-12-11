package main

import (
	"fmt"
	"time"
)

var (
	worktime = 25 * time.Minute

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
				return
			}

			backtrack(len(strings[0]+strings[1]) + 1)

			strings[0] = countdownString
			fmt.Print(strings[0] + " " + strings[1])
		}
	}
}

func spinner(spinnerChan chan<- string) {
	for {
		spinnerChan <- "|"
		time.Sleep(500 * time.Millisecond)

		spinnerChan <- "/"
		time.Sleep(500 * time.Millisecond)

		spinnerChan <- "-"
		time.Sleep(500 * time.Millisecond)

		spinnerChan <- "\\"
		time.Sleep(500 * time.Millisecond)
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
	fmt.Println("You're done! 🎉")
	fmt.Println("You did", worktime.String(), "of work! 🥳")
}
