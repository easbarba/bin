/// 2>/dev/null ; gorun "$0" "$@" ; exit $?

package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type shotter struct {
	exec    string
	full    string
	partial string
	common  string
	name    string
}

func main() {
	full, partial := parse()
	cmd_full, cmd_partial := chooser()
	if *full == true {
		system(cmd_full)
	}

	if *partial == true {
		system(cmd_partial)
	}

}

func chooser() (string, string) {
	shotter_full := []string{grim().exec, grim().common, grim().full, grim().name}
	shotter_partial := []string{grim().exec, grim().common, grim().partial, grim().name}

	return strings.Join(shotter_full, " "), strings.Join(shotter_partial, " ")
}

func shot_location() string {
	home, err := os.UserHomeDir()

	if err != nil {
		panic(err)
	}

	screen_folder := filepath.Join(home, "Pictures", "screenshots")

	if _, err := os.Stat(screen_folder); err != nil {
		err = os.Mkdir(screen_folder, 0755)

		if err != nil {
			panic(err)
		}
	}

	return screen_folder
}

func grim() shotter {
	shot_name := filepath.Join(shot_location(), time.Now().Format(time.RFC3339)+".png")

	return shotter{
		exec:    "grimshot",
		full:    "active",
		partial: "area",
		common:  "save",
		name:    shot_name,
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
	full := flag.Bool("full", false, "full screen shot")
	partial := flag.Bool("partial", false, "partial screen shot")

	flag.Parse()

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: volume [options]\n\n")
		flag.PrintDefaults()
	}

	if *full == false && *partial == false {
		flag.Usage()
		os.Exit(0)
	}

	return full, partial
}
