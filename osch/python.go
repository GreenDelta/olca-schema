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
	ln(w, "class RefType(Enum):")
	w.model.EachClass(func(class *YamlClass) {
		if w.model.IsRefEntity(class) &&
			class.Name != "Ref" {
			lni(w, 1, class.Name, " = '", class.Name, "'")
		}
	})
	ln(w)
	w.writeEnumGetter("RefType")
	ln(w)

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
	ln(w, "RootEntity = Union[")
	w.model.EachClass(func(class *YamlClass) {
		if w.model.IsRootEntity(class) {
			lni(w, 1, class.Name, ",")
		}
	})
	ln(w, "]")
	ln(w)
	ln(w)
	ln(w, "RefEntity = Union[RootEntity, Unit, NwSet]")
}

func (w *pyw) writeHeader() {
	ln(w, "# DO NOT CHANGE THIS CODE AS THIS IS GENERATED AUTOMATICALLY")
	ln(w, `
# This module contains a Python API for reading and writing data sets in
# the JSON based openLCA data exchange format. For more information see
# http://greendelta.github.io/olca-schema
`)

	// imports
	ln(w, "import datetime")
	ln(w, "import json")
	ln(w, "import uuid")
	ln(w)
	ln(w, "from enum import Enum")
	ln(w, "from dataclasses import dataclass")
	ln(w, "from typing import Any, Dict, List, Optional, Union")
	ln(w)
	ln(w)
}

func (w *pyw) writeEnum(enum *YamlEnum) {
	name := enum.Name
	ln(w, "class ", name, "(Enum):")
	ln(w)
	for _, item := range enum.Items {
		lni(w, 1, item.Name, " = '", item.Name, "'")
	}
	ln(w)
	w.writeEnumGetter(name)
	ln(w)
}

func (w *pyw) writeEnumGetter(name string) {
	lni(w, 1, "@staticmethod")
	lni(w, 1, "def get(v: Union[str, '", name, "'],")
	lni(w, 3, "default: Optional['", name, "'] = None) -> Optional['", name, "']:")
	lni(w, 2, "for i in ", name, ":")
	lni(w, 3, "if i == v or i.value == v or i.name == v:")
	lni(w, 4, "return i")
	lni(w, 2, "return default")
	ln(w)
}

