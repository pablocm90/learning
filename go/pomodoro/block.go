package main

import (
	"fmt"
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
		fmt.Printf("\033[2K\r %d", timer)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("\nGOOOOO")
	b.State = "working"
	for true {

		b.Current = 0
		var upperTimer int
		switch b.State {
		case "working":
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
			fmt.Printf("\033[2K\r %d minutes remaining", upperTimer-b.Current)
		}

		if b.State == "working" {
			b.State = "breaking"
		} else {
			b.State = "working"
		}

		fmt.Printf("\nYou are done buddy, %#v time! :)\n", b.State)
	}

	return false, nil
}
