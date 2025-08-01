package app

func ParseDirForce(args []string) (dir string, force bool) {
	dir = "."
	for i := 0; i < len(args); i++ {
		if args[i] == "--dir" && i+1 < len(args) {
			dir = args[i+1]
		}
		if args[i] == "--force" {
			force = true
		}
	}
	return
}
