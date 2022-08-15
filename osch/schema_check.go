package main

import (
	"fmt"
	"sort"
	"strings"
)

func checkSchema(args *args) {
	model, err := ReadYamlModel(args.yamlDir)
	if err != nil {
		fmt.Println("ERROR: Failed to parse YAML model:", err)
		return
	}

	checkClassHierarchy(model)
	checkBooleanPrefixes(model)
	checkPropertyOrder(model)
	checkFieldIndices(model)
}

func checkClassHierarchy(model *YamlModel) {
	// check that every class begins in Entity
	for _, t := range model.Types {
		if t.IsEnum() {
			continue
		}
		class := t.Class
		for {
			if class.Name == "Entity" {
				break
			}
			parent := model.ParentOf(class)
			if parent == nil {
				fmt.Println("ERROR: class hierarchy of '" +
					class.Name + "' does not starts in `Entity`")
				break
			}
			class = parent
		}
	}
}

func checkBooleanPrefixes(model *YamlModel) {
	boolPrefs := []string{
		"has", "is", "with",
	}
	for _, t := range model.Types {
		if t.IsEnum() {
			continue
		}
		for _, prop := range t.Class.Props {
			if prop.Type != "boolean" {
				continue
			}
			valid := false
			for _, pref := range boolPrefs {
				if strings.HasPrefix(prop.Name, pref) {
					valid = true
					break
				}
			}
			if !valid {
				fmt.Println("WARNING: boolean property '" + prop.Name +
					"' in class '" + t.Name() +
					"' should start with 'is', 'has', or 'with'")
			}
		}
	}
}

func checkPropertyOrder(model *YamlModel) {
	// Check if the properties in the classes are sorted by name. This is just for
	// the initial schema creation and should be removed later.
	for _, t := range model.Types {
		if t.IsEnum() {
			continue
		}
		c := t.Class
		props := c.Props
		sorted := true
		var last *YamlProp
		for i := range props {
			prop := props[i]
			if i == 0 {
				last = props[i]
				continue
			}
			if strings.Compare(prop.Name, last.Name) < 0 || prop.Index < last.Index {
				sorted = false
				break
			}
		}

		if sorted {
			continue
		}

		fmt.Println("WARNING: properties not in order:", c.Name)
		sort.Sort(YamlPropsByName(c.Props))
		for i, p := range c.Props {
			fmt.Println("  o + ", i, p.Name)
		}
	}
}

func checkFieldIndices(model *YamlModel) {

	for _, t := range model.Types {
		if t.IsClass() {
			continue
		}
		usedIndices := make(map[int]bool)
		for _, item := range t.Enum.Items {
			idx := item.Index
			if idx < 1 {
				fmt.Println("ERROR: invalid index", idx, "for item",
					item.Name, "in enum", t.Name())
				continue
			}
			if usedIndices[idx] {
				fmt.Println("ERROR: index", idx, "for item", item.Name,
					"in enum", t.Name(),
					"is already used by some other item in that enum")
				continue
			}
			usedIndices[idx] = true
		}
	}

	// check properties in classes
	for _, t := range model.Types {
		if t.IsEnum() {
			continue
		}
		props := model.AllPropsOf(t.Class)
		usedIndices := make(map[int]bool)
		for _, prop := range props {
			idx := prop.Index
			if idx < 1 {
				fmt.Println("ERROR: invalid index", idx, "for property",
					prop.Name, "in class", t.Name())
				continue
			}
			if usedIndices[idx] {
				fmt.Println("ERROR: index", idx, "for property", prop.Name,
					"in class", t.Name(),
					"is already used by some other property in the hierarchy")
				continue
			}
			usedIndices[idx] = true
		}
	}

}
