package main

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"
)

// BytesHint is a comment we add to fields with `bytes` as data type.
const ProtoBytesHint = `  // When we map to the bytes type it means that we have no matching message
  // type and just put the raw bytes into the field. This is specifically true
  // for our geometry data of locations which cannot be translated to valid
  // GeoJSON using Protocol Buffers (as they do not support arrays of arrays).
  // To indicate that this is a different field than the field in the
  // olca-schema definition, we append the _bytes suffix to the field name
`

const ProtoRootFooter = `// This enumeration type is added for compatibility with the @type attribute of
// the openLCA JSON-LD format. In the proto messages we limit its usage to
// instances of CategorizedEntity and Ref while it is allowed for every type in
// the JSON-LD format. Thus, you should use ignoringUnknownFields flag when
// parsing openLCA JSON-LD messages with the generated proto parsers.
enum ProtoType {
  Undefined = 0;
  Actor = 1;
  Currency = 2;
  DQSystem = 3;
  Epd = 4;
  Flow = 5;
  FlowProperty = 6;
  ImpactCategory = 7;
  ImpactMethod = 8;
  Location = 9;
  NwSet = 10;
  Parameter = 11;
  Process = 12;
  ProductSystem = 13;
  Project = 14;
  Result = 15;
  SocialIndicator = 16;
  Source = 17;
  Unit = 18;
  UnitGroup = 19;
}
`

type protoWriter struct {
	buff  *bytes.Buffer
	model *YamlModel
	pack  string
}

func writeProtos(args *args) {
	model, err := ReadYamlModel(args)
	check(err, "could not read YAML model")

	buildDir := filepath.Join(args.home, "build")
	mkdir(buildDir)

	for _, pack := range model.Packages() {
		var buff bytes.Buffer
		w := protoWriter{
			buff:  &buff,
			model: model,
			pack:  pack,
		}
		w.writePack()

		fileName := "olca.proto"
		if pack != "" {
			fileName = "olca." + pack + ".proto"
		}
		outFile := filepath.Join(buildDir, fileName)
		log.Println("write file:", fileName)
		err = os.WriteFile(outFile, buff.Bytes(), os.ModePerm)
		check(err, "failed to write to file", outFile)
	}
}

func (w *protoWriter) writePack() {
	w.writeFileHeader()

	for _, t := range w.model.Types {
		if t.Package != w.pack {
			continue
		}
		if t.IsEnum() {
			w.writeEnum(t.Enum)
			w.writeln()
		}
		if t.IsClass() {
			class := t.Class
			comment := formatComment(class.Doc, "")
			if comment != "" {
				w.write(comment)
			}
			w.writeln("message Proto" + class.Name + " {")
			w.writeln()
			w.writeFieldsOf(class)
			w.writeln("}")
			w.writeln()
			w.writeln()
		}
	}

	if w.pack == "" {
		w.write(ProtoRootFooter)
	}

}

func (w *protoWriter) writeFileHeader() {
	w.writeln("// Generated from olca-schema (https://github.com/GreenDelta/olca-schema)")
	w.writeln("// DO NOT EDIT!")
	w.writeln()
	w.writeln("syntax = \"proto3\";")
	w.writeln()

	// package names
	if w.pack == "" {
		w.writeln("package protolca;")
		w.writeln("option java_package = \"org.openlca.proto\";")
	} else {
		w.writeln("package protolca." + w.pack + ";")
		w.writeln("option java_package = \"org.openlca.proto." + w.pack + "\";")
	}

	// other options
	w.writeln("option java_outer_classname = \"Proto\";")
	w.writeln("option go_package = \".;protolca\";")
	w.writeln("option csharp_namespace = \"ProtoLCA\";")
	w.writeln("option java_multiple_files = true;")
	w.writeln()

	// imports
	if w.pack != "" {
		w.writeln("import \"olca.proto\"")
		w.writeln()
	}
	w.writeln()
}

