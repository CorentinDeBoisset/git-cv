package main

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

var (
	command *pflag.FlagSet

	help     bool
	dryRun   bool
	noVerify bool

	author string
	date   string

	amend       bool
	resetAuthor bool

	allEdits bool
)

func buildCommand() *pflag.FlagSet {
	command = pflag.NewFlagSet(os.Args[0], pflag.ContinueOnError)

	command.BoolVarP(&allEdits, "all", "a", false, "")
	command.BoolVar(&resetAuthor, "reset-author", false, "")
	command.StringVar(&author, "author", "", "Override the commit author")
	command.StringVar(&date, "date", "", "Override the author date used in the commit.")
	command.BoolVarP(&noVerify, "no-verify", "n", false, "This option bypasses the pre-commit hook.")
	command.BoolVar(&amend, "amend", false, "")
	command.BoolVar(&dryRun, "dry-run", false, "")
	command.BoolVarP(&help, "help", "h", false, "Get some help")

	return command
}

func main() {
	cmd := buildCommand()

	if err := cmd.Parse(os.Args[1:]); err != nil || help {
		// Print usage and stop
		fmt.Fprint(os.Stderr, "usage: git cv [options]\n\n")
		cmd.PrintDefaults()

		if err != nil {
			os.Exit(2)
			return
		}
		return
	}

	fmt.Print("all right\n")

	// Use the parsed arguments
}
