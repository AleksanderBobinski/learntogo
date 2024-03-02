package main

import (
	"os"
	"testing"
)

func TestInitialPrompt(t *testing.T) {
	stdin, _ := os.CreateTemp("", "")
	defer os.Remove(stdin.Name())
	stdout, _ := os.CreateTemp("", "")
	defer os.Remove(stdout.Name())
	expectedPrompt := ">"

	ret := connect(stdin, stdout)
	text, err := os.ReadFile(stdout.Name())
	returnedPrompt := string(text)

	if err != nil || ret != 0 || returnedPrompt != expectedPrompt {
		t.Fatalf(`Client = %q, ret = %v, want %q`, returnedPrompt, ret, expectedPrompt)
	}
}