func (w *pyw) writeClass(class *YamlClass) {
	ln(w, "@dataclass")
	ln(w, "class ", class.Name, ":")
	ln(w)

	// properties
	for _, prop := range w.model.AllPropsOf(class) {
		if prop.Name == "@type" {
			continue
		}
		propType := YamlPropType(prop.Type)
		lni(w, 1, prop.PyName(), ": Optional[", propType.ToPython(), "] = None")
	}
	if class.Name == "Ref" {
		lni(w, 1, "ref_type: Optional[RefType] = None")
	}
	ln(w)

	// __post_init__
	if w.model.IsRootEntity(class) {
		fields := []string{"id", "version", "last_change"}
		inits := []string{"str(uuid.uuid4())", "'01.00.000'",
			"datetime.datetime.now(datetime.timezone.utc).isoformat()"}
		lni(w, 1, "def __post_init__(self):")
		for i, field := range fields {
			lni(w, 2, "if self.", field, " is None:")
			lni(w, 3, "self.", field, " = ", inits[i])
		}
		ln(w)
	}

	// to_dict
	lni(w, 1, "def to_dict(self) -> Dict[str, Any]:")
	if w.model.IsRootEntity(class) {
		lni(w, 2, "d: Dict[str, Any] = {'@type': '", class.Name, "'}")
	} else {
		lni(w, 2, "d: Dict[str, Any] = {}")
	}

	if class.Name == "Ref" {
		lni(w, 2, "if self.ref_type is not None:")
		lni(w, 3, "d['@type'] = self.ref_type.value")
	}
	for _, prop := range w.model.AllPropsOf(class) {
		if prop.Name == "@type" {
			continue
		}
		selfProp := "self." + prop.PyName()
		dictProp := "d['" + prop.Name + "']"
		propType := prop.PropType()
		lni(w, 2, "if ", selfProp, " is not None:")
		if propType.IsPrimitive() ||
			(propType.IsList() && propType.UnpackList().IsPrimitive()) ||
			propType == "GeoJSON" {
			lni(w, 3, dictProp, " = ", selfProp)
		} else if propType.IsEnumOf(w.model) {
			lni(w, 3, dictProp, " = ", selfProp, ".value")
		} else if propType.IsList() {
			lni(w, 3, dictProp, " = [e.to_dict() for e in ", selfProp, "]")
		} else {
			lni(w, 3, dictProp, " = ", selfProp, ".to_dict()")
		}
	}
	lni(w, 2, "return d")
	ln(w)

	// to_json
	if w.model.IsRootEntity(class) {
		lni(w, 1, "def to_json(self) -> str:")
		lni(w, 2, "return json.dumps(self.to_dict(), indent=2)")
		ln(w)
	}

	// to_ref
	if w.model.IsRefEntity(class) {
		lni(w, 1, "def to_ref(self) -> 'Ref':")
		lni(w, 2, "ref = Ref(id=self.id, name=self.name)")
		if w.model.IsRootEntity(class) {
			lni(w, 2, "ref.category = self.category")
		}
		if class.Name != "Ref" {
			lni(w, 2, "ref.ref_type = RefType.get('"+class.Name+"')")
		} else {
			lni(w, 2, "ref.ref_type = self.ref_type")
		}
		lni(w, 2, "return ref")
		ln(w)
	}

	w.writeFromDict(class)

	// from_json
	if w.model.IsRootEntity(class) {
		lni(w, 1, "@staticmethod")
		lni(w, 1, "def from_json(data: Union[str, bytes]) -> '", class.Name, "':")
		lni(w, 2, "return "+class.Name+".from_dict(json.loads(data))")
		ln(w)
	}

	ln(w)
}

func (w *pyw) writeFromDict(class *YamlClass) {

	classOf := func(propType YamlPropType) string {
		if propType.IsRef() {
			return "Ref"
		} else {
			return string(propType)
		}
	}

	lni(w, 1, "@staticmethod")
	lni(w, 1, "def from_dict(d: Dict[str, Any]) -> '", class.Name, "':")
	instance := strings.ToLower(toSnakeCase(class.Name))
	lni(w, 2, instance, " = ", class.Name, "()")

	// clear the fields that are set in __post_init__
	if w.model.IsRootEntity(class) {
		for _, field := range []string{"id", "last_change", "version"} {
			lni(w, 2, instance, ".", field, " = None")
		}
	}

	if class.Name == "Ref" {
		lni(w, 2, instance, ".ref_type = RefType.get(d.get('@type', ''))")
	}

	for _, prop := range w.model.AllPropsOf(class) {
		if prop.Name == "@type" {
			continue
		}
		lni(w, 2, "if (v := d.get('", prop.Name, "')) or v is not None:")
		propType := prop.PropType()
		modelProp := instance + "." + prop.PyName()

		if propType.IsPrimitive() ||
			(propType.IsList() && propType.UnpackList().IsPrimitive()) ||
			propType == "GeoJSON" {

			// direct assignments for primitives, list of primitives, or GeoJson
			// objects
			lni(w, 3, modelProp, " = v")

		} else if propType.IsEnumOf(w.model) {

			// enum getters
			lni(w, 3, modelProp, " = ", prop.Type, ".get(v)")

		} else if propType.IsList() {

			// list conversion for non-primitive types
			u := propType.UnpackList()
			lni(w, 3, modelProp, " = [", classOf(u), ".from_dict(e) for e in v]")

		} else {

			// from-dict calls for entity types
			lni(w, 3, modelProp, " = ", classOf(propType), ".from_dict(v)")
		}
	}
	lni(w, 2, "return "+instance)
	ln(w)
}
