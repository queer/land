package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("starting init")
	exe := "%EXE%"
	finalArgs := []string{"%ARGS%"}

	fmt.Println("cmd prep")
	cmd := exec.Command(exe, finalArgs...)
	fmt.Println("cmd prerun")
	err := cmd.Run()
	fmt.Println("cmd run")

	if err != nil {
		fmt.Println("kaboom")
		fmt.Println(err.Error())
	}
}
