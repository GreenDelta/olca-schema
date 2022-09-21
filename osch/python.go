package main

import (
	"bytes"
	"path/filepath"
	"strings"
)

type pyWriter struct {
	buff  *bytes.Buffer
	model *YamlModel
}

// indentation levels
const pyInd1 = "    "
const pyInd2 = pyInd1 + pyInd1
const pyInd3 = pyInd2 + pyInd1
const pyInd4 = pyInd3 + pyInd1

func writePythonModule(args *args) {
	model, err := ReadYamlModel(args)
	check(err, "could not read YAML model")

	var buffer bytes.Buffer
	writer := pyWriter{
		buff:  &buffer,
		model: model,
	}
	writer.writeAll()

	modDir := filepath.Join(args.home, "py", "olca_schema")
	mkdir(modDir)
	modFile := filepath.Join(modDir, "schema.py")
	writeFile(modFile, buffer.String())
}

func (w *pyWriter) writeAll() {

	w.writeHeader()

	// enums and classes
	w.model.EachEnum(w.writeEnum)
	for _, class := range w.model.TopoSortClasses() {
		if w.model.IsAbstract(class) {
			continue
		}
		w.writeClass(class)
	}

	// write RootEntity type
	w.writeln("RootEntity = Union[")
	w.model.EachClass(func(class *YamlClass) {
		if w.model.IsRoot(class) {
			w.writeln(pyInd1 + class.Name + ",")
		}
	})
	w.writeln("]")
}

func (w *pyWriter) writeHeader() {
	w.writeln("# DO NOT CHANGE THIS CODE AS THIS IS GENERATED AUTOMATICALLY")
	w.writeln(`
# This module contains a Python API for reading and writing data sets in
# the JSON based openLCA data exchange format. For more information see
# http://greendelta.github.io/olca-schema
`)

	// imports
	w.writeln("import datetime")
	w.writeln("import json")
	w.writeln("import uuid")
	w.writeln()
	w.writeln("from enum import Enum")
	w.writeln("from dataclasses import dataclass")
	w.writeln("from typing import Any, Dict, List, Optional, Union")
	w.writeln()
	w.writeln()
}

func (w *pyWriter) writeEnum(enum *YamlEnum) {
	n := enum.Name
	w.writeln("class", n+"(Enum):")
	w.writeln()

	// write the items
	for _, item := range enum.Items {
		w.writeln(pyInd1 + item.Name + " = '" + item.Name + "'")
	}
	w.writeln()

	// writer the get-method
	w.writeln(pyInd1 + "def get(v: Union[str, '" + n + "'],\n" +
		pyInd3 + "default: Optional['" + n + "'] = None) -> '" + n + "':")
	w.writeln(pyInd2 + "for i in " + n + ":")
	w.writeln(pyInd3 + "if i == v or i.value == v or i.name == v:")
	w.writeln(pyInd4 + "return i")
	w.writeln(pyInd2 + "return default")
	w.writeln()
	w.writeln()
}

