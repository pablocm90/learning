package prompt

import "fmt"

// MessageUser the input and changes input to the next message
func MessageUser(input *string, nextMessage string) {
	fmt.Println(*input)
	*input = nextMessage
}
