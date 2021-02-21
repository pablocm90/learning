package prompt

import (
	"testing"
)

func TestMessageUser(t *testing.T) {
	message := "Message !"

	MessageUser(&message, "And then other")

	if message != "And then other" {
		t.Error("Expected: And then other, got: ", message)
	}

}
