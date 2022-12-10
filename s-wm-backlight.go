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
	up   string
	down string
}

func main() {
	up, down := parse()
	brighter := brightnessctl()

	if *up == true {
		print(brighter.exec + " " + brighter.up)
	}

	if *down == true {
		system(brighter.exec + " " + brighter.down)
	}
}

func brightnessctl() brighter {
	step := 3

	return brighter{
		exec: "brightnessctl",
		up:   fmt.Sprintf("set %d%%+", step),
		down: fmt.Sprintf("set %d%%-", step),
	}
}

func system(cmd string) {
	cm := exec.Command("sh", "-c", cmd) // TODO
	stdout, err := cm.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(stdout))
}

// command line arguments parser
func parse() (*bool, *bool) {
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
