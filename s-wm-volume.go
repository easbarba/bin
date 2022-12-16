/// 2>/dev/null ; gorun "$0" "$@" ; exit $?

package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

type manager struct {
	exec   string
	toggle []string
	up     []string
	down   []string
}

func main() {
	up, down, toggle := cli_parse()
	manager := picker()

	if *up == true {
		system(manager.exec, manager.up)
	}

	if *down == true {
		system(manager.exec, manager.down)
	}

	if *toggle == true {
		system(manager.exec, manager.toggle)
	}
}

func system(cmd string, action []string) {
	cm := exec.Command(cmd, action...)

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

func picker() manager {
	manager := pactl()

	if _, err := os.Stat(manager.exec); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	return manager
}

func pactl() manager {
	step := 2
	id := fmt.Sprintf("%s", "@DEFAULT_SINK@")

	result := manager{
		exec:   "/usr/bin/pactl",
		toggle: []string{"set-sink-mute", id, "toggle"},
		up:     []string{"set-sink-volume", id, fmt.Sprintf("+%d%%", step)},
		down:   []string{"set-sink-volume", id, fmt.Sprintf("-%d%%", step)},
	}

	return result
}

// command line arguments parser
func cli_parse() (*bool, *bool, *bool) {
	up := flag.Bool("up", false, "increase volume")
	down := flag.Bool("down", false, "Decrease volume")
	toggle := flag.Bool("toggle", false, "Toggle volume")

	flag.Parse()

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: volume [options]\n\n")
		flag.PrintDefaults()
	}

	if *up == false && *down == false && *toggle == false {
		flag.Usage()
		os.Exit(0)
	}

	return up, down, toggle
}
