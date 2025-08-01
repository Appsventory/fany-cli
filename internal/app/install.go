package app

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"

	"golang.org/x/sys/windows/registry"
)

// Install menempelkan binary ke user PATH
func Install(global bool) error {
	binName := "fany"
	if runtime.GOOS == "windows" {
		binName = "fany.exe"
	}

	exe, err := os.Executable()
	if err != nil {
		return err
	}

	var targetDir string
	if global {
		targetDir = globalBinDir()
	} else {
		targetDir = userBinDir()
	}
	// di fungsi Install, setelah copy file selesai
	if runtime.GOOS == "windows" {
		if err := addUserPath(targetDir); err == nil {
			fmt.Printf("✔ Added %s to user PATH (restart terminal)\n", targetDir)
		}
	}
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return err
	}
	target := filepath.Join(targetDir, binName)
	if _, err := os.Stat(target); err == nil {
		return fmt.Errorf("already installed at %s", target)
	}

	in, err := os.Open(exe)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err := io.Copy(out, in); err != nil {
		return err
	}

	fmt.Printf("✔ Installed to %s\n", target)
	return nil
}

// (di akhir file, setelah fungsi Install)
func addUserPath(dir string) error {
	if runtime.GOOS != "windows" {
		return nil // no-op di non-Windows
	}
	k, err := registry.OpenKey(registry.CURRENT_USER, `Environment`, registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	defer k.Close()
	val, _, err := k.GetStringValue("PATH")
	if err != nil && err != registry.ErrNotExist {
		return err
	}
	for _, p := range filepath.SplitList(val) {
		if p == dir {
			return nil
		}
	}
	newVal := val + ";" + dir
	return k.SetStringValue("PATH", newVal)
}

// Uninstall menghapus binary
func Uninstall(global bool) error {
	binName := "fany"
	if runtime.GOOS == "windows" {
		binName = "fany.exe"
	}
	target := filepath.Join(binDir(global), binName)
	if err := os.Remove(target); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("not found at %s", target)
		}
		return err
	}
	fmt.Printf("✔ Removed %s\n", target)
	return nil
}

func binDir(global bool) string {
	if global {
		return globalBinDir()
	}
	return userBinDir()
}

func userBinDir() string {
	if runtime.GOOS == "windows" {
		return filepath.Join(os.Getenv("USERPROFILE"), "bin")
	}
	return filepath.Join(os.Getenv("HOME"), ".local", "bin")
}

func globalBinDir() string {
	if runtime.GOOS == "windows" {
		return `C:\Program Files\Fany\bin`
	}
	return "/usr/local/bin"
}
