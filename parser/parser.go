package parser

import (
	"fmt"
	"strings"
)

// EOF is a special string to teach the end of shell script to the parser
const EOF string = "EOF"

// ShellCommand is a shell command from stdin or file
type ShellCommand struct {
	Description string
	Command     string
}

// Parser returns a closure to parse shell script
func Parser() func(string) ([]ShellCommand, error) {
	var serial = 0
	var commands = make([]ShellCommand, 0, 10)

	var desc = ""
	var preLine = ""

	return func(line string) ([]ShellCommand, error) {
		if line == EOF {
			if preLine != "" {
				return commands, fmt.Errorf("no trailing line after `%s`", preLine)
			}

			return commands, nil
		}

		// Blank line is ignored

		// Comment line is recorded as Description
		if strings.HasPrefix(line, "#") {
			desc = line
			return commands, nil
		}

		if strings.HasSuffix(line, "\\") {
			preLine += line + "\n"
			return commands, nil
		}

		serial++
		commands = append(commands, ShellCommand{desc, preLine + line})

		// Reset temporary values
		desc, preLine = "", ""

		return commands, nil
	}
}
