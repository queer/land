package main

import (
	"fmt"
	"os"
	"strings"
	"syscall"
)

func main() {
	exe := "%EXE%"
	finalArgs := []string{"%ARGS%"}
	finalEnv := []string{"%ENV%"}
	fmt.Println("init: exe =", exe, "args =", finalArgs, "env =", finalEnv)

	for _, arg := range finalEnv {
		if strings.HasPrefix(arg, "PATH=") {
			os.Setenv("PATH", arg[5:])
			break
		}
	}

	err := syscall.Exec(exe, finalArgs, finalEnv)
	if err != nil {
		fmt.Println(err.Error())
	}
}
