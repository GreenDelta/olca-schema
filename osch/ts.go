package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
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

	out, err := args.outputFileOrDefault("/build/schema.ts")
	check(err, "failed to create file "+out)
	os.WriteFile(out, buffer.Bytes(), os.ModePerm)
}

func (w *tsw) writeUtils() {
	ln(w, `// this file was generated automatically; do not change it but help to make
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
	ln(w, "export enum RefType {")
	w.model.EachClass(func(class *YamlClass) {
		if w.model.IsRefEntity(class) &&
			class.Name != "Ref" &&
			!w.model.IsAbstract(class) {
			lni(w, 1, class.Name+" = \""+class.Name+"\",")
		}
	})
	ln(w, "}")
	ln(w)
}

func (w *tsw) writeSumTypes() {
	text := "export type RootEntity = "
	w.model.EachClass(func(class *YamlClass) {
		if !w.model.IsAbstract(class) && w.model.IsRootEntity(class) {
			text += "\n  | " + class.Name
		}
	})
	text += ";\n"
	ln(w, text)

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
	ln(w, text)
}

func (w *tsw) writeEnums() {
	w.model.EachEnum(func(e *YamlEnum) {
		ln(w, "export enum ", e.Name, " {")
		for _, item := range e.Items {
			lni(w, 1, item.Name, " = \""+item.Name+"\",")
		}
		ln(w, "}")
		ln(w)
	})
}

func (w *tsw) writeClasses() {
	w.model.EachClass(func(class *YamlClass) {
		if w.model.IsAbstract(class) {
			return
		}

		// write a companion interfase
		ln(w, "interface I", class.Name, " {")
		w.writeProps(class)
		ln(w, "}")
		ln(w)

		// write the class
		ln(w, "export class ", class.Name, " {")
		w.writeProps(class)
		ln(w)
		w.writeOfFactory(class)
		if w.model.IsRefEntity(class) && class.Name != "Ref" {
			ln(w)
			w.writeToRef(class)
		}
		ln(w)
		w.writeToDict(class)
		ln(w)
		w.writeFromDict(class)

		// toJson & fromJson
		ln(w, fmt.Sprintf(`
  static fromJson(json: string | Dict): %[1]s | null {
    return typeof json === "string"
      ? %[1]s.fromDict(JSON.parse(json) as Dict)
      : %[1]s.fromDict(json);
  }

  toJson(): string {
    return JSON.stringify(this.toDict(), null, "  ");
  }`, class.Name))

		ln(w, "}")
		ln(w)
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
		lni(w, 1, propName, "?: ", w.typeOf(propType), " | null;")
	}
	if class.Name == "Ref" {
		lni(w, 1, "refType?: RefType | null;")
	}
}

func (w *tsw) writeOfFactory(class *YamlClass) {
	lni(w, 1, "static of(i: I", class.Name, "): ", class.Name, " {")
	lni(w, 2, "const e = new ", class.Name, "();")
	for _, prop := range w.model.AllPropsOf(class) {
		if prop.Name == "@type" {
			continue
		}
		propName := prop.Name
		if prop.Name == "@id" {
			propName = "id"
		}
		lni(w, 2, "e.", propName, " = i.", propName, ";")
	}
	if class.Name == "Ref" {
		lni(w, 2, "e.refType = i.refType;")
	}
	lni(w, 2, "return e;")
	lni(w, 1, "}")
}

func (w *tsw) writeToRef(class *YamlClass) {
	lni(w, 1, "toRef(): Ref {")
	lni(w, 2, "return Ref.of({")
	lni(w, 3, "refType: RefType.", class.Name, ",")
	lni(w, 3, "id: this.id,")
	lni(w, 3, "name: this.name,")
	if w.model.IsRootEntity(class) {
		lni(w, 3, "category: this.category,")
	}
	lni(w, 2, "});")
	lni(w, 1, "}")
}

func (w *tsw) writeToDict(class *YamlClass) {
	lni(w, 1, "toDict(): Dict {")
	lni(w, 2, "const d: Dict = {};")
	if w.model.IsRefEntity(class) && class.Name != "Ref" {
		lni(w, 2, "d[\"@type\"] = \"", class.Name, "\";")
	}
	if class.Name == "Ref" {
		lni(w, 2, "ifPresent(this.refType, (v) => d[\"@type\"] = v);")
	}
	for _, prop := range w.model.AllPropsOf(class) {
		if prop.Name == "@type" {
			continue
		}
		if prop.Name == "@id" {
			lni(w, 2, "ifPresent(this.id, (v) => d[\"@id\"] = v);")
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
		lni(w, 2, "ifPresent(this.", prop.Name,
			", (v) => d.", prop.Name, " = ", conv, ");")
	}
	lni(w, 2, "return d;")
	lni(w, 1, "}")
}

func (w *tsw) writeFromDict(class *YamlClass) {
	lni(w, 1, "static fromDict(d: Dict): ", class.Name, " | null {")
	lni(w, 2, "if (!d) return null;")
	lni(w, 2, "const e = new ", class.Name, "();")

	// set type attribute for instances of Ref
	if class.Name == "Ref" {
		lni(w, 2, "ifPresent(d[\"@type\"], (v) => e.refType = v as RefType);")
	}

	for _, prop := range w.model.AllPropsOf(class) {
		if prop.Name == "@type" {
			continue
		}
		if prop.Name == "@id" {
			lni(w, 2, "e.id = d[\"@id\"] as string;")
			continue
		}
		t := prop.PropType()
		if !t.IsList() {
			if t.IsPrimitive() || t == "GeoJSON" || t.IsEnumOf(w.model) {
				lni(w, 2, "e.", prop.Name, " = d.",
					prop.Name, " as ", w.typeOf(t), ";")
			} else {
				lni(w, 2, "e.", prop.Name, " = ",
					w.typeOf(t), ".fromDict(d.", prop.Name, " as Dict);")
			}
		} else {
			inner := t.UnpackList()
			if inner.IsPrimitive() {
				lni(w, 2, "e.", prop.Name, " = d.",
					prop.Name, " as ", w.typeOf(inner), "[];")
			} else {
				text := "    e." + prop.Name + " = d." + prop.Name + "\n"
				text += "      ? (d." + prop.Name + " as Dict[]).map(" +
					w.typeOf(inner) + ".fromDict) as " + w.typeOf(inner) + "[]\n" +
					"      : null;"
				ln(w, text)
			}
		}
	}
	lni(w, 2, "return e;")
	lni(w, 1, "}")
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
