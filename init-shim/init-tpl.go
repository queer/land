package main

import (
	"syscall"
)

func main() {
	exe := "%EXE%"
	finalArgs := []string{"%ARGS%"}

	syscall.Exec(exe, finalArgs, []string{})
}
