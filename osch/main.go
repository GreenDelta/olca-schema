package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	args := parseArgs()

	if args.command == "" ||
		args.command == "help" ||
		args.command == "-h" {
		printHelp()
		return
	}

	switch args.command {
	case "proto":
		proto(args)
	case "doc", "docs", "md", "mdbook", "markdown":
		writeMarkdownBook(args)
	case "py", "python":
		writePythonModule(args)
	case "context":
		writeContextJson(args)
	case "check":
		checkSchema(args)
	default:
		fmt.Println("unknown command:", args.command)
	}

}

func check(err error, msg ...interface{}) {
	if err != nil {
		fmt.Print("ERROR: ")
		fmt.Println(msg...)
		panic(err)
	}
}

func proto(args *args) {
	yamlModel, err := ReadYamlModel(args)
	check(err)
	proto := generateProto(yamlModel)
	buildDir := filepath.Join(args.home, "build")
	mkdir(buildDir)
	outFile := filepath.Join(buildDir, "olca.proto")
	err = os.WriteFile(outFile, []byte(proto), os.ModePerm)
	check(err, "failed to write to file", outFile)
}

func printHelp() {
	fmt.Println(`
osch

usage:

$ osch [command]

commands:

  help    - prints this help
  check   - checks the schema
	doc     - generates the schema documentation
  proto   - generates the Protocol Buffers schema
	py      - generates the Python classes
	context - writes the docs/context.jsonld file
  `)
}
