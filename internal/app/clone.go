package app

import (
	"github.com/Appsventory/fany-cli/internal/utils"
)

func GitClone(repo string, folder string) error {
	if folder == "" {
		folder = repo
	}
	url := "https://github.com/" + repo + ".git"
	return utils.GitClone(url, folder)
}
