package reader

import (
	"bufio"
	"os"
	"pomodoro/prompt"
	"strings"
)

// UserInput prompts the user, gets an input and changes currentMessage
func UserInput(message *string, newPrompt string) (result string, err error) {
	prompt.MessageUser(message, newPrompt)
	reader := bufio.NewReader((os.Stdin))
	text, _ := reader.ReadString('\n')
	if text == "" {
		return "", err
	}
	return strings.TrimSuffix(text, "\n"), nil
}
