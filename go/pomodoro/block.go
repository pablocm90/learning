package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/0xAX/notificator"
)

// Block is the timing block for our pomodoro machine
type Block struct {
	WorkTime, BreakTime int
	State               string
	Current             int
}

// AskTimings for the block to be executed next
func (b Block) AskTimings() (workTime int, breakTime int, err error) {
	reader := bufio.NewReader((os.Stdin))

	fmt.Println("How long should the work time be for this block?")
	tempWorkTime, _ := reader.ReadString('\n')
	workTime, err = strconv.Atoi(strings.TrimSuffix(tempWorkTime, "\n"))
	fmt.Printf("Cool, let's set a work session to %v minutes \n", workTime)
	fmt.Println("How long should the break time be for this block?")
	tempBreakTime, _ := reader.ReadString('\n')
	breakTime, err = strconv.Atoi(strings.TrimSuffix(tempBreakTime, "\n"))
	fmt.Printf("Cool, let's set a break after your session of %v minutes \n", breakTime)
	return
}

// ChunkFinishAnnouncements notifies the user at the end of blocks and changes the state
func (b *Block) ChunkFinishAnnouncements() {
	var notify *notificator.Notificator

	notify = notificator.New(notificator.Options{
		DefaultIcon: "assets/sliced_tomato.png",
		AppName:     "Pomodoro CLI",
	})
	if b.State == "working" {
		b.State = "breaking"
		notify.Push("Work is over", "Time to do a productive break", "assets/sliced_tomato.png", "Pomodoro CLI")
		fmt.Printf("\nYou are done buddy, break time! :)\n")
	} else {
		b.State = "working"
		notify.Push("Break is over", "Time to start a new Work Block!", "assets/tomato.png", "Pomodoro CLI")
		fmt.Printf("\nGood job! \n")
	}
}

// Run runs the block of work and pause alternativelly.
func (b *Block) Run() (active bool, err error) {
	var text string
	if b.State == "working" {
		b.WorkTime, b.BreakTime, err = b.AskTimings()
		b.Current = 0
		fmt.Printf("\033[2K\rWhat should we work on this block ? \n")
		reader := bufio.NewReader((os.Stdin))
		text, _ = reader.ReadString('\n')
		if text == "" {
			return false, err
		}
		// should log here the string
		writeTaskLog(strings.TrimSuffix(text, "\n"))
		fmt.Println("Starting in :")
		for timer := 3; timer > 0; timer-- {
			fmt.Printf("\033[2K\r %d", timer)
			time.Sleep(1 * time.Second)
		}
		fmt.Printf("\033[2K\rlet's start working on: '%v' \n", strings.TrimSuffix(text, "\n"))
	}

	b.Current = 0
	var upperTimer int
	switch b.State {
	case "working":
		upperTimer = b.WorkTime
	case "breaking":
		upperTimer = b.BreakTime
	default:
		return false, nil
	}

	for b.Current < upperTimer {
		time.Sleep(1 * time.Minute)
		b.Current++
		fmt.Printf("\033[2K\r%d minutes remaining", upperTimer-b.Current)
	}

	b.ChunkFinishAnnouncements()
	return true, nil
}
