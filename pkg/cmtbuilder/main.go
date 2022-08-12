package cmtbuilder

import (
	"fmt"
	"strings"

	"github.com/corentindeboisset/styledconsole"
)

func PromptCommit() (string, string, error) {
	// The conditions of the format are specified in the conventionnal commits standard, here:
	// https://www.conventionalcommits.org/en/v1.0.0/

	// The specs from commitlint are used as a reference, they can be found here:
	// https://github.com/conventional-changelog/commitlint/tree/master/@commitlint/config-conventional

	var commitType string
	choices := []string{
		"feat  - A new feature",
		"fix   - A bug fix",
		"docs  - Documentation only changes",
		"perf  - A code change that improves performance",
		"chore - A catch-all type for any other commits",
	}
	res, err := styledconsole.Choice("Select the type of change that you are committing", choices)
	if err != nil {
		return "", "", fmt.Errorf("failed to pick the type of commit: %w", err)
	}

	for i := 0; i < len(choices); i++ {
		if res == choices[i] {
			commitType = strings.TrimSpace(res[0:6])
		}
	}

	scopeName, err := styledconsole.Ask(
		"What is the scope of this change in the codebase",
		func(val string) bool {
			if len(commitType)+len(val) > 97 {
				fmt.Printf("The specified scope is too long\n")
				return false
			}
			return true
		},
	)
	if err != nil {
		return "", "", fmt.Errorf("failed to prompt the user for the scope of the commit: %w", err)
	}
	scopeName = strings.ToLower(strings.Trim(scopeName, " .\n\t\r\v\f"))
	if len(scopeName) > 0 {
		scopeName = fmt.Sprintf("(%s)", scopeName)
	}

	shortMsg, err := styledconsole.Ask(
		"Write a short, imperative tense description of the change",
		func(val string) bool {
			if len(commitType)+len(scopeName)+len(val) > 97 {
				fmt.Printf("The specified description is too long\n")
				return false
			}
			if len(val) == 0 {
				fmt.Printf("The specified description is too short\n")
				return false
			}
			return true
		},
	)
	if err != nil {
		return "", "", fmt.Errorf("failed to prompt the user for a short description of the commit: %w", err)
	}
	// Trim the message and force a lowercase first caracter
	shortMsg = strings.Trim(shortMsg, " .\n\t\r\v\f")
	shortMsg = strings.ToLower(shortMsg[0:1]) + shortMsg[1:]

	largeMsg, err := styledconsole.Ask("Provide a longer description of the change", nil)
	if err != nil {
		return "", "", fmt.Errorf("failed to prompt the user for a long description of the commit: %w", err)
	}

	breakingChange, err := styledconsole.ConfirmWithDefault("Are there any breaking changes? ", false)
	if err != nil {
		return "", "", fmt.Errorf("failed to prompt the user whether there are breaking changes in the commit: %w", err)
	}

	var title, body string
	if breakingChange {
		title = fmt.Sprintf("%s%s!: %s", commitType, scopeName, shortMsg)
		breakingChangeMessage, err := styledconsole.Ask("Describe the breaking changes", nil)
		if err != nil {
			return "", "", fmt.Errorf("failed to prompt the user for a description of the breaking changes in the commit: %w", err)
		}
		body = fmt.Sprintf("%s\n\nBREAKING CHANGE: %s", largeMsg, breakingChangeMessage)
	} else {
		title = fmt.Sprintf("%s%s: %s", commitType, scopeName, shortMsg)
		body = largeMsg
	}

	// TODO: add a question to explicit issues closed by the commit

	return title, limitLines(body), nil
}

func limitLines(paragraph string) string {
	lines := strings.Split(paragraph, "\n")
	limitedLines := make([]string, 0)

	for _, line := range lines {
		for {
			if len(line) > 99 {
				// Shave off the beginning of the line, and iterate until it's short enough
				lastSpaceIdx := strings.LastIndex(line[0:100], " ")
				if lastSpaceIdx != -1 {
					limitedLines = append(limitedLines, line[0:lastSpaceIdx])
					line = line[lastSpaceIdx:]
				} else {
					limitedLines = append(limitedLines, line[0:100]+"-")
					line = "-" + line[100:]
				}
				continue
			} else {
				limitedLines = append(limitedLines, line)
				break
			}
		}
	}

	return strings.Join(limitedLines, "\n")
}
