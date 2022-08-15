package main

import (
	"os"
	"path/filepath"
)

type args struct {
	command string
	home    string
}

func parseArgs() *args {
	args := &args{}
	osArgs := os.Args
	if len(osArgs) < 2 {
		return args
	}

	args.command = osArgs[1]
	args.home = findHome()
	return args
}

func findHome() string {
	dir, err := filepath.Abs(".")
	if err != nil {
		return ""
	}
	for {
		path := filepath.Join(dir, "olca-schema")
		if isDir(path) {
			return path
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return ""
		}
		dir = parent
	}
}
