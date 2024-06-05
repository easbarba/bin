/// 2>/dev/null ; gorun "$0" "$@" ; exit $?

package main

import (
	"flag"
	"fmt"
	"os"
)

type brighter struct {
	exec string
	up   []string
	down []string
}

func main() {
	up, down := bright_cli_parse()
	brighter := brightnessctl()

	if *up == true {
		system(brighter.exec, brighter.up)
	}

	if *down == true {
		system(brighter.exec, brighter.down)
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

// command line arguments parser
func bright_cli_parse() (*bool, *bool) {
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
