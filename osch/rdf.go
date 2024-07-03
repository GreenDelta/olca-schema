package main

import (
	"bytes"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type rdfWriter struct {
	buff  *bytes.Buffer
	model *YamlModel
}

func (w *rdfWriter) buffer() *bytes.Buffer {
	return w.buff
}

func (w *rdfWriter) indent() string {
	return "  "
}

type rdfProp struct {
	domain []*YamlClass
	yaml   *YamlProp
}

func (p *rdfProp) domainDef() string {
	if len(p.domain) == 1 {
		class := p.domain[0]
		return ":" + class.Name
	}
	text := "[ a owl:Class; owl:unionOf ("
	for i, class := range p.domain {
		if i > 0 {
			text += " "
		}
		text += ":" + class.Name
	}
	return text + ")]"
}

func (p *rdfProp) rangeDef() string {
	var _range func(t string) string
	_range = func(t string) string {
		if strings.HasPrefix(t, "Ref[") {
			return ":Ref"
		}
		if strings.HasPrefix(t, "List[") {
			return _range(t[5 : len(t)-1])
		}
		switch t {
		case "string":
			return "xsd:string"
		case "int", "integer":
			return "xsd:integer"
		case "double", "float":
			return "xsd:double"
		case "bool", "boolean":
			return "xsd:boolean"
		case "date":
			return "xsd:date"
		case "dateTime":
			return "xsd:dateTime"
		case "GeoJSON":
			return "http://purl.org/geojson/vocab#Feature"
		default:
			return ":" + t
		}
	}
	return _range(p.yaml.Type)
}

func (p *rdfProp) name() string {
	return p.yaml.Name
}

func writeRdf(args *args) {

	model, err := ReadYamlModel(args)
	check(err, "could not read YAML model")
	outDir := filepath.Join(args.home, "docs")
	mkdir(outDir)

	var buffer bytes.Buffer
	w := &rdfWriter{
		buff:  &buffer,
		model: model,
	}

	w.header()

	model.EachEnum(func(enum *YamlEnum) {
		ln(w, ":", enum.Name, " a rdfs:Class;")
		lni(w, 1, "rdfs:subClassOf :Enumeration;")
		lni(w, 1, "rdfs:comment \"", strip(enum.Doc), "\"")
		lni(w, 1, ".\n")

		for _, item := range enum.Items {
			ln(w, ":", item.Name, " a rdfs:Class;")
			lni(w, 1, "rdfs:subClassOf :", enum.Name, ";")
			lni(w, 1, "rdfs:comment \"", strip(item.Doc), "\"")
			lni(w, 1, ".\n")
		}
	})

	model.EachClass(func(class *YamlClass) {
		ln(w, ":", class.Name, " a rdfs:Class;")
		if class.SuperClass != "" {
			lni(w, 1, "rdfs:subClassOf :", class.SuperClass, ";")
		}
		lni(w, 1, "rdfs:comment \"", strip(class.Doc), "\"")
		lni(w, 1, ".\n")
	})

	w.writeProps()
	os.WriteFile(filepath.Join(outDir, "schema.ttl"), w.buff.Bytes(), os.ModePerm)
}

func (w *rdfWriter) header() {
	ln(w, "@prefix : <http://greendelta.github.io/olca-schema#> .")
	ln(w, "@prefix owl: <http://www.w3.org/2002/07/owl#> .")
	ln(w, "@prefix rdfs: <http://www.w3.org/2000/01/rdf-schema#> .")
	ln(w, "@prefix rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#> .")
	ln(w, "@prefix xsd: <http://www.w3.org/2001/XMLSchema#> .")
	ln(w)
	ln(w, `:Enumeration a rdfs:Class;
  rdfs:comment "The super-class of all enumeration types."
  .`)
	ln(w)
}

func (w *rdfWriter) writeProps() {
	for _, prop := range w.collectProps() {
		ln(w, ":", prop.name(), " a rdf:Property;")
		if len(prop.domain) == 1 {
			lni(w, 1, "rdfs:comment \"", strip(prop.yaml.Doc), "\";")
		}
		lni(w, 1, "rdfs:domain ", prop.domainDef(), ";")
		lni(w, 1, "rdfs:range ", prop.rangeDef())
		lni(w, 1, ".\n")
	}
}

func (w *rdfWriter) collectProps() []*rdfProp {
	dict := make(map[string]*rdfProp)
	w.model.EachClass(func(class *YamlClass) {
		for i, prop := range class.Props {
			if strings.HasPrefix(prop.Name, "@") {
				continue
			}
			p := dict[prop.Name]
			if p == nil {
				p = &rdfProp{
					yaml: class.Props[i],
				}
				dict[prop.Name] = p
			}
			p.domain = append(p.domain, class)
		}
	})
	props := make([]*rdfProp, 0, len(dict))
	for _, p := range dict {
		props = append(props, p)
	}
	sort.Slice(props, func(i, j int) bool {
		return strings.Compare(props[i].name(), props[j].name()) < 0
	})
	return props
}
