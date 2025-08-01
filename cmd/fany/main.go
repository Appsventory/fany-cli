package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: fany <command>")
		fmt.Println("Try: fany --help")
		return
	}

	switch os.Args[1] {
	case "--v", "--version":
		fmt.Println("Fany CLI v0.1.0-dev")
	case "--help", "-h":
		fmt.Println(`Fany CLI – bootstrap NineVerse projects
Commands:
  new <name>         create project
  get-new <name>     force online download
  cache-init <path>  register cache
  cache-update       update cache
  install            install CLI to PATH
  uninstall          remove CLI
  git-clone <repo>   clone shorthand
  cp-install         install Composer
  cp-upgrade         upgrade Composer
`)
	default:
		fmt.Printf("unknown command: %s\n", os.Args[1])
	}
}