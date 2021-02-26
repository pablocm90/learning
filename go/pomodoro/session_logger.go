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
		f, err := os.OpenFile("log/task_log.txt", os.O_APPEND|os.O_WRONLY, 0644)
		check(err)
		defer f.Close()

		f.WriteString("\n-" + message)
		return true
	}
	os.Mkdir("log", 0777)
	err := ioutil.WriteFile("log/task_log.txt", []byte("- "+message), 0777)
	check(err)
	return true
}

func readTaskLog() (taskLog bool) {
	if _, err := os.Stat("log/task_log.txt"); err == nil {
		f, err := ioutil.ReadFile("log/task_log.txt")
		check(err)

		fmt.Println("Here are the tasks you have tackled")
		fmt.Println(string(f))
		return true
	}
	fmt.Println("You have no tasks yet, good excuse to start working :D")
	return
}
