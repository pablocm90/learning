package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func writeTaskLog(message string) (status bool) {
	if _, err := os.Stat("log/task_log.txt"); err == nil {
		fmt.Println("here")
		f, err := os.OpenFile("log/task_log.txt", os.O_APPEND|os.O_WRONLY, 0644)
		check(err)
		defer f.Close()

		f.WriteString("\n-" + message)
		return true
	}
	fmt.Println("here after if")
	os.Mkdir("log", 0777)
	err := ioutil.WriteFile("log/task_log.txt", []byte("- "+message), 0777)
	check(err)
	return true
}

func readTaskLog() (taskLog bool) {
	return
}
