package main

import (
	"log"
	"strings"
)

type YamlProp struct {
	Name       string `yaml:"name"`
	Index      int    `yaml:"index"`
	Type       string `yaml:"type"`
	Doc        string `yaml:"doc"`
	IsOptional bool   `yaml:"optional"`
}

func (p *YamlProp) PropType() YamlPropType {
	return YamlPropType(p.Type)
}

func (prop *YamlProp) PyName() string {
	name := prop.Name
	switch name {
	case "@type":
		return "schema_type"
	case "@id":
		return "id"
	case "from":
		return "from_"
	default:
		return toSnakeCase(name)
	}
}

type YamlPropsByName []*YamlProp

func (s YamlPropsByName) Len() int { return len(s) }

func (s YamlPropsByName) Less(i, j int) bool {
	name_i := s[i].Name
	name_j := s[j].Name
	if name_i == name_j {
		return false
	}
	firstOrder := []string{"@type", "@id"}
	for _, f := range firstOrder {
		if name_i == f {
			return true
		}
		if name_j == f {
			return false
		}
	}
	return strings.Compare(name_i, name_j) < 0
}

func (s YamlPropsByName) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type YamlPropType string

func (t YamlPropType) IsList() bool {
	return strings.HasPrefix(string(t), "List[")
}

func (t YamlPropType) IsPrimitive() bool {
	return startsWithLower(string(t))
}

func (t YamlPropType) IsEnumOf(model *YamlModel) bool {
	if t := model.TypeMap[string(t)]; t != nil && t.IsEnum() {
		return true
	} else {
		return false
	}
}

func (t YamlPropType) IsClassOf(model *YamlModel) bool {
	if t := model.TypeMap[string(t)]; t != nil && t.IsClass() {
		return true
	} else {
		return false
	}
}

func (t YamlPropType) UnpackList() YamlPropType {
	s := strings.TrimPrefix(string(t), "List[")
	return YamlPropType(strings.TrimSuffix(s, "]"))
}

func (t YamlPropType) IsRef() bool {
	return strings.HasPrefix(string(t), "Ref[")
}

func (t YamlPropType) UnpackRef() YamlPropType {
	s := strings.TrimPrefix(string(t), "Ref[")
	return YamlPropType(strings.TrimSuffix(s, "]"))
}

func (t YamlPropType) ToPython() string {
	if t.IsList() {
		param := t.UnpackList()
		return "List[" + param.ToPython() + "]"
	}
	if t.IsRef() {
		return "Ref"
	}
	switch t {
	case "string", "date", "dateTime":
		return "str"
	case "double", "float":
		return "float"
	case "int", "integer":
		return "int"
	case "bool", "boolean":
		return "bool"
	case "GeoJSON":
		return "Dict[str, Any]"
	default:
		if startsWithLower(string(t)) {
			log.Println("WARNING: unknown primitive type:", t)
			return "object"
		} else {
			return string(t)
		}
	}
}
