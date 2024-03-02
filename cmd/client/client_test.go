package main

import (
	"os"
	"testing"

	"github.com/AleksanderBobinski/learntogo/v2/pkg/connection"
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

func TestSendFailIfNoServer(t *testing.T) {
	err := connection.Send("localhost:10", "")
	if err == nil {
		t.Fatalf("Send should fail if it can't deliver a message.")
	}
}
