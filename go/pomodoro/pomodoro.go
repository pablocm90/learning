package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func menuPrompt() (action string) {

	fmt.Println("What would you like to do ?")
	fmt.Println("[l] list the tasks I have done")
	fmt.Println("[n] new work block")
	fmt.Println("[s] stop")
	reader := bufio.NewReader((os.Stdin))
	action, _ = reader.ReadString('\n')
	action = strings.TrimSuffix(action, "\n")
	return
}

func main() {
	fmt.Println("Welcome to pomodoro")
	run := true
	block := &Block{
		State: "working",
	}

	for run {
		var err error
		if block.State == "breaking" {
			block.Run()
		}
		action := menuPrompt()
		switch action {
		case "s":
			run = false
		case "n":
			run, err = block.Run()
		case "l":
			readTaskLog()
		default:
			fmt.Println("I do not know this command. Please try again")
		}
		if err != nil {
			panic(err)
		}
	}
}
