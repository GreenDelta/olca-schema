package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

// Writes the `context.jsonld` file to the `docs` folder that is hosted via
// GitHub pages.
func writeContextJson(args *args) {

	model, err := ReadYamlModel(args)
	if err != nil {
		return
	}

	enums := make(map[string]bool)
	model.EachEnum(func(enum *YamlEnum) {
		enums[enum.Name] = true
	})

	enumFields := make(map[string]bool)
	model.EachClass(func(class *YamlClass) {
		for _, prop := range class.Props {
			if enums[prop.Type] {
				enumFields[prop.Name] = true
			}
		}
	})

	// build the context dictionary
	context := map[string]any{
		"@vocab": "http://greendelta.github.io/olca-schema#",
		// "@base":  "http://greendelta.github.io/olca-schema#",
	}
	vocab := map[string]string{"@type": "@vocab"}
	for field := range enumFields {
		context[field] = vocab
	}
	dict := map[string]any{
		"@context": context,
	}

	// serialize it to Json
	bytes, err := json.MarshalIndent(dict, "", "  ")
	if err != nil {
		log.Fatalln("ERROR: failed to serialize JSON-LD context:", err)
		return
	}

	// write it to the docs-folder
	docDir := filepath.Join(args.home, "docs")
	if _, err := os.Stat(docDir); err != nil {
		if err := os.MkdirAll(docDir, os.ModePerm); err != nil {
			log.Fatalln("ERROR: failed to create `docs` folder:", docDir, err)
			return
		}
	}

	path := filepath.Join(docDir, "context.jsonld")
	if err := os.WriteFile(path, bytes, os.ModePerm); err != nil {
		log.Fatalln("ERROR: failed to write context.jsonld:", err)
		return
	}
	log.Println("wrote docs/conext.jsonld")

}
