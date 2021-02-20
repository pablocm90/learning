package prompt

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestWelcomeMessage(t *testing.T) {
	want := regexp.MustCompile("Welcome to Pomodoro!\n")
	msg, err := WelcomeMessage()
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`WelcomeMessage() = %q, %v, should match for %#q, nil`, msg, err, want)
	}
}

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestMessageUser(t *testing.T) {
	want := regexp.MustCompile("Message !\n")
	msg, err := MessageUser("Message !")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`WelcomeMessage() = %q, %v, should match for %#q, nil`, msg, err, want)
	}
}
