package gitadapter

import (
	"fmt"
	"os"
	"os/exec"
)

func PrecommitHook() error {
	preCommitHookCmd := exec.Command("git", "hook", "run", "--ignore-missing", "pre-commit")
	preCommitHookCmd.Stdout = os.Stdout
	preCommitHookCmd.Stderr = os.Stderr

	if err := preCommitHookCmd.Run(); err != nil {
		return fmt.Errorf("The pre-commit hook failed: %w", err)
	}

	return nil
}

func CreateCommit(message, body string) error {
	// Then, prepare a commit file from the submitted message/body
	tmp, err := os.CreateTemp("", "COMMIT_FILE_*")
	if err != nil {
		return fmt.Errorf("failed to create a temp file to store the commit content: %w", err)
	}
	defer os.Remove(tmp.Name())

	if _, err := fmt.Fprintf(tmp, "%s\n%s", message, body); err != nil {
		return fmt.Errorf("failed to write the content of the commit to a temporary file: %w", err)
	}

	// Run the commit-msg hook
	commitMessageHookCmd := exec.Command("git", "hook", "run", "--ignore-missing", "commit-msg", "--", tmp.Name())
	commitMessageHookCmd.Stdout = os.Stdout
	commitMessageHookCmd.Stderr = os.Stderr
	if err := commitMessageHookCmd.Run(); err != nil {
		return fmt.Errorf("The commit-message hook failed: %w", err)
	}

	// Then, create the actual commit, skipping the two hooks that were already run thanks to --no-verify
	commitCmd := exec.Command("git", "commit", "-F", tmp.Name(), "--no-verify")
	commitCmd.Stdout = os.Stdout
	commitCmd.Stderr = os.Stderr
	if err := commitCmd.Run(); err != nil {
		return fmt.Errorf("The command git commit failed: %w", err)
	}

	return nil
}
