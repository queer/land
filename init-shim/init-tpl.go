package main

import (
	"fmt"
	"syscall"
)

func main() {
	exe := "%EXE%"
	finalArgs := []string{"%ARGS%"}
	finalEnv := []string{"%ENV%"}
	fmt.Println("init: exe =", exe, "args =", finalArgs)

	syscall.Exec(exe, finalArgs, finalEnv)
}
