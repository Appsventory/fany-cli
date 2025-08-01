package app

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/Appsventory/fany-cli/internal/utils"
)

func GitInstall() error {
	if runtime.GOOS == "windows" {
		return fmt.Errorf("Windows: run 'winget install Git.Git' or download from git-scm.com")
	}
	cmd := exec.Command("sh", "-c", `
		if command -v apt >/dev/null; then sudo apt update && sudo apt install -y git;
		elif command -v dnf >/dev/null; then sudo dnf install -y git;
		elif command -v brew >/dev/null; then brew install git;
		else echo "Please install Git manually"; exit 1; fi
	`)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	return cmd.Run()
}

func ComposerInstall() error {
	if runtime.GOOS == "windows" {
		return fmt.Errorf("Windows: run 'winget install Composer.Composer' or use Composer-Setup.exe")
	}
	cmd := exec.Command("sh", "-c", `
		if command -v composer >/dev/null; then composer self-update;
		elif command -v brew >/dev/null; then brew install composer;
		else php -r "copy('https://getcomposer.org/installer','composer-setup.php');" && \
		     php composer-setup.php --install-dir=$HOME/.local/bin --filename=composer;
		fi
	`)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	return cmd.Run()
}

func ComposerUpgrade() error {
	if _, err := exec.LookPath("composer"); err != nil {
		return fmt.Errorf("composer not found, run: fany cp-install")
	}
	cmd := exec.Command("composer", "self-update")
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	return cmd.Run()
}

func CpInit() error {
	if _, err := exec.LookPath("composer"); err != nil {
		return fmt.Errorf("composer not found, run: fany cp-install")
	}
	cmd := exec.Command("composer", "install")
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	return cmd.Run()
}

func GitInit() error {
	if !utils.IsGitInstalled() {
		return fmt.Errorf("git not found, run: fany git-install")
	}
	steps := [][]string{
		{"git", "init"},
		{"git", "add", "."},
		{"git", "commit", "-m", "Initial commit"},
	}
	for _, step := range steps {
		cmd := exec.Command(step[0], step[1:]...)
		cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
		if err := cmd.Run(); err != nil {
			return err
		}
	}
	return nil
}
