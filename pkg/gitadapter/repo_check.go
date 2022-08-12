package gitadapter

import (
	"fmt"
	"os/exec"
	"strings"
)

func CheckRepo(path string) error {
	cmd := exec.Command("git", "rev-parse")
	output, err := cmd.CombinedOutput()
	if err != nil {
		if _, ok := err.(*exec.ExitError); ok && strings.Contains(string(output), "not a git repository") {
			return fmt.Errorf("The specified path is not a git repository")
		}

		// There was an unknown error
		return fmt.Errorf("An error occured checking if the directory is a git repository: %w", err)
	}

	return nil
}
