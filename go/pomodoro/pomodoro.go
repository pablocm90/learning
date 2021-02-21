package main

import (
	"fmt"
	"pomodoro/prompt"
	"pomodoro/reader"
	"strconv"
	"time"
)

// Block is the timing block for our pomodoro machine
type Block struct {
	WorkTime, BreakTime int
	State               string
	Current             int
}

// Run runs the block of work and pause alternativelly.
func (b *Block) Run() (active bool, err error) {
	fmt.Println("Starting Work...")
	for timer := 3; timer > 0; timer-- {
		fmt.Println(timer)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("GOOOOO")

	b.State = "working"
	var upperTimer int
	switch b.State {
	case "working":
		fmt.Println("I am in working")
		upperTimer = b.WorkTime
	case "breaking":
		upperTimer = b.BreakTime
	default:
		fmt.Println("I am in default")
		return false, nil
	}

	for b.Current < upperTimer {
		time.Sleep(1 * time.Minute)
		b.Current++
		fmt.Println(upperTimer-b.Current, " minutes remaining")
	}

	fmt.Println("You are done buddy :)")

	return false, nil
}

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
