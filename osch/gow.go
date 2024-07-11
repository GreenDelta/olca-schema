package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
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
	w.structs()

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

func (w *gow) structs() {
	w.model.EachClass(func(c *YamlClass) {
		if w.model.IsAbstract(c) {
			return
		}
		ln(w, "type ", c.Name, " struct {")

		if c.Name == "Ref" {
			lni(w, 1, "Type string `json:\"@type,omitempty\"`")
		}

		for _, prop := range w.model.AllPropsOf(c) {
			if prop.Name == "@type" {
				continue
			}
			ptype := prop.PropType()
			fieldType := w.typeOf(ptype)
			if ptype.IsClassOf(w.model) || fieldType == "Ref" ||
				(prop.IsOptional && ptype.IsPrimitive() && fieldType != "string") {
				fieldType = "*" + fieldType
			}
			lni(w, 1, w.fieldOf(prop), " ", fieldType, " ", w.jsonTagOf(prop))
		}

		ln(w, "}")
		ln(w)
	})
}

func (w *gow) fieldOf(prop *YamlProp) string {
	if prop.Name == "@id" {
		return "ID"
	}
	name := prop.Name
	return strings.ToUpper(name[0:1]) + name[1:]
}

func (w *gow) typeOf(t YamlPropType) string {
	if t.IsList() {
		return "[]" + w.typeOf(t.UnpackList())
	}
	if t.IsRef() {
		return "Ref"
	}
	if t.IsClassOf(w.model) || t.IsEnumOf(w.model) {
		return string(t)
	}
	switch t {
	case "string", "date", "dateTime":
		return "string"
	case "double", "float":
		return "float64"
	case "int", "integer":
		return "int"
	case "bool", "boolean":
		return "bool"
	case "GeoJSON":
		return "map[string]any"
	default:
		return "?"
	}
}

func (w *gow) jsonTagOf(prop *YamlProp) string {
	tag := "`json:" + "\"" + prop.Name
	t := prop.PropType()
	if t.IsList() ||
		w.typeOf(t) == "string" ||
		t.IsEnumOf(w.model) ||
		t.IsClassOf(w.model) ||
		t.IsRef() ||
		prop.IsOptional {
		tag += ",omitempty"
	}
	return tag + "\"`"
}
