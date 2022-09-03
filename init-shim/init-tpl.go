package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("AAAAAAAAAAAAAAAA")
	exe := "%EXE%"
	finalArgs := []string{"%ARGS%"}

	cmd := exec.Command(exe, finalArgs...)
	err := cmd.Run()

	if err != nil {
		fmt.Println(err.Error())
	}
}
