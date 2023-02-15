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
	writer.writeUtils()
	writer.writeRefType()
	writer.writeEnums()
	writer.writeClasses()

	outDir := filepath.Join(args.home, "build")
	mkdir(outDir)
	file := filepath.Join(outDir, "schema.ts")
	os.WriteFile(file, buffer.Bytes(), os.ModePerm)
}

func (w *tsWriter) writeUtils() {
	w.writeln(`
// #region: utils
type Dict = {[field: string]: any};

interface Dictable {
  toDict: () => Dict,
}

function ifPresent<T>(val: T | undefined, consumer: (val: T) => void) {
  if (val !== null && val !== undefined) {
    consumer(val);
  }
}

function dictAll(list: Array<Dictable>): Array<Dict> {
  return list.map(e => e.toDict());
}
// #endregion
`)
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
		w.writeOfFactory(class)
		if w.model.IsRefEntity(class) && class.Name != "Ref" {
			w.writeln()
			w.writeToRef(class)
		}
		w.writeln()
		w.writeToDict(class)

		w.writeln()
		w.writeln("  toJson(): string {")
		w.writeln("    return JSON.stringify(this.toDict(), null, \"  \");")
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
			propName = "id"
		}
		propType := YamlPropType(prop.Type)
		w.writeln("  ", propName, "?: ", w.typeOf(propType), ";")
	}
	if class.Name == "Ref" {
		w.writeln("  refType?: RefType;")
	}
}

func (w *tsWriter) writeOfFactory(class *YamlClass) {
	w.writeln("  static of(i: I", class.Name, "): ", class.Name, " {")
	w.writeln("    const e = new ", class.Name, "();")
	for _, prop := range w.model.AllPropsOf(class) {
		if prop.Name == "@type" {
			continue
		}
		propName := prop.Name
		if prop.Name == "@id" {
			propName = "id"
		}
		w.writeln("    e.", propName, " = i.", propName, ";")
	}
	if class.Name == "Ref" {
		w.writeln("    e.refType = i.refType;")
	}
	w.writeln("    return e;")
	w.writeln("  }")
}

func (w *tsWriter) writeToRef(class *YamlClass) {
	w.writeln("  toRef(): Ref {")
	w.writeln("    return Ref.of({")
	w.writeln("      refType: RefType.", class.Name, ",")
	w.writeln("      id: this.id,")
	w.writeln("      name: this.name,")
	if w.model.IsRootEntity(class) {
		w.writeln("      category: this.category,")
	}
	w.writeln("    });")
	w.writeln("  }")
}

func (w *tsWriter) writeToDict(class *YamlClass) {
	w.writeln("  toDict(): Dict {")
	w.writeln("    const d: Dict = {};")
	for _, prop := range w.model.AllPropsOf(class) {
		if prop.Name == "@type" {
			continue
		}
		propName := prop.Name
		if prop.Name == "@id" {
			propName = "id"
		}
		t := prop.PropType()
		conv := "v"
		if t.IsList() && !t.UnpackList().IsPrimitive() {
			conv = "dictAll(v)"
		}
		w.writeln("    ifPresent(this.", propName,
			", v => d.", propName, " = ", conv, ");")
	}
	w.writeln("    return d;")
	w.writeln("  }")
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
