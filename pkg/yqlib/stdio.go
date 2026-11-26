package yqlib

import (
	"io"
	"os"

	"github.com/mattn/go-isatty"
)

// Stdin is the stream yq reads for the "-" file argument and piped input.
// Embedders that run yq without process streams replace it before each run;
// the default reads whatever os.Stdin holds at read time.
var Stdin io.Reader = processStdin{}

// Stderr receives yq's log output; the default writes to os.Stderr at write time.
var Stderr io.Writer = processStderr{}

type processStdin struct{}

func (processStdin) Read(p []byte) (int, error) { return os.Stdin.Read(p) }

type processStderr struct{}

func (processStderr) Write(p []byte) (int, error) { return os.Stderr.Write(p) }

// StdinIsPiped reports whether input is piped, so yq reads it when no file is given.
var StdinIsPiped = func() bool {
	stat, err := os.Stdin.Stat()
	if err != nil {
		GetLogger().Debugf("error getting stdin: %v", err)
	}
	return stat != nil && (stat.Mode()&os.ModeCharDevice) == 0
}

// StdoutIsTerminal reports whether output goes to a terminal, which enables colours.
var StdoutIsTerminal = func() bool {
	fd := os.Stdout.Fd()
	return isatty.IsTerminal(fd) || isatty.IsCygwinTerminal(fd)
}

type currentStderr struct{}

func (currentStderr) Write(p []byte) (int, error) { return Stderr.Write(p) }

// StderrWriter writes to whichever writer Stderr holds at write time.
var StderrWriter io.Writer = currentStderr{}
