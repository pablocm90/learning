package main

import (
	"fmt"
	"pomodoro/prompt"
	"pomodoro/reader"
	"strconv"
)

func main() {
	currentMessage := "Welcome to pomodoro"
	prompt.MessageUser(&currentMessage, "How long should the session be ? ")

	run := true

	for run {
		workTime, _ := reader.UserInput(&currentMessage, "What about breaktime ?")
		fmt.Println("Cool, let's set a work session to", workTime)
		breakTime, _ := reader.UserInput(&currentMessage, "All Set! ")
		fmt.Println("Cool, let's set a break session to", breakTime)
		prompt.MessageUser(&currentMessage, "")

		workTimeConverted, _ := strconv.Atoi(workTime)
		breakTimeConverted, _ := strconv.Atoi(breakTime)
		blockState := "idle"
		current := 0
		block := &Block{
			WorkTime:  workTimeConverted,
			BreakTime: breakTimeConverted,
			State:     blockState,
			Current:   current,
		}

		var err error
		run, err = block.Run()

		if err != nil {
			panic("I panic!")
		}
	}
}
