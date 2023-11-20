package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner()
	time.Sleep(25 * time.Minute)
	done()
}

func spinner() {
	for {
		fmt.Print("|")
		time.Sleep(500 * time.Millisecond)
		fmt.Print("\b")
		fmt.Print("/")
		time.Sleep(500 * time.Millisecond)
		fmt.Print("\b")
		fmt.Print("-")
		time.Sleep(500 * time.Millisecond)
		fmt.Print("\b")
		fmt.Print("\\")
		time.Sleep(500 * time.Millisecond)
		fmt.Print("\b")
	}
}

func done() {
	fmt.Print("\b")
	fmt.Println("You're done! 🎉")
}
