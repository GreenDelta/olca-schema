package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type tsw struct {
	buff  *bytes.Buffer
	model *YamlModel
}

func (w *tsw) buffer() *bytes.Buffer {
	return w.buff
}

func (w *tsw) indent() string {
	return "  "
}

func writeTypeScriptModule(args *args) {

	model, err := ReadYamlModel(args)
	check(err, "could not read YAML model")

	var buffer bytes.Buffer
	w := tsw{
		buff:  &buffer,
		model: model,
	}
	w.writeUtils()
	w.writeRefType()
	w.writeSumTypes()
	w.writeEnums()
	w.writeClasses()

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

func (w *tsw) writeUtils() {
	wln(w, `// this file was generated automatically; do not change it but help to make
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

func (w *tsw) writeRefType() {
	wln(w, "export enum RefType {")
	w.model.EachClass(func(class *YamlClass) {
		if w.model.IsRefEntity(class) &&
			class.Name != "Ref" &&
			!w.model.IsAbstract(class) {
			wlni(w, 1, class.Name+" = \""+class.Name+"\",")
		}
	})
	wln(w, "}")
	wln(w)
}

func (w *tsw) writeSumTypes() {
	text := "export type RootEntity = "
	w.model.EachClass(func(class *YamlClass) {
		if !w.model.IsAbstract(class) && w.model.IsRootEntity(class) {
			text += "\n  | " + class.Name
		}
	})
	text += ";\n"
	wln(w, text)

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
	wln(w, text)
}

func (w *tsw) writeEnums() {
	w.model.EachEnum(func(e *YamlEnum) {
		wln(w, "export enum ", e.Name, " {")
		for _, item := range e.Items {
			wlni(w, 1, item.Name, " = \""+item.Name+"\",")
		}
		wln(w, "}")
		wln(w)
	})
}

func (w *tsw) writeClasses() {
	w.model.EachClass(func(class *YamlClass) {
		if w.model.IsAbstract(class) {
			return
		}

		// write a companion interfase
		wln(w, "interface I", class.Name, " {")
		w.writeProps(class)
		wln(w, "}")
		wln(w)

		// write the class
		wln(w, "export class ", class.Name, " {")
		w.writeProps(class)
		wln(w)
		w.writeOfFactory(class)
		if w.model.IsRefEntity(class) && class.Name != "Ref" {
			wln(w)
			w.writeToRef(class)
		}
		wln(w)
		w.writeToDict(class)
		wln(w)
		w.writeFromDict(class)

		// toJson & fromJson
		wln(w, fmt.Sprintf(`
  static fromJson(json: string | Dict): %[1]s | null {
    return typeof json === "string"
      ? %[1]s.fromDict(JSON.parse(json) as Dict)
      : %[1]s.fromDict(json);
  }

  toJson(): string {
    return JSON.stringify(this.toDict(), null, "  ");
  }`, class.Name))

		wln(w, "}")
		wln(w)
	})
}

func (w *tsw) writeProps(class *YamlClass) {
	for _, prop := range w.model.AllPropsOf(class) {
		if prop.Name == "@type" {
			continue
		}
		propName := prop.Name
		if prop.Name == "@id" {
			propName = "id"
		}
		propType := YamlPropType(prop.Type)
		wlni(w, 1, propName, "?: ", w.typeOf(propType), " | null;")
	}
	if class.Name == "Ref" {
		wlni(w, 1, "refType?: RefType | null;")
	}
}

func (w *tsw) writeOfFactory(class *YamlClass) {
	wlni(w, 1, "static of(i: I", class.Name, "): ", class.Name, " {")
	wlni(w, 2, "const e = new ", class.Name, "();")
	for _, prop := range w.model.AllPropsOf(class) {
		if prop.Name == "@type" {
			continue
		}
		propName := prop.Name
		if prop.Name == "@id" {
			propName = "id"
		}
		wlni(w, 2, "e.", propName, " = i.", propName, ";")
	}
	if class.Name == "Ref" {
		wlni(w, 2, "e.refType = i.refType;")
	}
	wlni(w, 2, "return e;")
	wlni(w, 1, "}")
}

func (w *tsw) writeToRef(class *YamlClass) {
	wlni(w, 1, "toRef(): Ref {")
	wlni(w, 2, "return Ref.of({")
	wlni(w, 3, "refType: RefType.", class.Name, ",")
	wlni(w, 3, "id: this.id,")
	wlni(w, 3, "name: this.name,")
	if w.model.IsRootEntity(class) {
		wlni(w, 3, "category: this.category,")
	}
	wlni(w, 2, "});")
	wlni(w, 1, "}")
}

func (w *tsw) writeToDict(class *YamlClass) {
	wlni(w, 1, "toDict(): Dict {")
	wlni(w, 2, "const d: Dict = {};")
	if w.model.IsRefEntity(class) && class.Name != "Ref" {
		wlni(w, 2, "d[\"@type\"] = \"", class.Name, "\";")
	}
	if class.Name == "Ref" {
		wlni(w, 2, "ifPresent(this.refType, (v) => d[\"@type\"] = v);")
	}
	for _, prop := range w.model.AllPropsOf(class) {
		if prop.Name == "@type" {
			continue
		}
		if prop.Name == "@id" {
			wlni(w, 2, "ifPresent(this.id, (v) => d[\"@id\"] = v);")
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
		wlni(w, 2, "ifPresent(this.", prop.Name,
			", (v) => d.", prop.Name, " = ", conv, ");")
	}
	wlni(w, 2, "return d;")
	wlni(w, 1, "}")
}

func (w *tsw) writeFromDict(class *YamlClass) {
	wlni(w, 1, "static fromDict(d: Dict): ", class.Name, " | null {")
	wlni(w, 2, "if (!d) return null;")
	wlni(w, 2, "const e = new ", class.Name, "();")

	// set type attribute for instances of Ref
	if class.Name == "Ref" {
		wlni(w, 2, "ifPresent(d[\"@type\"], (v) => e.refType = v as RefType);")
	}

	for _, prop := range w.model.AllPropsOf(class) {
		if prop.Name == "@type" {
			continue
		}
		if prop.Name == "@id" {
			wlni(w, 2, "e.id = d[\"@id\"] as string;")
			continue
		}
		t := prop.PropType()
		if !t.IsList() {
			if t.IsPrimitive() || t == "GeoJSON" || t.IsEnumOf(w.model) {
				wlni(w, 2, "e.", prop.Name, " = d.",
					prop.Name, " as ", w.typeOf(t), ";")
			} else {
				wlni(w, 2, "e.", prop.Name, " = ",
					w.typeOf(t), ".fromDict(d.", prop.Name, " as Dict);")
			}
		} else {
			inner := t.UnpackList()
			if inner.IsPrimitive() {
				wlni(w, 2, "e.", prop.Name, " = d.",
					prop.Name, " as ", w.typeOf(inner), "[];")
			} else {
				text := "    e." + prop.Name + " = d." + prop.Name + "\n"
				text += "      ? (d." + prop.Name + " as Dict[]).map(" +
					w.typeOf(inner) + ".fromDict) as " + w.typeOf(inner) + "[]\n" +
					"      : null;"
				wln(w, text)
			}
		}
	}
	wlni(w, 2, "return e;")
	wlni(w, 1, "}")
}

func (w *tsw) typeOf(t YamlPropType) string {
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
