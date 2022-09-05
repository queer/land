package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/vishvananda/netlink"
)

func main() {
	fmt.Println("init: starting up!")
	exe := "%EXE%"
	finalArgs := []string{"%ARGS%"}
	finalEnv := []string{"%ENV%"}
	entrypoint := "%ENTRYPOINT%"
	entrypointArgs := []string{"%ENTRYPOINT_ARGS%"}

	// Set up initial path if provided
	for _, arg := range finalEnv {
		if strings.HasPrefix(arg, "PATH=") {
			path := arg[5:]
			os.Setenv("PATH", path)
			// Trick the shell into being usable if you boot one
			fmt.Println("init: set PATH =", path)
			break
		}
	}

	fmt.Println("init: enabling eth0")
	eth0, err := netlink.LinkByName("eth0")
	failFast(err, "netlink get eth0")
	err = netlink.LinkSetUp(eth0)
	failFast(err, "netlink set eth0 up")

	fmt.Println("init: adding eth0 ip")
	addr, err := netlink.ParseAddr("172.22.0.2/16")
	addr.Broadcast = net.IPv4(172, 22, 255, 255)
	failFast(err, "netlink parse addr")
	err = netlink.AddrAdd(eth0, addr)
	failFast(err, "netlink add addr")

	fmt.Println("init: adding default route")
	route := &netlink.Route{
		Scope: netlink.SCOPE_UNIVERSE,
		Dst:   nil,
		Gw:    net.IPv4(172, 22, 0, 1),
	}
	err = netlink.RouteAdd(route)
	failFast(err, "netlink add route")

	// Run entrypoint
	if entrypoint != "" {
		cmdArgs := append([]string{exe}, finalArgs...)
		entrypointArgs = append(entrypointArgs, cmdArgs...)
		fmt.Println("init: entrypoint: exe =", entrypoint, "args =", entrypointArgs, "env =", finalEnv)
		// entrypointExec := exec.Command(entrypoint, entrypointArgs...)
		// entrypointExec.Env = finalEnv

		// err = entrypointExec.Run()
		err = syscall.Exec(entrypoint, entrypointArgs, finalEnv)
		failFast(err, "entrypoint exec")
	} else {
		// Look up the exe since syscall.Exec doesn't
		exe, err := exec.LookPath(exe)
		failFast(err, "cmd exe resolve")

		// Run cmd
		fmt.Println("init: cmd: exe =", exe, "args =", finalArgs, "env =", finalEnv)
		err = syscall.Exec(exe, append([]string{exe}, finalArgs...), finalEnv)
		failFast(err, "cmd exec")
	}
}

func failFast(err error, ctx string) {
	if err != nil {
		fmt.Println("init:", ctx, "--", err.Error())
		os.Exit(1)
	}
}
