package main

import (
	"fmt"
	"io/ioutil"
	"os"
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
	case "md", "mdbook", "markdown":
		writeMarkdownBook(args)
	case "py", "python":
		writePythonModule(args)
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
	yamlModel, err := ReadYamlModel(args.yamlDir)
	check(err)

	proto := GenProto(yamlModel)

	// print to console or write to file
	if len(os.Args) < 3 {
		fmt.Println(proto)
	} else {
		outFile := os.Args[2]
		err := ioutil.WriteFile(outFile, []byte(proto), os.ModePerm)
		check(err, "failed to write to file", outFile)
	}
}

func printHelp() {
	fmt.Println(`
osch

usage:

$ osch [command] [options]

commands:

  help  - prints this help
  check - checks the schema
  proto - converts the schema to ProtocolBuffers

  `)
}
