package app

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/Appsventory/fany-cli/internal/utils"
)

// New membuat project baru: online dulu, fallback cache
func New(name, dir string, force bool) error {
	dest := filepath.Join(dir, name)
	if !force {
		if _, err := os.Stat(dest); err == nil {
			return os.ErrExist
		}
	} else {
		_ = os.RemoveAll(dest)
	}

	if utils.HasInternet() {
		repo := "https://github.com/Appsventory/NineVerse.git"
		if err := utils.GitClone(repo, dest); err != nil {
			return err
		}
	} else {
		cache, err := utils.CacheDir()
		if err != nil {
			return err
		}
		if _, err := os.Stat(cache); os.IsNotExist(err) {
			return err
		}
		if err := utils.CopyDir(cache, dest); err != nil {
			return err
		}
	}

	if err := dropWrapper(dest); err != nil {
		fmt.Println("warning: could not create wrapper:", err)
	}
	return patchEnv(name, dest)
}

// GetNew selalu download online (tanpa fallback)
func GetNew(name, dir string, force bool) error {
	dest := filepath.Join(dir, name)
	if !force {
		if _, err := os.Stat(dest); err == nil {
			return os.ErrExist
		}
	} else {
		_ = os.RemoveAll(dest)
	}
	repo := "https://github.com/Appsventory/NineVerse.git"
	if err := utils.GitClone(repo, dest); err != nil {
		return err
	}
	if err := dropWrapper(dest); err != nil {
		fmt.Println("warning: could not create wrapper:", err)
	}
	return patchEnv(name, dest)
}

// patchEnv: copy .env.example → .env lalu set APP_NAME & APP_VERSION
func patchEnv(name, dest string) error {
	envEx := filepath.Join(dest, ".env.example")
	envDst := filepath.Join(dest, ".env")

	data, err := os.ReadFile(envEx)
	if err != nil {
		return err
	}
	// copy
	if err := os.WriteFile(envDst, data, 0644); err != nil {
		return err
	}
	// edit
	content := string(data)
	content = strings.ReplaceAll(content, "APP_NAME=", "APP_NAME="+name)
	content = strings.ReplaceAll(content, "APP_VERSION=", "APP_VERSION=1.0")
	return os.WriteFile(envDst, []byte(content), 0644)
}

func dropWrapper(dest string) error {
	if runtime.GOOS == "windows" {
		return os.WriteFile(filepath.Join(dest, "fany.bat"), []byte(wrapperBat), 0644)
	}
	shFile := filepath.Join(dest, "fany")
	err := os.WriteFile(shFile, []byte(wrapperSh), 0755)
	if err == nil {
		// chmod +x
		return os.Chmod(shFile, 0755)
	}
	return err
}

// file wrapper literal
const wrapperSh = `#!/usr/bin/env bash
php "$(dirname "$0")/fany" "$@"
`

const wrapperBat = `@echo off
php "%~dp0fany" %*
`
