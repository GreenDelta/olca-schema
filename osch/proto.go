package main

import (
	"bytes"
	"strconv"
	"strings"
	"unicode"
)

// FileHeader is the file header that is written to the generated proto3 file.
// This is the place where you want to define global options
const FileHeader = `// Generated from olca-schema (https://github.com/GreenDelta/olca-schema).
// DO NOT EDIT!

syntax = "proto3";

package protolca;

option csharp_namespace = "ProtoLCA";
option go_package = ".;protolca";
option java_package = "org.openlca.proto";
option java_outer_classname = "Proto";
option java_multiple_files = true;


`

// BytesHint is a comment we add to fields with `bytes` as data type.
const BytesHint = `  // When we map to the bytes type it means that we have no matching message
  // type and just put the raw bytes into the field. This is specifically true
  // for our geometry data of locations which cannot be translated to valid
  // GeoJSON using Protocol Buffers (as they do not support arrays of arrays).
  // To indicate that this is a different field than the field in the
  // olca-schema definition, we append the _bytes suffix to the field name
`

const FileFooter = `
// This enumeration type is added for compatibility with the @type attribute of
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

func generateProto(yaml *YamlModel) string {
	var buff bytes.Buffer
	buff.WriteString(FileHeader)

	// write the message and enumeration types
	for _, typeDef := range yaml.Types {
		switch typeDef.Name() {
		case "Entity", "RefEntity", "RootEntity":
			continue
		}

		// write a class definition
		class := typeDef.Class
		if class != nil {
			comment := formatComment(class.Doc, "")
			if comment != "" {
				buff.WriteString(comment)
			}
			buff.WriteString("message Proto" + class.Name + " {\n\n")
			writeProtoFields(class, &buff, yaml.TypeMap)
			buff.WriteString("}\n\n")
			continue
		}

		// write an enumeration
		enum := typeDef.Enum
		if enum != nil {
			if enum.Name == "ModelType" {
				continue
			}
			comment := formatComment(enum.Doc, "")
			if comment != "" {
				buff.WriteString(comment)
			}

			buff.WriteString("enum Proto" + enum.Name + " {\n\n")

			buff.WriteString("  // This default option was added automatically\n")
			buff.WriteString("  // and means that no values was set.\n")
			buff.WriteString("  " + protoUndefinedOf(enum) + " = 0;\n\n")
			for i, item := range enum.Items {
				comment := formatComment(item.Doc, "  ")
				if comment != "" {
					buff.WriteString(comment)
				}
				buff.WriteString("  " + item.Name + " = " +
					strconv.Itoa(i+1) + ";\n\n")
			}
			buff.WriteString("}\n\n")
		}
	}

	buff.WriteString(FileFooter)
	return buff.String()
}

// Writes the fields of the given class to the given buffer. This function
// climbs up the class hierarchy and inlines the fields of the corresponding
// super classes (as there is no extension mechanism in proto3).
func writeProtoFields(class *YamlClass, buff *bytes.Buffer, types map[string]*YamlType) {

	// write fields of super classes recursively
	if class.SuperClass != "" {
		super := types[class.SuperClass]
		if super != nil && super.Class != nil {
			writeProtoFields(super.Class, buff, types)
		}
	}

	// @type field
	if class.Name == "Ref" || class.Name == "RootEntity" {
		buff.WriteString("  // The type name of the respective entity.\n")
		buff.WriteString("  ProtoType type = 1 [json_name = \"@type\"];\n\n")
	}

	// @id field
	if class.Name == "RefEntity" {
		buff.WriteString("  // The reference ID (typically an UUID) of the entity.\n")
		buff.WriteString("  string id = 2 [json_name = \"@id\"];\n\n")
	}

	// write fields
	for _, field := range class.Props {

		if strings.HasPrefix(field.Name, "@") {
			continue
		}

		// field comment
		comment := formatComment(field.Doc, "  ")
		if comment != "" {
			buff.WriteString(comment)
		}

		protoType := toProtoType(field.Type)
		protoField := toSnakeCase(field.Name)
		if protoType == "bytes" {
			buff.WriteString(BytesHint)
			protoField += "_bytes"
		}

		buff.WriteString("  " + protoType + " " + protoField +
			" = " + strconv.Itoa(field.Index) + ";\n\n")
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
