package reader

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestUserInput(t *testing.T) {
	userInput := "25"
	in, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer in.Close()
	_, err = io.WriteString(in, "25\n")
	if err != nil {
		t.Fatal(err)
	}

	_, err = in.Seek(0, os.SEEK_SET)
	if err != nil {
		t.Fatal(err)
	}

	n, err := UserInput(userInput)
	fmt.Println("IN TEST USER INPUT")
	fmt.Print(n)
	if n != "25" {
		t.Error("expected: 25, got:", n)
	}
}
