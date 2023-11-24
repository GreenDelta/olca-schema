package main

import (
	"bytes"
	"fmt"
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
	writer.writeSumTypes()
	writer.writeEnums()
	writer.writeClasses()

	var outFile string
	if args.output != "" {
		mkdir(filepath.Dir(outFile))
		outFile = args.output
	} else {
		outDir := filepath.Join(args.home, "build")
		mkdir(outDir)
		outFile = filepath.Join(outDir, "schema.ts")
	}

	os.WriteFile(outFile, buffer.Bytes(), os.ModePerm)
}

func (w *tsWriter) writeUtils() {
	w.writeln(`// this file was generated automatically; do not change it but help to make
// the code generator better; see:
// https://github.com/GreenDelta/olca-schema/tree/master/osch

// #region: utils
type Dict = Record<string, unknown>;

interface Dictable {
  toDict: () => Dict;
}

function ifPresent<T>(val: T | undefined, consumer: (val: T) => void) {
  if (val !== null && val !== undefined) {
    consumer(val);
  }
}

function dictAll(list: Array<Dictable> | null): Array<Dict> {
  return list ? list.map((e) => e.toDict()) : [];
}
// #endregion
`)
}

func (w *tsWriter) writeRefType() {
	w.writeln("export enum RefType {")
	w.model.EachClass(func(class *YamlClass) {
		if w.model.IsRefEntity(class) &&
			class.Name != "Ref" &&
			!w.model.IsAbstract(class) {
			w.writeln("  ", class.Name+" = \""+class.Name+"\",")
		}
	})
	w.writeln("}")
	w.writeln()
}

func (w *tsWriter) writeSumTypes() {
	text := "export type RootEntity = "
	w.model.EachClass(func(class *YamlClass) {
		if !w.model.IsAbstract(class) && w.model.IsRootEntity(class) {
			text += "\n  | " + class.Name
		}
	})
	text += ";\n"
	w.writeln(text)

	text = "export type RefEntity =\n  | RootEntity"
	w.model.EachClass(func(class *YamlClass) {
		if !w.model.IsAbstract(class) &&
			w.model.IsRefEntity(class) &&
			!w.model.IsRootEntity(class) &&
			class.Name != "Ref" {
			text += "\n  | " + class.Name
		}
	})
	text += ";\n"
	w.writeln(text)
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
		w.writeFromDict(class)

		// toJson & fromJson
		w.writeln(fmt.Sprintf(`
  static fromJson(json: string | Dict): %[1]s | null {
    return typeof json === "string"
      ? %[1]s.fromDict(JSON.parse(json) as Dict)
      : %[1]s.fromDict(json);
  }

  toJson(): string {
    return JSON.stringify(this.toDict(), null, "  ");
  }`, class.Name))

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
		w.writeln("  ", propName, "?: ", w.typeOf(propType), " | null;")
	}
	if class.Name == "Ref" {
		w.writeln("  refType?: RefType | null;")
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
	if w.model.IsRefEntity(class) && class.Name != "Ref" {
		w.writeln("    d[\"@type\"] = \"", class.Name, "\";")
	}
	if class.Name == "Ref" {
		w.writeln("    ifPresent(this.refType, (v) => d[\"@type\"] = v);")
	}
	for _, prop := range w.model.AllPropsOf(class) {
		if prop.Name == "@type" {
			continue
		}
		if prop.Name == "@id" {
			w.writeln("    ifPresent(this.id, (v) => d[\"@id\"] = v);")
			continue
		}

		t := prop.PropType()
		conv := "v"
		if t.IsList() && !t.UnpackList().IsPrimitive() {
			conv = "dictAll(v)"
		} else if !t.IsList() &&
			!t.IsPrimitive() &&
			!t.IsEnumOf(w.model) &&
			prop.Type != "GeoJSON" {
			conv = "v?.toDict()"
		}
		w.writeln("    ifPresent(this.", prop.Name,
			", (v) => d.", prop.Name, " = ", conv, ");")
	}
	w.writeln("    return d;")
	w.writeln("  }")
}

func (w *tsWriter) writeFromDict(class *YamlClass) {
	w.writeln("  static fromDict(d: Dict): ", class.Name, " | null {")
	w.writeln("    if (!d) return null;")
	w.writeln("    const e = new ", class.Name, "();")

	// set type attribute for instances of Ref
	if class.Name == "Ref" {
		w.writeln("    ifPresent(d[\"@type\"], (v) => e.refType = v as RefType);")
	}
	
	for _, prop := range w.model.AllPropsOf(class) {
		if prop.Name == "@type" {
			continue
		}
		if prop.Name == "@id" {
			w.writeln("    e.id = d[\"@id\"] as string;")
			continue
		}
		t := prop.PropType()
		if !t.IsList() {
			if t.IsPrimitive() || t == "GeoJSON" || t.IsEnumOf(w.model) {
				w.writeln("    e.", prop.Name, " = d.",
					prop.Name, " as ", w.typeOf(t), ";")
			} else {
				w.writeln("    e.", prop.Name, " = ",
					w.typeOf(t), ".fromDict(d.", prop.Name, " as Dict);")
			}
		} else {
			inner := t.UnpackList()
			if inner.IsPrimitive() {
				w.writeln("    e.", prop.Name, " = d.",
					prop.Name, " as ", w.typeOf(inner), "[];")
			} else {
				text := "    e." + prop.Name + " = d." + prop.Name + "\n"
				text += "      ? (d." + prop.Name + " as Dict[]).map(" +
					w.typeOf(inner) + ".fromDict) as " + w.typeOf(inner) + "[]\n" +
					"      : null;"
				w.writeln(text)
			}
		}
	}
	w.writeln("    return e;")
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
		return "Record<string, unknown>"
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
