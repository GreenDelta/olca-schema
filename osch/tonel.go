package main

import (
	"bytes"
	"os"
	"path/filepath"
)

type tonelWriter struct {
	buff  *bytes.Buffer
	model *YamlModel
}

func writeTonelFiles(args *args) {

	model, err := ReadYamlModel(args)
	check(err, "could not read YAML model")

	outDir := filepath.Join(args.home, "build", "tonel")
	mkdir(outDir)

	model.EachClass(func(class *YamlClass) {
		var buffer bytes.Buffer
		writer := tonelWriter{
			buff:  &buffer,
			model: model,
		}
		writer.writeClass(class)

		outFile := filepath.Join(outDir, "Lca"+class.Name+".class.st")
		os.WriteFile(outFile, buffer.Bytes(), os.ModePerm)

	})

}

func (w *tonelWriter) writeClass(class *YamlClass) {

	className := "Lca" + class.Name

	// class declaration
	w.writeln("class {")
	w.writeln("\t#name : #", className, ",")
	var super string
	if class.SuperClass != "" {
		super = "Lca" + class.SuperClass
	} else {
		super = "Object"
	}
	w.writeln("\t#superclass : #", super, ",")
	w.writeln("\t#instVars : [")
	w.eachPropOf(class, func(prop string) {
		w.writeln("\t\t'", prop, "',")
	})
	w.writeln("\t],")
	w.writeln("\t#category : #'openLCA-Model'")
	w.writeln("}")

	// accessors
	w.eachPropOf(class, func(prop string) {
		w.writeln()

		// getter
		w.writeln("{ #category : #accessing }")
		w.writeln(className, " >> ", prop, " [")
		w.writeln()
		w.writeln("\t^ ", prop)
		w.writeln("]")
		w.writeln()

		// setter
		w.writeln("{ #category : #accessing }")
		w.writeln(className, " >> ", prop, ": anObject [")
		w.writeln()
		w.writeln("\t", prop, " := anObject")
		w.writeln("]")
		w.writeln()
	})

}

func (w *tonelWriter) eachPropOf(class *YamlClass, fn func(name string)) {
	for _, prop := range class.Props {
		if prop.Name == "@type" {
			continue
		}
		propName := prop.Name
		if propName == "@id" {
			propName = "id"
		}
		fn(propName)
	}
}

func (w *tonelWriter) writeln(xs ...string) {
	for _, x := range xs {
		w.buff.WriteString(x)
	}
	w.buff.WriteRune('\n')
}
