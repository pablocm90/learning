package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	b.State = "working"
	for true {
		var text string
		if b.State == "working" {
			fmt.Printf("\033[2K\r What should we work on this block ? \n")
			reader := bufio.NewReader((os.Stdin))
			text, _ = reader.ReadString('\n')
			if text == "" {
				return false, err
			}
			// should log here the string
			writeTaskLog(strings.TrimSuffix(text, "\n"))
			fmt.Printf("let's start with %#v block \n", strings.TrimSuffix(text, "\n"))
		}
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
