package main

import (
	"fmt"
	"syscall"
)

func main() {
	exe := "%EXE%"
	finalArgs := []string{"%ARGS%"}
	fmt.Println("init: exe=", exe, "args=", finalArgs)

	syscall.Exec(exe, finalArgs, []string{})
}
