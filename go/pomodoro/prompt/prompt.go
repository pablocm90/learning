package prompt

import "fmt"

func WelcomeMessage() (str string, err error) {
	return fmt.Sprintln("Welcome to Pomodoro!"), nil
}

func MessageUser(input string) (str string, err error) {
	fmt.Println(input)
	return fmt.Sprintln(input), nil
}
