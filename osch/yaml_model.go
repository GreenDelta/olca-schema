package main

import (
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v2"
)

type YamlType struct {
	Class   *YamlClass `yaml:"class"`
	Enum    *YamlEnum  `yaml:"enum"`
	Package string
}

func (yt *YamlType) IsClass() bool {
	return yt.Class != nil
}

func (yt *YamlType) IsEnum() bool {
	return yt.Enum != nil
}

func (yt *YamlType) IsCoreType() bool {
	return yt.Package == ""
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

// Packages returns the used package names of model. Note that the name of the
// "root" or "core" package is just the empty string, and will be also returned
// in that list.
func (model *YamlModel) Packages() []string {
	packs := make([]string, 0)
	handled := make(map[string]bool)
	for _, t := range model.Types {
		pack := t.Package
		if handled[pack] {
			continue
		}
		handled[pack] = true
		packs = append(packs, pack)
	}
	sort.Strings(packs)
	return packs
}

func (model *YamlModel) IsRootPackage(pack string) bool {
	return pack == ""
}

func (model *YamlModel) EachEnum(consumer func(enum *YamlEnum)) {
	for i := range model.Types {
		t := model.Types[i]
		if t.IsEnum() {
			consumer(t.Enum)
		}
	}
}

func (model *YamlModel) EachClass(consumer func(class *YamlClass)) {
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

func (model *YamlModel) PackageOfClass(class *YamlClass) string {
	for _, t := range model.Types {
		if t.IsClass() && t.Class.Name == class.Name {
			return t.Package
		}
	}
	return ""
}

func (model *YamlModel) PackageOfEnum(enum *YamlEnum) string {
	for _, t := range model.Types {
		if t.IsEnum() && t.Enum.Name == enum.Name {
			return t.Package
		}
	}
	return ""
}

func (model *YamlModel) IsEmpty() bool {
	return len(model.Types) == 0
}

func ReadYamlModel(args *args) (*YamlModel, error) {

	types := make([]*YamlType, 0)

	// collect Yaml definitions from the `yaml` folder and its sub-folders
	yamlRoot := filepath.Join(args.home, "yaml")
	dirQueue := make([]string, 1, 5)
	dirQueue[0] = yamlRoot
	for len(dirQueue) > 0 {
		nextDir := dirQueue[0]
		dirQueue = dirQueue[1:]

		pack := ""
		if nextDir != yamlRoot {
			pack = filepath.Base(nextDir)
		}
		log.Println("parse folder:", nextDir)

		files, err := os.ReadDir(nextDir)
		if err != nil {
			return nil, err
		}

		for _, file := range files {
			name := file.Name()
			if file.IsDir() {
				dirQueue = append(dirQueue, filepath.Join(nextDir, name))
				continue
			}

			log.Println("Parse YAML file", name)
			path := filepath.Join(nextDir, name)
			data, err := os.ReadFile(path)
			if err != nil {
				return nil, err
			}
			typeDef := &YamlType{}
			if err := yaml.Unmarshal(data, typeDef); err != nil {
				return nil, err
			}
			typeDef.Package = pack

			types = append(types, typeDef)
		}

	}

	log.Println("Collected", len(types), "YAML types")

	typeMap := make(map[string]*YamlType)
	for i := range types {
		typeDef := types[i]
		if typeMap[typeDef.Name()] != nil {
			log.Fatalln("Duplicate type name: '"+typeDef.Name()+"'.",
				"The name of each type must be unique.")
		}

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

// IsRootEntity returns true if the given class is a root entity. This is the case
// when `RootEntity` is a parent class of the given class.
func (model *YamlModel) IsRootEntity(class *YamlClass) bool {
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

// TopoSortClasses returns the classes of the given Yaml model in topological
// order: A class X that has a dependency to a class Y, comes after Y in the
// returned slice.
func (model *YamlModel) TopoSortClasses() []*YamlClass {

	// check if there is a link between a class A and another class B where B is
	// dependent from A. B is dependent from A if it has a property of type A.
	isLinked := func(class, dependent *YamlClass) bool {
		if class == dependent {
			return false
		}
		for _, prop := range dependent.Props {
			propType := YamlPropType(prop.Type)
			if propType.IsList() {
				propType = propType.UnpackList()
			}
			if propType.ToPython() == class.Name {
				return true
			}
		}
		return false
	}

	// collect the dependencies
	dependencyCount := make(map[string]int)
	dependents := make(map[string][]string)
	model.EachClass(func(class *YamlClass) {
		if _, ok := dependencyCount[class.Name]; !ok {
			dependencyCount[class.Name] = 0
		}
		model.EachClass(func(dependent *YamlClass) {
			if isLinked(class, dependent) {
				c := class.Name
				d := dependent.Name
				dependencyCount[d] += 1
				dependents[c] = append(dependents[c], d)
			}
		})
	})

	// make sure that every RootEntity is dependent from 'Ref' as we generate a
	// to_ref method where the Ref type should be known
	refDeps, ok := dependents["Ref"]
	if !ok {
		refDeps = make([]string, 0)
	}
	model.EachClass(func(class *YamlClass) {
		if !model.IsRootEntity(class) && class.Name != "Unit" {
			return
		}
		contains := false
		for _, dep := range refDeps {
			if class.Name == dep {
				contains = true
				break
			}
		}
		if !contains {
			refDeps = append(refDeps, class.Name)
			dependencyCount[class.Name] += 1
		}
	})
	dependents["Ref"] = refDeps

	// sort dependencies in topological order
	order := make([]string, 0)
	for len(dependencyCount) > 0 {

		// find next node with no dependencies; if there are multiple options, try
		// to do this in alphabetical order so that we get a stable sort order
		node := ""
		for n, count := range dependencyCount {
			if count > 0 {
				continue
			}
			if node == "" ||
				strings.Compare(strings.ToLower(n), strings.ToLower(node)) < 0 {
				node = n
			}
		}

		if node == "" {
			log.Println("ERROR: could not sort classes in topological order")
			break
		}
		delete(dependencyCount, node)
		order = append(order, node)

		// remove the handled dependency from its dependents
		for _, dependent := range dependents[node] {
			dependencyCount[dependent] -= 1
		}
	}

	sorted := make([]*YamlClass, 0, len(order))
	for _, name := range order {
		next := model.TypeMap[name]
		if next != nil && next.IsClass() {
			sorted = append(sorted, next.Class)
		}
	}
	return sorted
}
