package app

import (
	"os"

	"github.com/Appsventory/fany-cli/internal/utils"
)

func CacheInit(src string) error {
	dst, err := utils.CacheDir()
	if err != nil {
		return err
	}
	if err := os.RemoveAll(dst); err != nil {
		return err
	}
	return utils.CopyDir(src, dst)
}

func CacheUpdate() error {
	cacheDir, err := utils.CacheDir()
	if err != nil {
		return err
	}
	repo := "https://github.com/Appsventory/NineVerse.git"
	// jika folder cache ada -> clone ulang
	if err := os.RemoveAll(cacheDir); err != nil {
		return err
	}
	return utils.GitClone(repo, cacheDir)
}
