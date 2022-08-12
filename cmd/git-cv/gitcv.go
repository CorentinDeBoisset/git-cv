package main

import (
	"fmt"
	"os"

	"github.com/corentindeboisset/git-cv/pkg/cmtbuilder"
	"github.com/corentindeboisset/git-cv/pkg/gitadapter"
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

	curDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("An error occured: %v\n", err)
		os.Exit(1)
		return
	}

	if err := gitadapter.CheckRepo(curDir); err != nil {
		fmt.Printf("An error occured: %v\n", err)
		os.Exit(1)
		return
	}

	// TODO: check there are files to commit (or that the allow-empty flag is on)

	if err := gitadapter.PrecommitHook(); err != nil {
		fmt.Printf("An error occured: %v\n", err)
		os.Exit(1)
		return
	}

	title, body, err := cmtbuilder.PromptCommit()
	if err != nil {
		fmt.Printf("An error occured: %v\n", err)
		os.Exit(1)
		return
	}

	// TODO: add the parsed arguments into the git command
	if err := gitadapter.CreateCommit(title, body); err != nil {
		fmt.Printf("An error occured: %v\n", err)
		os.Exit(1)
		return
	}
}
