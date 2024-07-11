package main

import (
	"os"
	"path/filepath"
	"strings"
)

type args struct {
	command string
	home    string
	output  string
}

func parseArgs() *args {
	args := &args{
		home: findHome(),
	}
	osArgs := os.Args
	if len(osArgs) < 2 {
		return args
	}

	args.command = osArgs[1]
	flag := ""
	for _, arg := range osArgs {
		if strings.HasPrefix(arg, "-") {
			flag = arg
			continue
		}
		switch flag {
		case "-o", "-output":
			args.output = arg
		}
		flag = ""
	}
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

func (a *args) outputFileOrDefault(name string) (string, error) {
	var path string
	if a.output != "" {
		path = a.output
	} else {
		path = filepath.Join(a.home, name)
	}
	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		return "", err
	}
	return path, nil
}
