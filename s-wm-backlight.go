/// 2>/dev/null ; gorun "$0" "$@" ; exit $?

package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

type brighter struct {
	exec string
	up   []string
	down []string
}

func main() {
	up, down := cli_parse()
	brighter := brightnessctl()

	if *up == true {
		system(brighter.exec, brighter.up...)
	}

	if *down == true {
		system(brighter.exec, brighter.down...)
	}
}

func brightnessctl() brighter {
	step := 3

	return brighter{
		exec: "brightnessctl",
		up:   []string{"set", fmt.Sprintf("%d%%+", step)},
		down: []string{"set", fmt.Sprintf("%d%%-", step)},
	}
}

func system(cmd string, args ...string) {
	cm := exec.Command(cmd, args...)

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

// command line arguments parser
func cli_parse() (*bool, *bool) {
	up := flag.Bool("up", false, "turn backlight up")
	down := flag.Bool("down", false, "turn backlight down")

	flag.Parse()

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: volume [options]\n\n")
		flag.PrintDefaults()
	}

	if *up == false && *down == false {
		flag.Usage()
		os.Exit(0)
	}

	return up, down
}
