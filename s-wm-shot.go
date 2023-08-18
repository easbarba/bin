/// 2>/dev/null ; gorun "$0" "$@" ; exit $?

package main

import (
	"flag"
	"fmt"
	"os"
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

func (s *Shotter) end(partial bool) []string {
	if partial {
		return []string{s.common + s.partial + s.name}
	}

	return []string{s.common + s.full + s.name}
}

func main() {
	_, partial := shot_cli_parse()

	system(grim().exec, grim().end(partial))
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
