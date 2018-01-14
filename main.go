package main

import (
	"bufio"
	"fmt"
	"os"

	"os/exec"

	"github.com/muziyoshiz/shellflow/parser"
	flag "github.com/spf13/pflag"
)

var (
	dryRunOpt  = flag.Bool("dry-run", false, `dry run`)
	versionOpt = flag.BoolP("version", "v", false, `version`)
	helpOpt    = flag.BoolP("help", "h", false, `help`)
)

func main() {
	os.Exit(_main())
}

func _main() int {
	flag.Parse()

	if *helpOpt {
		flag.PrintDefaults()
		return 0
	}

	if *versionOpt {
		fmt.Printf("shellflow version: %s (commit: %s)\n", version, commit)
		return 0
	}

	parse := parser.Parser()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		_, err := parse(line)
		if err != nil {
			// Format error
			return 1
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading standard input: ", err)
	}

	commands, err := parse(parser.EOF)
	if err != nil {
		// Execute commands
		fmt.Print(err)
	}

	for num, cmd := range commands {
		fmt.Printf("# %d. %s\n", num+1, cmd.Description)
		fmt.Printf("%s\n", cmd.Command)
	}

	if *dryRunOpt {
		return 0
	}

	for num, cmd := range commands {
		fmt.Printf("# %d. %s\n", num+1, cmd.Description)
		fmt.Printf("%s\n", cmd.Command)

		out, err := exec.Command("sh", "-c", cmd.Command).Output() // #nosec
		if err != nil {
			return 1
		}

		fmt.Printf("%s", out)
	}

	return 0
}
