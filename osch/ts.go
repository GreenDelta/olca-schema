package main

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
)

type tsWriter struct {
	buff  *bytes.Buffer
	model *YamlModel
}

func writeTypeScriptModule(args *args) {

	model, err := ReadYamlModel(args)
	check(err, "could not read YAML model")

	var buffer bytes.Buffer
	writer := tsWriter{
		buff:  &buffer,
		model: model,
	}
	writer.writeRefType()
	writer.writeEnums()
	writer.writeClasses()

	outDir := filepath.Join(args.home, "build")
	mkdir(outDir)
	file := filepath.Join(outDir, "schema.ts")
	os.WriteFile(file, buffer.Bytes(), os.ModePerm)
}

func (w *tsWriter) writeRefType() {
	w.writeln("export enum RefType {")
	w.model.EachClass(func(class *YamlClass) {
		if w.model.IsRefEntity(class) &&
			class.Name != "Ref" {
			w.writeln("  ", class.Name+" = \""+class.Name+"\",")
		}
	})
	w.writeln("}")
	w.writeln()
}

func (w *tsWriter) writeEnums() {
	w.model.EachEnum(func(e *YamlEnum) {
		w.writeln("export enum ", e.Name, " {")
		for _, item := range e.Items {
			w.writeln("  ", item.Name, " = \""+item.Name+"\",")
		}
		w.writeln("}")
		w.writeln()
	})
}

func (w *tsWriter) writeClasses() {
	w.model.EachClass(func(class *YamlClass) {
		if w.model.IsAbstract(class) {
			return
		}

		// write a companion interfase
		w.writeln("interface I", class.Name, " {")
		w.writeProps(class)
		w.writeln("}")
		w.writeln()

		// write the class
		w.writeln("export class ", class.Name, " {")
		w.writeProps(class)
		w.writeln()

		// write the `of` factory method
		w.writeln("  static of(i: I", class.Name, "): ", class.Name, " {")
		w.writeln("    const e = new ", class.Name, "();")
		for _, prop := range w.model.AllPropsOf(class) {
			if prop.Name == "@type" {
				continue
			}
			propName := prop.Name
			if prop.Name == "@id" {
				propName = "uid"
			}
			w.writeln("    e.", propName, " = i.", propName, ";")
		}
		if class.Name == "Ref" {
			w.writeln("    e.refType = i.refType;")
		}
		w.writeln("    return e;")
		w.writeln("  }")

		w.writeln("}")
		w.writeln()
	})
}

func (w *tsWriter) writeProps(class *YamlClass) {
	for _, prop := range w.model.AllPropsOf(class) {
		if prop.Name == "@type" {
			continue
		}
		propName := prop.Name
		if prop.Name == "@id" {
			propName = "uid"
		}
		propType := YamlPropType(prop.Type)
		w.writeln("  ", propName, "?: ", w.typeOf(propType), ";")
	}
	if class.Name == "Ref" {
		w.writeln("  refType?: RefType;")
	}
}

func (w *tsWriter) typeOf(t YamlPropType) string {
	if t.IsList() {
		param := t.UnpackList()
		return "Array<" + w.typeOf(param) + ">"
	}
	if t.IsRef() {
		return "Ref"
	}
	switch t {
	case "string", "date", "dateTime":
		return "string"
	case "double", "float", "int", "integer":
		return "number"
	case "bool", "boolean":
		return "boolean"
	case "GeoJSON":
		return "any"
	default:
		if startsWithLower(string(t)) {
			log.Println("WARNING: unknown primitive type:", t)
			return "any"
		} else {
			return string(t)
		}
	}
}

func (w *tsWriter) writeln(xs ...string) {
	for _, x := range xs {
		w.buff.WriteString(x)
	}
	w.buff.WriteRune('\n')
}
