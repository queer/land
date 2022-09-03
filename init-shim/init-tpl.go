package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func main() {
	finalArgs := []string{"%ARGS%"}
	finalEnv := []string{"%ENV%"}

	for _, arg := range finalEnv {
		fmt.Println(arg)
		if strings.HasPrefix(arg, "PATH=") {
			path := arg[5:]
			os.Setenv("PATH", path)
			fmt.Println("init: set PATH =", path)
			break
		}
	}

	exe, err := exec.LookPath("%EXE%")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("init: exe =", exe, "args =", finalArgs, "env =", finalEnv)

	err = syscall.Exec(exe, finalArgs, finalEnv)
	if err != nil {
		fmt.Println(err.Error())
	}
}
