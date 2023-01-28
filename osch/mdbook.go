package main

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type mdWriter struct {
	target string
	model  *YamlModel
	args   *args
}

func writeMarkdownBook(args *args) {
	model, err := ReadYamlModel(args)
	check(err, "could not read YAML model")
	target := cleanDir(args.home, "build", "docs")
	writer := &mdWriter{
		model:  model,
		target: target,
		args:   args}
	writer.writeBook()
}

func (w *mdWriter) writeBook() {

	w.file("book.toml", `[book]
language = "en"
multilingual = false
src = "src"
title = "openLCA schema"

[output.html]
mathjax-support = true
`)

	w.dir("src")
	w.file("src/SUMMARY.md", w.summary())

	// try to copy the schema README and CHANGES
	mds := []string{"intro.md", "CHANGES.md"}
	for _, md := range mds {
		mdPath := filepath.Join(w.args.home, md)
		if _, err := os.Stat(mdPath); err == nil {
			if text, err := os.ReadFile(mdPath); err == nil {
				w.file("src/"+md, string(text))
			} else {
				log.Println("WARNING: failed to copy", mdPath)
			}
		}
	}

	w.dir("src/classes")
	for _, t := range w.model.Types {
		if t.IsEnum() {
			continue
		}
		w.file("src/classes/"+t.Name()+".md", w.docClassOf(t.Class))
	}

	w.dir("src/enums")
	for _, t := range w.model.Types {
		if t.IsClass() {
			continue
		}
		w.file("src/enums/"+t.Name()+".md", w.docEnumOf(t.Enum))
	}

}

func (w *mdWriter) dir(path string) string {
	fullPath := filepath.Join(w.target, path)
	mkdir(fullPath)
	return fullPath
}

func (w *mdWriter) file(path, content string) {
	fullPath := filepath.Join(w.target, path)
	writeFile(fullPath, content)
}

func (w *mdWriter) summary() string {

	outerTypeOf := w.directCompositions()

	buff := NewBuffer()
	buff.Writeln("# Summary\n")
	buff.Writeln("[Introduction](./intro.md)")
	buff.Writeln("[Changes](./CHANGES.md)")

	addClassLinks := func(class *YamlClass) {
		if class == nil || outerTypeOf[class.Name] != "" {
			return
		}
		buff.Writeln(" - [" + class.Name + "](./classes/" + class.Name + ".md)")
		// write direct components
		w.model.EachClass(func(inner *YamlClass) {
			if outerTypeOf[inner.Name] == class.Name {
				buff.Writeln(
					"   - [" + inner.Name + "](./classes/" + inner.Name + ".md)")
			}
		})
	}

	// root entities and their direct components (the can only exist in
	// the root package)
	buff.Writeln("# Root entities\n")
	w.model.EachClass(func(class *YamlClass) {
		if w.model.IsRootEntity(class) {
			addClassLinks(class)
		}
	})

	// other shared components in the root package
	buff.Writeln("# Other components\n")
	w.model.EachClass(func(class *YamlClass) {
		if w.model.IsRootEntity(class) ||
			!w.model.IsRootPackage(w.model.PackageOfClass(class)) {
			return
		}
		addClassLinks(class)
	})

	// write the enumerations of the root package
	buff.Writeln("\n# Enumerations\n")
	for _, t := range w.model.Types {
		if t.IsClass() || !w.model.IsRootPackage(t.Package) {
			continue
		}
		buff.Writeln(" - [" + t.Name() + "](./enums/" + t.Name() + ".md)")
	}

	// write types in other packages
	for _, pack := range w.model.Packages() {
		if w.model.IsRootPackage(pack) {
			continue
		}
		buff.Writeln("\n# Package: ", pack)
		buff.Writeln("\n## Classes")
		for _, t := range w.model.Types {
			if !t.IsClass() || t.Package != pack {
				continue
			}
			addClassLinks(t.Class)
		}

		hasEnums := false
		for _, t := range w.model.Types {
			if !t.IsEnum() || t.Package != pack {
				continue
			}
			if !hasEnums {
				hasEnums = true
				buff.Writeln("\n## Enumerations")
			}
			buff.Writeln(" - [" + t.Name() + "](./enums/" + t.Name() + ".md)")
		}
	}

	return buff.String()
}

