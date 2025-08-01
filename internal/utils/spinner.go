package utils

import (
	"os"

	"github.com/schollz/progressbar/v3"
)

func NewBar(max int64) *progressbar.ProgressBar {
	return progressbar.NewOptions64(max,
		progressbar.OptionSetWriter(os.Stderr),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetDescription("cloning"),
		progressbar.OptionSetTheme(progressbar.Theme{Saucer: "█", SaucerPadding: "░", BarStart: "▕", BarEnd: "▏"}),
	)
}
