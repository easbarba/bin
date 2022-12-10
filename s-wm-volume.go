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
	toggle string
	up     string
	down   string
}

func main() {
	up, down, toggle := parse()
	manager := chooser()

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

func system(cmd string, action string) {
	cm := exec.Command("sh", "-c", cmd+" "+action) // TODO
	stdout, err := cm.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(stdout))
}

func chooser() manager {
	manager := pactl()

	if _, err := os.Stat(manager.exec); err != nil {
		fmt.Print(err)
	}

	return manager
}

func pactl() manager {
	step := 3
	id := "@DEFAULT_SINK@"

	result := manager{
		exec:   "/usr/bin/pactl",
		toggle: fmt.Sprintf("set-sink-mute %s toggle", id),
		up:     fmt.Sprintf("set-sink-volume %s +%d%%", id, step),
		down:   fmt.Sprintf("set-sink-volume %s -%d%%", id, step),
	}

	return result
}

// command line arguments parser
func parse() (*bool, *bool, *bool) {
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
