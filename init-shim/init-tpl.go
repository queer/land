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
	entrypoint := "%ENTRYPOINT%"
	entrypointArgs := []string{"%ENTRYPOINT_ARGS%"}

	// Set up initial path if provided
	for _, arg := range finalEnv {
		fmt.Println(arg)
		if strings.HasPrefix(arg, "PATH=") {
			path := arg[5:]
			os.Setenv("PATH", path)
			fmt.Println("init: set PATH =", path)
			break
		}
	}

	// Look up the exe since syscall.Exec doesn't
	exe, err := exec.LookPath("%EXE%")
	failFast(err)

	// Run entrypoint
	if entrypoint != "" {
		fmt.Println("init: entrypoint: exe =", entrypoint, "args =", entrypointArgs, "env =", finalEnv)

		cmdArgs := append([]string{exe}, finalArgs...)
		entrypointArgs = append(entrypointArgs, cmdArgs...)
		err = syscall.Exec(exe, finalArgs, finalEnv)
		failFast(err)
	} else {
		// Run cmd
		fmt.Println("init: cmd: exe =", exe, "args =", finalArgs, "env =", finalEnv)
		err = syscall.Exec(exe, finalArgs, finalEnv)
		failFast(err)
	}
}

func failFast(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
