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

	// read and check the model
	model, err := ReadYamlModel(args)
	check(err, "could not read YAML model")

	// prepare the package folder
	modDir := filepath.Join(args.home, "py", "olca_schema")
	mkdir(modDir)

	// write the packages
	for _, pack := range model.Packages() {

		var buffer bytes.Buffer
		writer := pyWriter{
			buff:  &buffer,
			model: model,
		}
		writer.writePackage(pack)

		var fileName string
		if model.IsRootPackage(pack) {
			fileName = "schema.py"
		} else {
			fileName = pack + ".py"
		}
		modFile := filepath.Join(modDir, fileName)
		writeFile(modFile, buffer.String())
	}
}

func (w *pyWriter) writePackage(pack string) {

	w.writeHeader(pack)

	// enums and classes
	w.model.EachEnum(func(enum *YamlEnum) {
		if w.model.PackageOfEnum(enum) == pack {
			w.writeEnum(enum)
		}
	})

	for _, class := range w.model.TopoSortClasses() {
		if w.model.PackageOfClass(class) != pack || w.model.IsAbstract(class) {
			continue
		}
		w.writeClass(class)
	}

	// write RootEntity type
	if w.model.IsRootPackage(pack) {
		w.writeln("RootEntity = Union[")
		w.model.EachClass(func(class *YamlClass) {
			if w.model.IsRootEntity(class) {
				w.wrind1ln(class.Name + ",")
			}
		})
		w.writeln("]")
	}
}

func (w *pyWriter) writeHeader(pack string) {
	w.writeln("# DO NOT CHANGE THIS CODE AS THIS IS GENERATED AUTOMATICALLY")
	w.writeln(`
# This module contains a Python API for reading and writing data sets in
# the JSON based openLCA data exchange format. For more information see
# http://greendelta.github.io/olca-schema
`)

	// imports
	if w.model.IsRootPackage(pack) {
		w.writeln("import datetime")
		w.writeln("import json")
		w.writeln("import uuid")
		w.writeln()
	}

	w.writeln("from enum import Enum")
	w.writeln("from dataclasses import dataclass")
	w.writeln("from typing import Any, Dict, List, Optional, Union")

	if !w.model.IsRootPackage(pack) {
		w.writeln()
		w.writeln("from .schema import *")
	}

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
		w.wrind1ln("model_type: str = ''")
	}
	w.writeln()

	// __post_init__
	if w.model.IsRootEntity(class) {
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
	if w.model.IsRootEntity(class) {
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
		w.wrind2ln("if " + selfProp + " is not None:")
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
	if w.model.IsRootEntity(class) {
		w.wrind1ln("def to_json(self) -> str:")
		w.wrind2ln("return json.dumps(self.to_dict(), indent=2)")
		w.writeln()
	}

	// to_ref
	if w.model.IsRootEntity(class) || class.Name == "Unit" {
		w.wrind1ln("def to_ref(self) -> 'Ref':")
		w.wrind2ln("ref = Ref(id=self.id, name=self.name)")
		if w.model.IsRootEntity(class) {
			w.wrind2ln("ref.category = self.category")
		}
		w.wrind2ln("ref.model_type = '" + class.Name + "'")
		w.wrind2ln("return ref")
		w.writeln()
	}

	w.writeFromDict(class)

	// from_json
	if w.model.IsRootEntity(class) {
		w.wrind1ln("@staticmethod")
		w.wrind1ln("def from_json(data: Union[str, bytes]) -> '" +
			class.Name + "':")
		w.wrind2ln("return " + class.Name + ".from_dict(json.loads(data))")
		w.writeln()
	}

	w.writeln()
}

func (w *pyWriter) writeFromDict(class *YamlClass) {

	classOf := func(propType YamlPropType) string {
		if propType.IsRef() {
			return "Ref"
		} else {
			return string(propType)
		}
	}

	w.wrind1ln("@staticmethod")
	w.wrind1ln("def from_dict(d: Dict[str, Any]) -> '" + class.Name + "':")
	instance := strings.ToLower(toSnakeCase(class.Name))
	w.wrind2ln(instance + " = " + class.Name + "()")
	if class.Name == "Ref" {
		w.wrind2ln(instance + ".model_type = d.get('@type', '')")
	}

	for _, prop := range w.model.AllPropsOf(class) {
		w.wrind2ln("if (v := d.get('" + prop.Name + "')) or v is not None:")
		propType := prop.PropType()
		modelProp := instance + "." + prop.PyName()

		if propType.IsPrimitive() ||
			(propType.IsList() && propType.UnpackList().IsPrimitive()) ||
			propType == "GeoJSON" {

			// direct assignments for primitives, list of primitives, or GeoJson
			// objects
			w.wrind3ln(modelProp + " = v")

		} else if propType.IsEnumOf(w.model) {

			// enum getters
			w.wrind3ln(modelProp + " = " + prop.Type + ".get(v)")

		} else if propType.IsList() {

			// list conversion for non-primitive types
			u := propType.UnpackList()
			w.wrind3ln(modelProp + " = [" + classOf(u) + ".from_dict(e) for e in v]")

		} else {

			// from-dict calls for entity types
			w.wrind3ln(modelProp + " = " + classOf(propType) + ".from_dict(v)")
		}
	}
	w.wrind2ln("return " + instance)
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
