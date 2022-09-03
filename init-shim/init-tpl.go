package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func main() {
	exe := "%EXE%"
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

	// Run entrypoint
	if entrypoint != "" {
		cmdArgs := append([]string{exe}, finalArgs...)
		entrypointArgs = append(entrypointArgs, cmdArgs...)
		fmt.Println("init: entrypoint: exe =", entrypoint, "args =", []string{entrypointArgs[0]}, "env =", finalEnv)
		// entrypointExec := exec.Command(entrypoint, entrypointArgs...)
		// entrypointExec.Env = finalEnv
		// err := entrypointExec.Run()
		err := syscall.Exec(entrypoint, []string{entrypointArgs[0]}, finalEnv)
		failFast(err, "entrypointExec.Run()")
	} else {
		// Look up the exe since syscall.Exec doesn't
		exe, err := exec.LookPath(exe)
		failFast(err, "cmd exe resolve")
		// Run cmd
		fmt.Println("init: cmd: exe =", exe, "args =", finalArgs, "env =", finalEnv)
		err = syscall.Exec(exe, finalArgs, finalEnv)
		failFast(err, "cmd exec")
	}
}

func failFast(err error, ctx string) {
	if err != nil {
		fmt.Println(ctx, "-", err.Error())
		os.Exit(1)
	}
}
