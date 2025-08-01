package utils

import (
	"os"
	"path/filepath"
)

func CacheDir() (string, error) {
	base := os.Getenv("HOME")
	if base == "" {
		base = os.Getenv("USERPROFILE") // Windows
	}
	dir := filepath.Join(base, ".cache", "fany", "nineverse")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", err
	}
	return dir, nil
}
