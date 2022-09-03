package main

import (
	"bytes"
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

	out, _ := exec.Command("whoami").Output()
	fmt.Println("init: whoami =", string(out))

	// Run entrypoint
	if entrypoint != "" {
		cmdArgs := append([]string{exe}, finalArgs...)
		entrypointArgs = append(entrypointArgs, cmdArgs...)
		fmt.Println("init: entrypoint: exe =", entrypoint, "args =", entrypointArgs, "env =", finalEnv)
		entrypointExec := exec.Command(entrypoint, entrypointArgs...)
		entrypointExec.Env = finalEnv

		var outb, errb bytes.Buffer
		entrypointExec.Stdout = &outb
		entrypointExec.Stderr = &errb
		err := entrypointExec.Run()
		fmt.Println("out:", outb.String(), "err:", errb.String())
		failFast(err, "entrypoint exec")
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
