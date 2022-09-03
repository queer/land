package main

import (
	"fmt"
	"syscall"
)

func main() {
	fmt.Println("starting init")
	exe := "%EXE%"
	finalArgs := []string{"%ARGS%"}

	syscall.Exec(exe, finalArgs, []string{})
}
