package main

import (
	"bytes"
	"fmt"
	"os"
)

type gow struct {
	buff  *bytes.Buffer
	model *YamlModel
}

func (w *gow) buffer() *bytes.Buffer {
	return w.buff
}

func (w *gow) indent() string {
	return "\t"
}

func writeGo(args *args) error {

	model, err := ReadYamlModel(args)
	if err != nil {
		return fmt.Errorf("failed to read YAML model: %w", err)
	}

	var buffer bytes.Buffer
	w := gow{&buffer, model}
	w.header()
	w.enums()

	out, err := args.outputFileOrDefault("/build/golcas/schema.go")
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	if err := os.WriteFile(out, buffer.Bytes(), os.ModePerm); err != nil {
		return fmt.Errorf("failed to write file %s: %w", out, err)
	}
	return nil
}

func (w *gow) header() {
	ln(w, "// DO NOT CHANGE THIS FILE AS THIS WAS GENERATED AUTOMATICALLY")
	ln(w, "// instead modify the generator that produces this file, see")
	ln(w, "// http://greendelta.github.io/olca-schema")
	ln(w)
	ln(w, "package golcas")
	ln(w)
}

func (w *gow) enums() {

	w.model.EachEnum(func(e *YamlEnum) {
		ln(w, "type ", e.Name, " string")
		ln(w, "const(")
		for _, i := range e.Items {
			lni(w, 1, i.Name, " ", e.Name, " = ", "\""+i.Name+"\"")
		}
		ln(w, ")")
		ln(w)
	})

}