func (w *protoWriter) writeEnum(enum *YamlEnum) {
	if enum.Name == "ModelType" {
		return
	}
	comment := formatComment(enum.Doc, "")
	if comment != "" {
		w.write(comment)
	}
	w.writeln("enum Proto" + enum.Name + " {")
	w.writeln()

	// write the 'UNDEFINED_*' item
	w.indln("// This default option was added automatically")
	w.indln("// and means that no values was set.")
	w.indln(protoUndefinedOf(enum) + " = 0;")
	w.writeln()

	for _, item := range enum.Items {
		comment := formatComment(item.Doc, "  ")
		if comment != "" {
			w.write(comment)
		}
		w.indln(item.Name + " = " + strconv.Itoa(item.Index) + ";")
		w.writeln()
	}
	w.writeln("}")
	w.writeln()
}

// Writes the fields of the given class to the given buffer. This function
// climbs up the class hierarchy and inlines the fields of the corresponding
// super classes (as there is no extension mechanism in proto3).
func (w *protoWriter) writeFieldsOf(class *YamlClass) {

	// write fields of super classes recursively
	if class.SuperClass != "" {
		super := w.model.TypeMap[class.SuperClass]
		if super != nil && super.Class != nil {
			w.writeFieldsOf(super.Class)
		}
	}

	// @type field
	if class.Name == "Ref" || class.Name == "RootEntity" {
		w.indln("// The type name of the respective entity.")
		w.indln("ProtoType type = 1 [json_name = \"@type\"];")
		w.writeln()
	}

	// @id field
	if class.Name == "Ref" || class.Name == "RefEntity" {
		w.indln("// The reference ID (typically an UUID) of the entity.")
		w.indln("string id = 2 [json_name = \"@id\"];")
		w.writeln()
	}

	// write fields
	for _, field := range class.Props {

		if strings.HasPrefix(field.Name, "@") {
			continue
		}

		// field comment
		comment := formatComment(field.Doc, "  ")
		if comment != "" {
			w.write(comment)
		}

		protoType := toProtoType(field.Type)
		protoField := toSnakeCase(field.Name)
		if protoType == "bytes" {
			w.write(ProtoBytesHint)
			protoField += "_bytes"
		}

		w.write("  ")
		if field.IsOptional {
			w.write("optional ")
		}
		w.writeln(protoType + " " + protoField +
			" = " + strconv.Itoa(field.Index) + ";")
		w.writeln()
	}

}

// Maps the given olca-schema type to a corresponding proto3 type.
func toProtoType(schemaType string) string {
	switch schemaType {
	case "string", "double", "float":
		return schemaType
	case "dateTime", "date":
		return "string"
	case "int", "integer":
		return "int32"
	case "boolean":
		return "bool"
	case "GeoJSON":
		return "bytes"
	case "ModelType":
		return "ProtoCategoryType"
	}

	if strings.HasPrefix(schemaType, "Ref[") {
		return "ProtoRef"
	}
	if strings.HasPrefix(schemaType, "List[") {
		t := strings.TrimSuffix(
			strings.TrimPrefix(schemaType, "List["), "]")
		return "repeated " + toProtoType(t)
	}

	return "Proto" + schemaType
}

// Generates the name of the `UNDEFINED` option for the given
// enumeration type. As this option has to have a unique name
// we include the name of the enumeration into that name.
func protoUndefinedOf(enum *YamlEnum) string {
	var buff bytes.Buffer
	for _, char := range enum.Name {
		if unicode.IsUpper(char) {
			buff.WriteRune('_')
		}
		buff.WriteRune(char)
	}
	return "UNDEFINED" + strings.ToUpper(buff.String())
}

func (w *protoWriter) indln(s string) {
	w.writeln("  ", s)
}

func (w *protoWriter) writeln(xs ...string) {
	for _, x := range xs {
		w.buff.WriteString(x)
	}
	w.buff.WriteRune('\n')
}

func (w *protoWriter) write(s string) {
	w.buff.WriteString(s)
}
