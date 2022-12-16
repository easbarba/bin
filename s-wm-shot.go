/// 2>/dev/null ; gorun "$0" "$@" ; exit $?

package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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
	cfull, cpartial := cli_parse()
	a_full, a_partial := action_picker()

	if *cfull == true {
		system(a_full[0], a_full[:1]...)
	}

	if *cpartial == true {
		system(a_partial[0], a_partial[:1]...)
	}

}

func action_picker() ([]string, []string) {
	full := []string{grim().exec, grim().common, grim().full, grim().name}
	partial := []string{grim().exec, grim().common, grim().partial, grim().name}

	return full, partial
}

func shot_location() string {
	home, err := os.UserHomeDir()

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	screen_folder := filepath.Join(home, "Pictures", "screenshots")

	if _, err := os.Stat(screen_folder); err != nil {
		err = os.Mkdir(screen_folder, 0755)

		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
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
