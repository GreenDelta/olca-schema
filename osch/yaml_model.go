package main

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v2"
)

type YamlType struct {
	Class *YamlClass `yaml:"class"`
	Enum  *YamlEnum  `yaml:"enum"`
}

func (yt *YamlType) IsClass() bool {
	return yt.Class != nil
}

func (yt *YamlType) IsEnum() bool {
	return yt.Enum != nil
}

func (yt *YamlType) String() string {
	if yt.Class != nil {
		return "ClassDef " + yt.Class.Name
	}
	if yt.Enum != nil {
		return "EnumDef " + yt.Enum.Name
	}
	return "Unknown TypeDef"
}

func (yt *YamlType) Name() string {
	if yt.Class != nil {
		return yt.Class.Name
	}
	if yt.Enum != nil {
		return yt.Enum.Name
	}
	return "Unknown"
}

type YamlClass struct {
	Name       string      `yaml:"name"`
	SuperClass string      `yaml:"superClass"`
	Doc        string      `yaml:"doc"`
	Props      []*YamlProp `yaml:"properties"`
}

type YamlEnum struct {
	Name  string          `yaml:"name"`
	Doc   string          `yaml:"doc"`
	Items []*YamlEnumItem `yaml:"items"`
}

type YamlEnumItem struct {
	Name  string `yaml:"name"`
	Doc   string `yaml:"doc"`
	Index int    `yaml:"index"`
}

type YamlModel struct {
	Types   []*YamlType
	TypeMap map[string]*YamlType
}

func (model *YamlModel) EachEnum(consumer func(enum *YamlEnum)) {
	for i := range model.Types {
		t := model.Types[i]
		if t.IsEnum() {
			consumer(t.Enum)
		}
	}
}

func (model *YamlModel) EachClass(consumer func(enum *YamlClass)) {
	for i := range model.Types {
		t := model.Types[i]
		if t.IsClass() {
			consumer(t.Class)
		}
	}
}

func (model *YamlModel) ParentOf(class *YamlClass) *YamlClass {
	if class == nil {
		return nil
	}
	parentName := class.SuperClass
	if parentName == "" {
		return nil
	}
	parent := model.TypeMap[parentName]
	if parent == nil || parent.IsEnum() {
		return nil
	}
	return parent.Class
}

// IsAbstract returns true when the given class is an abstract super class.
// Only the leafs of the class hierarchy are non-abstract classes.
func (model *YamlModel) IsAbstract(class *YamlClass) bool {
	for _, t := range model.Types {
		if t.IsClass() {
			if t.Class.SuperClass == class.Name {
				return true
			}
		}
	}
	return false
}

func (model *YamlModel) IsEmpty() bool {
	return len(model.Types) == 0
}

func ReadYamlModel(dir string) (*YamlModel, error) {

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	types := make([]*YamlType, 0)
	for _, file := range files {
		name := file.Name()
		if !strings.HasSuffix(name, ".yaml") {
			continue
		}

		log.Println("Parse YAML file", name)
		path := filepath.Join(dir, name)
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, err
		}
		typeDef := &YamlType{}
		if err := yaml.Unmarshal(data, typeDef); err != nil {
			return nil, err
		}

		types = append(types, typeDef)
	}
	log.Println("Collected", len(types), "YAML types")

	typeMap := make(map[string]*YamlType)
	for i := range types {
		typeDef := types[i]
		typeMap[typeDef.Name()] = typeDef
	}

	model := YamlModel{Types: types, TypeMap: typeMap}

	return &model, nil
}

// AllPropsOf returns all properties of the given class including the properties
// of all its parent classes.
func (model *YamlModel) AllPropsOf(class *YamlClass) []*YamlProp {
	props := make([]*YamlProp, 0, len(class.Props)+1)
	c := class
	for {
		if c == nil {
			break
		}
		props = append(props, c.Props...)
		c = model.ParentOf(c)
	}
	sort.Sort(YamlPropsByName(props))
	return props
}

// IsRoot returns true if the given class is a root entity. This is the case
// when `RootEntity` is a parent class of the given class.
func (model *YamlModel) IsRoot(class *YamlClass) bool {
	c := class
	for {
		parent := model.ParentOf(c)
		if parent == nil {
			return false
		}
		if parent.Name == "RootEntity" {
			return true
		}
		c = parent
	}
}
