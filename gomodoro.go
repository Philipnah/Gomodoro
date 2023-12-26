package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/gen2brain/beeep"
)

var (
	// channels for notifications within the system
	countdownChan = make(chan string)
	spinnerChan   = make(chan string)

	workFlag = flag.Int("work", 25, "Set the duration of the timer")
	restFlag = flag.Int("rest", 5, "Set the duration of the timer")
)

func main() {
	flag.Parse()
	workDuration := time.Duration(*workFlag) * time.Minute
	restDuration := time.Duration(*restFlag) * time.Minute

	go spinner(spinnerChan)

	for {
		go countdown(workDuration, countdownChan)

		fmt.Print("📌 Work:\n\t")
		printer()
		printDone(workDuration)
		notifyUser("Time to relax!")

		fmt.Print("\n")

		go countdown(restDuration, countdownChan)

		fmt.Print("🌟 Rest:\n\t")
		printer()
		printDone(restDuration)
		notifyUser("Time to work!")
	}
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
			backtrack(len(strings[0]+strings[1]) + 1)
			if countdownString == "done" {
				return
			}

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

func countdown(duration time.Duration, countdownChan chan<- string) {
	end := time.Now().Add(duration)

	for time.Until(end).Milliseconds() > 0 {
		countdownChan <- time.Until(end).Round(time.Second).String()
		time.Sleep(time.Second)
	}

	countdownChan <- "done"
}

func backtrack(length int) {
	var backChars string
	var backSpaces string
	for i := 0; i < length; i++ {
		backChars = backChars + "\b"
		backSpaces = backSpaces + " "
	}

	// Go back to beginning of line,
	// replace all previous chars with spaces to clear them,
	// then go back to beginning of line to have a clear slate.
	fmt.Print(backChars)
	fmt.Print(backSpaces)
	fmt.Print(backChars)
}

func printDone(worktime time.Duration) {
	fmt.Println("Congratulations! 🎉")
	fmt.Println("You did", worktime.String(), "of work! 🥳")
}

func notifyUser(message string) {
	notifyErr := beeep.Notify("Gomodoro", message, "info.png")
	if notifyErr != nil {
		panic(notifyErr)
	}

	beepErr := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	if beepErr != nil {
		panic(beepErr)
	}
}
