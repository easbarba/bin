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

type Shotter struct {
	exec    string
	full    string
	partial string
	common  string
	name    string
}

func (s *Shotter) full_end() []string {
	return []string{s.common, s.full, s.name}
}

func (s *Shotter) partial_end() []string {
	return []string{s.common, s.partial, s.name}
}

func main() {
	full, partial := shot_cli_parse()
	grim := grim()

	if *full == true {
		shot_command(grim.exec, grim.full_end())
	}

	if *partial == true {
		shot_command(grim.exec, grim.partial_end())
	}
}

func shot_location() string {
	home, err := os.UserHomeDir()

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	screen_folder := filepath.Join(home, "Pictures", "Screenshots")

	if _, err := os.Stat(screen_folder); err != nil {
		err = os.Mkdir(screen_folder, 0755)

		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	}

	return screen_folder
}

func grim() Shotter {
	shot_name := filepath.Join(shot_location(), time.Now().Format(time.RFC3339)+".png")

	return Shotter{
		exec:    "grimshot",
		full:    "active",
		partial: "area",
		common:  "save",
		name:    shot_name,
	}
}

// command line arguments parser
func shot_cli_parse() (*bool, *bool) {
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

func shot_command(cmd string, args []string) {
	command := exec.Command(cmd, args...)

	if err := command.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	if stdout, err := command.Output(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	} else {
		fmt.Println(string(stdout))
	}
}
