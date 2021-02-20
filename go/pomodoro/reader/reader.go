package reader

import (
	"bufio"
	"fmt"
	"os"
	"pomodoro/prompt"
)

func UserInput(message *string) (result string, err error) {
	prompt.MessageUser(*message)
	reader := bufio.NewReader((os.Stdin))
	text, _ := reader.ReadString('\n')
	if text == "" {
		return "", err
	}
	*message = "I DID CHANGE!"
	return fmt.Sprint(text), nil
}
