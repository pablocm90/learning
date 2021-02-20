package main

import (
	"fmt"
	"pomodoro/prompt"
	"pomodoro/reader"
)

func main() {
	currentMessage, err := prompt.WelcomeMessage()

	run := true

	if err == nil {
		fmt.Print(currentMessage)
	}

	for run {
		currentMessage = "How long should the session be ? "
		feedback, _ := reader.UserInput(&currentMessage)
		fmt.Println(feedback)
		fmt.Println("'----------------'")
		fmt.Println("My message should have changed")

		fmt.Println(currentMessage)
		run = false
	}
}