func (w *mdWriter) docClassOf(class *YamlClass) string {
	var buff bytes.Buffer
	buff.WriteString("# " + class.Name + "\n\n")
	buff.WriteString(class.Doc + "\n\n")

	buff.WriteString("## Properties\n\n")

	parents := make([]*YamlClass, 0)
	parent := w.model.ParentOf(class)
	for {
		if parent == nil {
			break
		}
		parents = append([]*YamlClass{parent}, parents...)
		parent = w.model.ParentOf(parent)
	}

	for _, p := range parents {
		for _, prop := range p.Props {
			buff.WriteString("### `" + prop.Name + "`\n\n")
			buff.WriteString("Inherited from [" + p.Name + "." + prop.Name +
				"](./" + p.Name + ".md#" + prop.Name + ")\n\n")
			buff.WriteString(w.docPropOf(prop))
		}
	}

	for _, prop := range class.Props {
		buff.WriteString("### `" + prop.Name + "`\n\n")
		if prop.Doc != "" {
			buff.WriteString(prop.Doc + "\n\n")
		}
		buff.WriteString(w.docPropOf(prop))
	}

	buff.WriteString("## Python class stub\n\n")
	buff.WriteString(`
The snippet below shows the names of the properties of the corresponding
Python class of the [olca-schema](https://pypi.org/project/olca-schema)
package. Note that this is not the full class definition but just shows
the names of the class and its properties.
`)
	buff.WriteString("\n\n```python\n\n")
	buff.WriteString("@dataclass\n")
	buff.WriteString("class " + class.Name + ":\n")
	for _, prop := range w.model.AllPropsOf(class) {
		if prop.Name == "@type" {
			continue
		}
		buff.WriteString("  " + prop.PyName() + ": " +
			prop.PropType().ToPython() + "\n")
	}
	buff.WriteString("\n```\n")

	example := w.getJsonExample(class)
	if example != "" {
		buff.WriteString("\n\n## JSON example\n\n")
		buff.WriteString("```json\n")
		buff.WriteString(example)
		buff.WriteString("\n```\n")
	}

	return buff.String()
}

func (w *mdWriter) docPropOf(prop *YamlProp) string {
	var buff bytes.Buffer
	buff.WriteString("* _Type:_ ")
	if prop.IsOptional {
		buff.WriteString("`optional` ")
	}
	buff.WriteString(w.docTypeOf(prop.Type) + "\n")

	buff.WriteString("* _Proto-Index:_ " + strconv.Itoa(prop.Index) + "\n")
	return buff.String()
}

func (w *mdWriter) docEnumOf(enum *YamlEnum) string {
	var buff bytes.Buffer
	buff.WriteString("# " + enum.Name + "\n\n")
	buff.WriteString(enum.Doc + "\n\n")

	buff.WriteString("## Items\n\n")

	for _, item := range enum.Items {
		buff.WriteString("### `" + item.Name + "`\n\n")
		if item.Doc != "" {
			buff.WriteString(item.Doc + "\n\n")
		}
		buff.WriteString("* _Proto-Index:_ " + strconv.Itoa(item.Index) + "\n")
	}

	return buff.String()
}

func (w *mdWriter) docTypeOf(yamlType string) string {

	if yamlType == "" {
		return "__ERROR! EMPTY__"
	}

	if strings.HasPrefix(yamlType, "List[") {
		unpacked := strings.TrimPrefix(strings.TrimSuffix(yamlType, "]"), "List[")
		return "`List` of " + w.docTypeOf(unpacked)
	}

	if strings.HasPrefix(yamlType, "Ref[") {
		unpacked := strings.TrimPrefix(strings.TrimSuffix(yamlType, "]"), "Ref[")
		return "[Ref](./Ref.md) of " + w.docTypeOf(unpacked)
	}

	if yamlType == "GeoJSON" {
		return "`GeoJSON` ([external doc](https://tools.ietf.org/html/rfc7946))"
	}

	if startsWithLower(yamlType) {
		return "`" + yamlType +
			"` ([external doc](http://www.w3.org/TR/xmlschema-2/#" + yamlType + "))"
	}

	t := w.model.TypeMap[yamlType]
	if t == nil {
		log.Println("WARNING: unknown type:", yamlType)
		return "`" + yamlType + "`"
	}
	if t.IsEnum() {
		return "[" + yamlType + "](../enums/" + yamlType + ".md)"
	} else {
		return "[" + yamlType + "](./" + yamlType + ".md)"
	}

}

// Returns a map `inner type -> outer type` of types that are only used in
// in a specific outer type (like Exchange in Processes).
func (w *mdWriter) directCompositions() map[string]string {
	m := make(map[string]string)
	for _, innerType := range w.model.Types {

		// filter out enums and root entitities
		if innerType.IsEnum() {
			continue
		}
		parent := w.model.ParentOf(innerType.Class)
		if parent == nil || parent.Name == "RootEntity" {
			continue
		}

		// checks if the given outer type has a property that
		// references the inner type
		hasInner := func(outerType *YamlClass) bool {
			for _, prop := range outerType.Props {
				propType := prop.Type
				if strings.HasPrefix(propType, "List[") {
					propType = strings.TrimPrefix(
						strings.TrimSuffix(propType, "]"), "List[")
				}
				if strings.HasPrefix(propType, "Ref[") {
					propType = strings.TrimPrefix(
						strings.TrimSuffix(propType, "]"), "Ref[")
				}
				if propType == innerType.Name() {
					return true
				}
			}
			return false
		}

		// search for a single outer class that references
		// the inner type
		candidate := ""
		for _, outer := range w.model.Types {
			if outer.IsEnum() {
				continue
			}
			if !hasInner(outer.Class) {
				continue
			}
			if candidate == "" {
				candidate = outer.Name()
			} else {
				candidate = ""
				break
			}
		}
		if candidate != "" {
			m[innerType.Name()] = candidate
		}

	}

	return m
}

func (w *mdWriter) getJsonExample(class *YamlClass) string {
	file := strings.ToLower(toSnakeCase(class.Name)) + ".json"
	path := filepath.Join(w.args.home, "examples", file)
	if _, err := os.Stat(path); err != nil {
		return ""
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	return string(data)
}
