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
	t := p.yaml.Type
	if strings.HasPrefix(t, "List[") {
		return "rdf:List"
	}
	if strings.HasPrefix(t, "Ref[") {
		return ":Ref"
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
	default:
		return ":" + t
	}
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
	w := rdfWriter{
		buff:  &buffer,
		model: model,
	}

	w.header()

	model.EachEnum(func(enum *YamlEnum) {
		w.ln(":", enum.Name, " a rdfs:Class;")
		w.ln("  rdfs:subClassOf :Enumeration;")
		w.ln("  rdfs:comment \"", strip(enum.Doc), "\"")
		w.ln("  .\n")

		for _, item := range enum.Items {
			w.ln(":", item.Name, " a rdfs:Class;")
			w.ln("  rdfs:subClassOf :", enum.Name, ";")
			w.ln("  rdfs:comment \"", strip(item.Doc), "\"")
			w.ln("  .\n")
		}
	})

	model.EachClass(func(class *YamlClass) {
		w.ln(":", class.Name, " a rdfs:Class;")
		if class.SuperClass != "" {
			w.ln("  rdfs:subClassOf :", class.SuperClass, ";")
		}
		w.ln("  rdfs:comment \"", strip(class.Doc), "\"")
		w.ln("  .\n")
	})

	w.writeProps()
	os.WriteFile(filepath.Join(outDir, "schema.ttl"), w.buff.Bytes(), os.ModePerm)
}

func (w *rdfWriter) header() {
	w.ln("@prefix : <http://greendelta.github.io/olca-schema#> .")
	w.ln("@prefix owl: <http://www.w3.org/2002/07/owl#> .")
	w.ln("@prefix rdfs: <http://www.w3.org/2000/01/rdf-schema#> .")
	w.ln("@prefix rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#> .")
	w.ln("@prefix xsd: <http://www.w3.org/2001/XMLSchema#> .")
	w.ln()
	w.ln(`:Enumeration a rdfs:Class;
  rdfs:comment "The super-class of all enumeration types."
  .`)
	w.ln()
}

func (w *rdfWriter) writeProps() {
	for _, prop := range w.collectProps() {
		w.ln(":", prop.name(), " a rdf:Property;")
		if len(prop.domain) == 1 {
			w.ln("  rdfs:comment \"", strip(prop.yaml.Doc), "\";")
		}
		w.ln("  rdfs:domain ", prop.domainDef(), ";")
		w.ln("  rdfs:range ", prop.rangeDef())
		w.ln("  .\n")
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

func (w *rdfWriter) ln(xs ...string) {
	for _, x := range xs {
		w.buff.WriteString(x)
	}
	w.buff.WriteRune('\n')
}
