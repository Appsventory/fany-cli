package utils

import "github.com/fatih/color"

var (
	Red    = color.New(color.FgRed).SprintfFunc()
	Green  = color.New(color.FgGreen).SprintfFunc()
	Yellow = color.New(color.FgYellow).SprintfFunc()
	Blue   = color.New(color.FgBlue).SprintfFunc()
)
