package main

import (
	"os"
	"path/filepath"
	"strings"
)

type args struct {
	command string
	yamlDir string
	target  string
}

func parseArgs() *args {
	args := &args{}
	osArgs := os.Args
	if len(osArgs) < 2 {
		return args
	}

	args.command = osArgs[1]
	flag := ""
	for i := 2; i < len(osArgs); i++ {
		arg := osArgs[i]
		if strings.HasPrefix(arg, "-") {
			flag = arg
			continue
		}
		if flag == "" {
			continue
		}
		switch flag {
		case "-i", "-s", "-input", "-schema":
			args.yamlDir = arg
		case "-o", "-output":
			args.target = arg
		}
	}

	if args.yamlDir == "" {
		// try to find the schema home by going up the directory tree
		schemaHome := findSchemaHome()
		if schemaHome != "" {
			args.yamlDir = filepath.Join(schemaHome, "yaml")
		} else {
			args.yamlDir = "."
		}
	}

	return args
}

func findSchemaHome() string {
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
