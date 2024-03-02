package main

import "os"

func connect(stdin, stdout *os.File) int {
	stdout.WriteString(">")
	return 0
}

func main() {
	os.Exit(connect(os.Stdin, os.Stdout))
}
