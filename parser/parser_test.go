package parser_test

import (
	"testing"

	"github.com/muziyoshiz/shellflow/parser"
	"strings"
)

// Works with one-liner
func TestParseWithOneLiner(t *testing.T) {
	line := "echo \"Hello world\""

	parse := parser.Parser()

	commands, err := parse(line)
	if err != nil {
		t.Fatalf("Expected No error; got %+v", err)
	}

	if len(commands) != 1 {
		t.Fatalf("Expected 1; got %d", len(commands))
	}

	cmd := commands[0]
	if cmd.Number != 1 {
		t.Fatalf("Expected 1; got %d", cmd.Number)
	}
	if cmd.Description != "" {
		t.Fatalf("Expected blank; got %s", cmd.Description)
	}
	if cmd.Command != line {
		t.Fatalf("Expected %s; got %s", line, cmd.Command)
	}
}

// Works with multi-line command
func TestParseWithMultiLineCommand(t *testing.T) {
	str := `echo \
"Hello world"`

	lines := strings.Split(str, "\n")
	if len(lines) != 2 {
		t.Fatalf("Failed to create test data: length = %d", len(lines))
	}

	parse := parser.Parser()

	commands, err := parse(lines[0])
	if len(commands) != 0 {
		t.Fatalf("Expected 0; got %d", len(commands))
	}
	if err != nil {
		t.Fatalf("Expected No error; got %+v", err)
	}

	commands, err = parse(lines[1])
	if len(commands) != 1 {
		t.Fatalf("Expected 1; got %d", len(commands))
	}

	cmd := commands[0]
	if cmd.Number != 1 {
		t.Fatalf("Expected 1; got %d", cmd.Number)
	}
	if cmd.Description != "" {
		t.Fatalf("Expected blank; got %s", cmd.Description)
	}
	if cmd.Command != str {
		t.Fatalf("Expected %s; got %s", str, cmd.Command)
	}
}
