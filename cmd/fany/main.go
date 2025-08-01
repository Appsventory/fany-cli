package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/Appsventory/fany-cli/internal/app"
	"github.com/Appsventory/fany-cli/internal/utils"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}
	if proxyToProject() {
		return // sudah diteruskan ke PHP
	}

	switch os.Args[1] {
	case "--v", "--version":
		printVersion()
	case "--help", "-h":
		printHelp()
	case "new":
		if len(os.Args) < 3 {
			printHelp()
			return
		}
		name, dir, force := parseArgs(os.Args[2:])
		if err := app.New(name, dir, force); err != nil {
			fmt.Println("error:", err)
		}
	case "get-new":
		if len(os.Args) < 3 {
			printHelp()
			return
		}
		name, dir, force := parseArgs(os.Args[2:])
		if err := app.GetNew(name, dir, force); err != nil {
			fmt.Println("error:", err)
		}
	case "git-clone":
		if len(os.Args) < 3 {
			printHelp()
			return
		}
		repo := os.Args[2]
		folder := ""
		if len(os.Args) > 3 {
			folder = os.Args[3]
		}
		if !utils.IsGitInstalled() {
			fmt.Println("Git not found. Run: fany git-install")
			return
		}
		if err := app.GitClone(repo, folder); err != nil {
			fmt.Println("error:", err)
		}
	case "git-install":
		if err := app.GitInstall(); err != nil {
			fmt.Println("git-install:", err)
		}
	case "git-init":
		if err := app.GitInit(); err != nil {
			fmt.Println("git-init:", err)
		}
	case "cp-install":
		if err := app.ComposerInstall(); err != nil {
			fmt.Println("cp-install:", err)
		}
	case "cp-upgrade":
		if err := app.ComposerUpgrade(); err != nil {
			fmt.Println("cp-upgrade:", err)
		}
	case "cp-init":
		if err := app.CpInit(); err != nil {
			fmt.Println("cp-init:", err)
		}
	case "cache-init":
		if len(os.Args) < 3 {
			printHelp()
			return
		}
		if err := app.CacheInit(os.Args[2]); err != nil {
			fmt.Println("error:", err)
		}
	case "cache-update":
		if err := app.CacheUpdate(); err != nil {
			fmt.Println("error:", err)
		}
	case "install":
		global := false
		for _, a := range os.Args[2:] {
			if a == "--global" {
				global = true
			}
		}
		if err := app.Install(global); err != nil {
			fmt.Println("install:", err)
		}
	case "uninstall":
		global := false
		for _, a := range os.Args[2:] {
			if a == "--global" {
				global = true
			}
		}
		if err := app.Uninstall(global); err != nil {
			fmt.Println("uninstall:", err)
		}
	default:
		printHelp()
	}
}

func printVersion() {
	fmt.Println("Fany CLI v0.1.0-dev")
}

func printHelp() {
	fmt.Println(`Usage: fany <command> [args...]
Commands:
  new <name> [--dir <path>] [--force]    create project
  get-new <name> [...]                   force online download
  git-clone <repo> [folder]              clone repo
  cache-init <path>                      register cache
  cache-update                           update cache
  --v, --help                            info
`)
}

// proxyToProject dijalankan hanya jika cwd berisi file "fany" (PHP CLI)
func proxyToProject() bool {
	// cek apakah ada fany PHP di cwd
	phpFile := "fany"
	if runtime.GOOS == "windows" {
		phpFile = "fany"
	}
	if _, err := os.Stat(phpFile); err == nil {
		// jalankan php fany <arg...>
		cmd := exec.Command("php", append([]string{phpFile}, os.Args[1:]...)...)
		cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
		_ = cmd.Run()
		return true
	}
	return false
}

func parseArgs(args []string) (name, dir string, force bool) {
	name = args[0]
	dir = "."
	for i := 1; i < len(args); i++ {
		switch args[i] {
		case "--dir":
			if i+1 < len(args) {
				dir = args[i+1]
			}
		case "--force":
			force = true
		}
	}
	return
}
