package main

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/AleksanderBobinski/learntogo/v2/cmd/internal/tt"
	"github.com/AleksanderBobinski/learntogo/v2/pkg/connection"
)

func TestLogging(t *testing.T) {
	stdout, _ := os.CreateTemp("", "")
	defer os.Remove(stdout.Name())
	srvret := make(chan int)
	timeout := 100 * time.Millisecond
	expected := tt.RandStringRunes(rand.Int()%20 + 1)
	port := 8080

	go func() { srvret <- serve(stdout, port) }()
	for err := connection.Send(fmt.Sprintf("localhost:%d", port), expected); err != nil; {
		t.Log(err)
	}
	select {
	case ret := <-srvret:
		t.Fatalf("The server quit on it's own with %v", ret)
	case <-time.After(timeout):
	}
	text, _ := os.ReadFile(stdout.Name())
	logged := string(text)

	if logged != expected {
		t.Fatalf("Server logged %q but %q was expected.", logged, expected)
	}
}
