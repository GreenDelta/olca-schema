package main

import (
	"fmt"
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
		writeProtos(args)
	case "doc", "docs", "md", "mdbook", "markdown":
		writeMarkdownBook(args)
	case "py", "python":
		writePythonModule(args)
	case "st", "tonel":
		writeTonelFiles(args)
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
