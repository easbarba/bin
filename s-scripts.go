package main

import (
	"fmt"
	"os"
	"os/exec"
)

func system(cmd string, args []string) {
	cm := exec.Command(cmd, args)

	err := cm.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	stdout, err := cm.Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	fmt.Println(string(stdout))
}
