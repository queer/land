package main

import (
	"fmt"
	"syscall"
)

func main() {
	exe := "%EXE%"
	finalArgs := []string{"%ARGS%"}
	finalEnv := []string{"%ENV%"}
	fmt.Println("init: exe =", exe, "args =", finalArgs, "env =", finalEnv)

	err := syscall.Exec(exe, finalArgs, finalEnv)
	if err != nil {
		fmt.Println(err.Error())
	}
}