func (w *pyWriter) writeClass(class *YamlClass) {
	w.writeln("@dataclass")
	w.writeln("class", class.Name+":")
	w.writeln()

	// properties
	for _, prop := range w.model.AllPropsOf(class) {
		if prop.Name == "@type" {
			continue
		}
		propType := YamlPropType(prop.Type)
		w.writeln(pyInd1 + prop.PyName() +
			": Optional[" + propType.ToPython() + "] = None")
	}
	if class.Name == "Ref" {
		w.writeln("    model_type: str = ''")
	}
	w.writeln()

	// __post_init__
	if w.model.IsRoot(class) {
		fields := []string{"id", "version", "last_change"}
		inits := []string{"str(uuid.uuid4())", "'01.00.000'",
			"datetime.datetime.utcnow().isoformat() + 'Z'"}
		w.writeln(pyInd1 + "def __post_init__(self):")
		for i, field := range fields {
			w.writeln(pyInd2 + "if self." + field + " is None:")
			w.writeln(pyInd3 + "self." + field + " = " + inits[i])
		}
		w.writeln()
	}

	// to_dict
	w.writeln(pyInd1 + "def to_dict(self) -> Dict[str, Any]:")
	w.writeln(pyInd2 + "d: Dict[str, Any] = {}")
	if w.model.IsRoot(class) {
		w.writeln(pyInd2 + "d['@type'] = '" + class.Name + "'")
	}
	if class.Name == "Ref" {
		w.writeln(pyInd2 + "d['@type'] = self.model_type")
	}
	for _, prop := range w.model.AllPropsOf(class) {
		if prop.Name == "@type" {
			continue
		}
		selfProp := "self." + prop.PyName()
		dictProp := pyInd3 + "d['" + prop.Name + "']"
		propType := prop.PropType()
		w.writeln("        if " + selfProp + ":")
		if propType.IsPrimitive() ||
			(propType.IsList() && propType.UnpackList().IsPrimitive()) ||
			propType == "GeoJSON" {
			w.writeln(dictProp + " = " + selfProp)
		} else if propType.IsEnumOf(w.model) {
			w.writeln(dictProp + " = " + selfProp + ".value")
		} else if propType.IsList() {
			w.writeln(dictProp + " = [e.to_dict() for e in " + selfProp + "]")
		} else {
			w.writeln(dictProp + " = " + selfProp + ".to_dict()")
		}
	}
	w.writeln(pyInd2 + "return d")
	w.writeln()

	// to_json
	if w.model.IsRoot(class) {
		w.writeln(pyInd1 + "def to_json(self) -> str:")
		w.writeln(pyInd2 + "return json.dumps(self.to_dict(), indent=2)")
		w.writeln()
	}

	// to_ref
	if w.model.IsRoot(class) || class.Name == "Unit" {
		w.writeln(pyInd1 + "def to_ref(self) -> 'Ref':")
		w.writeln(pyInd2 + "ref = Ref(id=self.id, name=self.name)")
		if w.model.IsRoot(class) {
			w.writeln(pyInd2 + "ref.category = self.category")
		}
		w.writeln(pyInd2 + "ref.model_type = '" + class.Name + "'")
		w.writeln(pyInd2 + "return ref")
		w.writeln()
	}

	// from_dict
	w.writeln(pyInd1 + "@staticmethod")
	w.writeln(pyInd1 + "def from_dict(d: Dict[str, Any]) -> '" + class.Name + "':")
	instance := strings.ToLower(toSnakeCase(class.Name))
	w.writeln(pyInd2 + instance + " = " + class.Name + "()")
	if class.Name == "Ref" {
		w.writeln(pyInd2 + instance + ".model_type = d.get('@type', '')")
	}
	for _, prop := range w.model.AllPropsOf(class) {
		w.writeln("        if v := d.get('" + prop.Name + "'):")
		propType := prop.PropType()
		modelProp := "            " + instance + "." + prop.PyName()
		if propType.IsPrimitive() ||
			(propType.IsList() && propType.UnpackList().IsPrimitive()) ||
			propType == "GeoJSON" {
			w.writeln(modelProp + " = v")
		} else if propType.IsEnumOf(w.model) {
			w.writeln(modelProp + " = " + prop.Type + ".get(v)")
		} else if propType.IsList() {
			u := propType.UnpackList()
			var typeStr string
			if u.IsRef() {
				typeStr = "Ref"
			} else {
				typeStr = string(u)
			}
			w.writeln(modelProp + " = [" + typeStr + ".from_dict(e) for e in v]")
		} else {
			w.writeln(modelProp + " = " + string(propType) + ".from_dict(v)")
		}
	}
	w.writeln("        return " + instance)
	w.writeln()

	// from_json
	if w.model.IsRoot(class) {
		w.writeln("    @staticmethod")
		w.writeln("    def from_json(data: Union[str, bytes]) -> '" +
			class.Name + "':")
		w.writeln("        return " + class.Name + ".from_dict(json.loads(data))")
		w.writeln()
	}

	w.writeln()
}

func (w *pyWriter) writeln(args ...string) {
	w.write(args...)
	w.buff.WriteRune('\n')
}

func (w *pyWriter) write(args ...string) {
	for i, arg := range args {
		if i > 0 {
			w.buff.WriteRune(' ')
		}
		w.buff.WriteString(arg)
	}
}
