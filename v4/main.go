package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("type command!")
		os.Exit(1)
	}

	cmd := os.Args[1]
	switch cmd {
	case "run":
		run()
	case "child":
		child()
	default:
		fmt.Println("wrong command!")
		os.Exit(1)
	}
}

func run() {
	fmt.Println("pid", os.Getpid())
	fmt.Println("args", os.Args[2:])

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}

	cmd.Run()
}

func child() {
	fmt.Println("hello from container")
	fmt.Println("pid", os.Getpid())
	fmt.Println("args", os.Args[2:])

	cgroup()

	must(syscall.Sethostname([]byte("inception")))
	must(syscall.Chroot("/root/container-fs/diff"))
	must(syscall.Chdir("/"))
	syscall.Mount("proc", "proc", "proc", 0, "")

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	must(cmd.Run())

	syscall.Unmount("/proc", 0)
}

func cgroup() {
	cgroup := "/sys/fs/cgroup/inception"
	must(os.MkdirAll(cgroup, 0755))

	must(os.WriteFile(cgroup+"/pids.max", []byte("10"), 0700))
	must(os.WriteFile(cgroup+"/cgroup.procs", []byte(fmt.Sprintf("%d", os.Getpid())), 0700))

	must(os.WriteFile(cgroup+"/memory.max", []byte("52428800"), 0700))
	must(os.WriteFile(cgroup+"/cgroup.procs", []byte(fmt.Sprintf("%d", os.Getpid())), 0700))
}

func must(err error) {
	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}
}
