package main

import (
	"bytes"
	"path/filepath"
	"strings"
)

type pyw struct {
	buff  *bytes.Buffer
	model *YamlModel
}

func (w *pyw) buffer() *bytes.Buffer {
	return w.buff
}

func (w *pyw) indent() string {
	return "    "
}

func writePythonModule(args *args) {

	// read and check the model
	model, err := ReadYamlModel(args)
	check(err, "could not read YAML model")

	// prepare the package folder
	modDir := filepath.Join(args.home, "py", "olca_schema")
	mkdir(modDir)

	// write the schema
	var buffer bytes.Buffer
	writer := pyw{
		buff:  &buffer,
		model: model,
	}
	writer.writeSchema()
	fileName := "schema.py"
	modFile := filepath.Join(modDir, fileName)
	writeFile(modFile, buffer.String())
}

func (w *pyw) writeSchema() {

	w.writeHeader()

	// RefType
	w.writeln("class RefType(Enum):")
	w.model.EachClass(func(class *YamlClass) {
		if w.model.IsRefEntity(class) &&
			class.Name != "Ref" {
			w.wrind1ln(class.Name + " = '" + class.Name + "'")
		}
	})
	w.writeln()
	w.writeEnumGetter("RefType")
	w.writeln()

	// enums
	w.model.EachEnum(func(enum *YamlEnum) {
		w.writeEnum(enum)
	})

	// classes
	for _, class := range w.model.TopoSortClasses() {
		if w.model.IsAbstract(class) {
			continue
		}
		w.writeClass(class)
	}

	// RootEntity and RefEntity
	w.writeln("RootEntity = Union[")
	w.model.EachClass(func(class *YamlClass) {
		if w.model.IsRootEntity(class) {
			w.wrind1ln(class.Name + ",")
		}
	})
	w.writeln("]")
	w.writeln()
	w.writeln()
	w.writeln("RefEntity = Union[RootEntity, Unit, NwSet]")
}

func (w *pyw) writeHeader() {
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

func (w *pyw) writeEnum(enum *YamlEnum) {
	name := enum.Name
	w.writeln("class " + name + "(Enum):")
	w.writeln()
	for _, item := range enum.Items {
		w.wrind1ln(item.Name + " = '" + item.Name + "'")
	}
	w.writeln()
	w.writeEnumGetter(name)
	w.writeln()
}

func (w *pyw) writeEnumGetter(name string) {
	w.wrind1ln("@staticmethod")
	w.wrind1ln("def get(v: Union[str, '" + name + "'],")
	wlni(w, 3, "default: Optional['", name, "'] = None) -> Optional['", name, "']:")
	w.wrind2ln("for i in " + name + ":")
	wlni(w, 3, "if i == v or i.value == v or i.name == v:")
	wlni(w, 4, "return i")
	w.wrind2ln("return default")
	w.writeln()
}

func (w *pyw) writeClass(class *YamlClass) {
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
		w.wrind1ln("ref_type: Optional[RefType] = None")
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
			wlni(w, 3, "self.", field, " = ", inits[i])
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
		w.wrind2ln("if self.ref_type is not None:")
		wlni(w, 3, "d['@type'] = self.ref_type.value")
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
			wlni(w, 3, dictProp, " = ", selfProp)
		} else if propType.IsEnumOf(w.model) {
			wlni(w, 3, dictProp, " = ", selfProp, ".value")
		} else if propType.IsList() {
			wlni(w, 3, dictProp, " = [e.to_dict() for e in ", selfProp, "]")
		} else {
			wlni(w, 3, dictProp, " = ", selfProp, ".to_dict()")
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
	if w.model.IsRefEntity(class) {
		w.wrind1ln("def to_ref(self) -> 'Ref':")
		w.wrind2ln("ref = Ref(id=self.id, name=self.name)")
		if w.model.IsRootEntity(class) {
			w.wrind2ln("ref.category = self.category")
		}
		if class.Name != "Ref" {
			w.wrind2ln("ref.ref_type = RefType.get('" + class.Name + "')")
		} else {
			w.wrind2ln("ref.ref_type = self.ref_type")
		}
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

func (w *pyw) writeFromDict(class *YamlClass) {

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
		w.wrind2ln(instance + ".ref_type = RefType.get(d.get('@type', ''))")
	}

	for _, prop := range w.model.AllPropsOf(class) {
		if prop.Name == "@type" {
			continue
		}
		w.wrind2ln("if (v := d.get('" + prop.Name + "')) or v is not None:")
		propType := prop.PropType()
		modelProp := instance + "." + prop.PyName()

		if propType.IsPrimitive() ||
			(propType.IsList() && propType.UnpackList().IsPrimitive()) ||
			propType == "GeoJSON" {

			// direct assignments for primitives, list of primitives, or GeoJson
			// objects
			wlni(w, 3, modelProp, " = v")

		} else if propType.IsEnumOf(w.model) {

			// enum getters
			wlni(w, 3, modelProp, " = ", prop.Type, ".get(v)")

		} else if propType.IsList() {

			// list conversion for non-primitive types
			u := propType.UnpackList()
			wlni(w, 3, modelProp, " = [", classOf(u), ".from_dict(e) for e in v]")

		} else {

			// from-dict calls for entity types
			wlni(w, 3, modelProp, " = ", classOf(propType), ".from_dict(v)")
		}
	}
	w.wrind2ln("return " + instance)
	w.writeln()
}

func (w *pyw) wrind1ln(s string) {
	w.writeln("    ", s)
}

func (w *pyw) wrind2ln(s string) {
	w.writeln("        ", s)
}

func (w *pyw) writeln(xs ...string) {
	for _, x := range xs {
		w.buff.WriteString(x)
	}
	w.buff.WriteRune('\n')
}
