package main

import (
	"bytes"
	"os"
	"path/filepath"
)

type rdfWriter struct {
	buff  *bytes.Buffer
	model *YamlModel
}

func writeRdf(args *args) {

	model, err := ReadYamlModel(args)
	check(err, "could not read YAML model")
	outDir := filepath.Join(args.home, "build")
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

	os.WriteFile(filepath.Join(outDir, "schema.ttl"), w.buff.Bytes(), os.ModePerm)
}

func (w *rdfWriter) header() {
	w.ln("@prefix : <http://greendelta.github.io/olca-schema/schema.ttl#> .")
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

func (w *rdfWriter) ln(xs ...string) {
	for _, x := range xs {
		w.buff.WriteString(x)
	}
	w.buff.WriteRune('\n')
}
