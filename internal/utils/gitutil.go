package utils

import (
	"os"
	"os/exec"
)

func GitClone(repoURL, dest string) error {
	cmd := exec.Command("git", "clone", repoURL, dest)
	cmd.Stdout = nil // progress git ke stderr kita
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func IsGitInstalled() bool {
	_, err := exec.LookPath("git")
	return err == nil
}
