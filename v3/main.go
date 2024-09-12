package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("I dont know what to do")
	}
}

func run() {
	fmt.Printf("[run] running %v as PID: %d \n", os.Args[2:], os.Getpid())

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func child() {
	fmt.Printf("starting container \n")
	fmt.Printf("[child] running %v as PID: %d \n", os.Args[2:], os.Getpid())

	must(syscall.Sethostname([]byte("inception")))
	must(syscall.Chroot("/root/container-fs/diff")) // put path to proper linux fs here
	must(syscall.Chdir("/"))
	must(syscall.Mount("proc", "proc", "proc", 0, ""))

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	must(cmd.Run())

	fmt.Printf("quitting container \n")
}

func must(err error) {
	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}
}
