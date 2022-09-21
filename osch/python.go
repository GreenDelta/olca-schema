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
			w.wrind1ln(class.Name + ",")
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
	w.writeln("class " + n + "(Enum):")
	w.writeln()

	// write the items
	for _, item := range enum.Items {
		w.wrind1ln(item.Name + " = '" + item.Name + "'")
	}
	w.writeln()

	// writer the get-method
	w.wrind1ln("def get(v: Union[str, '" + n + "'],")
	w.wrind3ln("default: Optional['" + n + "'] = None) -> '" + n + "':")
	w.wrind2ln("for i in " + n + ":")
	w.wrind3ln("if i == v or i.value == v or i.name == v:")
	w.wrind4ln("return i")
	w.wrind2ln("return default")
	w.writeln()
	w.writeln()
}

func (w *pyWriter) writeClass(class *YamlClass) {
	w.writeln("@dataclass")
	w.writeln("class " + class.Name + ":")
	w.writeln()

	// properties
	for _, prop := range w.model.AllPropsOf(class) {
		if prop.Name == "@type" {
			continue
		}
		propType := YamlPropType(prop.Type)
		w.wrind1ln(prop.PyName() +
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
		w.wrind1ln("def __post_init__(self):")
		for i, field := range fields {
			w.wrind2ln("if self." + field + " is None:")
			w.wrind3ln("self." + field + " = " + inits[i])
		}
		w.writeln()
	}

	// to_dict
	w.wrind1ln("def to_dict(self) -> Dict[str, Any]:")
	w.wrind2ln("d: Dict[str, Any] = {}")
	if w.model.IsRoot(class) {
		w.wrind2ln("d['@type'] = '" + class.Name + "'")
	}
	if class.Name == "Ref" {
		w.wrind2ln("d['@type'] = self.model_type")
	}
	for _, prop := range w.model.AllPropsOf(class) {
		if prop.Name == "@type" {
			continue
		}
		selfProp := "self." + prop.PyName()
		dictProp := "d['" + prop.Name + "']"
		propType := prop.PropType()
		w.writeln("        if " + selfProp + ":")
		if propType.IsPrimitive() ||
			(propType.IsList() && propType.UnpackList().IsPrimitive()) ||
			propType == "GeoJSON" {
			w.wrind3ln(dictProp + " = " + selfProp)
		} else if propType.IsEnumOf(w.model) {
			w.wrind3ln(dictProp + " = " + selfProp + ".value")
		} else if propType.IsList() {
			w.wrind3ln(dictProp + " = [e.to_dict() for e in " + selfProp + "]")
		} else {
			w.wrind3ln(dictProp + " = " + selfProp + ".to_dict()")
		}
	}
	w.wrind2ln("return d")
	w.writeln()

	// to_json
	if w.model.IsRoot(class) {
		w.wrind1ln("def to_json(self) -> str:")
		w.wrind2ln("return json.dumps(self.to_dict(), indent=2)")
		w.writeln()
	}

	// to_ref
	if w.model.IsRoot(class) || class.Name == "Unit" {
		w.wrind1ln("def to_ref(self) -> 'Ref':")
		w.wrind2ln("ref = Ref(id=self.id, name=self.name)")
		if w.model.IsRoot(class) {
			w.wrind2ln("ref.category = self.category")
		}
		w.wrind2ln("ref.model_type = '" + class.Name + "'")
		w.wrind2ln("return ref")
		w.writeln()
	}

	w.writeFromDict(class)

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

func (w *pyWriter) writeFromDict(class *YamlClass) {

	w.wrind1ln("@staticmethod")
	w.wrind1ln("def from_dict(d: Dict[str, Any]) -> '" + class.Name + "':")
	instance := strings.ToLower(toSnakeCase(class.Name))
	w.wrind2ln(instance + " = " + class.Name + "()")
	if class.Name == "Ref" {
		w.wrind2ln(instance + ".model_type = d.get('@type', '')")
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
}

func (w *pyWriter) wrind1ln(s string) {
	w.writeln("    ", s)
}

func (w *pyWriter) wrind2ln(s string) {
	w.writeln("        ", s)
}

func (w *pyWriter) wrind3ln(s string) {
	w.writeln("            ", s)
}

func (w *pyWriter) wrind4ln(s string) {
	w.writeln("                ", s)
}

func (w *pyWriter) writeln(xs ...string) {
	for _, x := range xs {
		w.buff.WriteString(x)
	}
	w.buff.WriteRune('\n')
}
