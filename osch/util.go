package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

// Converts the given identifier from camelCase to snake_case.
func toSnakeCase(identifier string) string {
	var buff bytes.Buffer
	for i, char := range identifier {
		if i > 0 && unicode.IsUpper(char) {
			buff.WriteRune('_')
		}
		buff.WriteRune(unicode.ToLower(char))
	}
	return buff.String()
}

// Formats the given comment to have a line length of max. 80 characters.
func formatComment(comment string, indent string) string {
	if strings.TrimSpace(comment) == "" {
		return ""
	}

	// split words by whitespaces
	var words []string
	var word bytes.Buffer
	for _, char := range comment {
		if unicode.IsSpace(char) {
			if word.Len() > 0 {
				words = append(words, word.String())
			}
			word.Reset()
			continue
		}
		word.WriteRune(char)
	}
	if word.Len() > 0 {
		words = append(words, word.String())
	}
	if len(words) == 0 {
		return ""
	}

	// format the comment
	text := ""
	line := indent + "//"
	for _, word := range words {
		nextLine := line + " " + word
		if len(nextLine) < 80 {
			line = nextLine
		} else {
			text += line + "\n"
			line = indent + "// " + word
		}
	}
	if line != indent+"// " {
		text += line + "\n"
	}
	return text
}

func writeFile(file, content string) {
	err := ioutil.WriteFile(file, []byte(content), os.ModePerm)
	check(err, "failed to write file: "+file)
}

func isDir(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

func mkdir(path string) {
	stat, err := os.Stat(path)

	if err == nil {
		if stat.IsDir() {
			return
		}
		fmt.Println("ERROR:", path, "exists but it is not a folder")
		os.Exit(1)
	}
	if os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
	}
	check(err, "could not create folder: "+path)
}

func cleanDir(path ...string) string {
	dir := filepath.Join(path...)
	_, err := os.Stat(dir)
	if err == nil {
		os.RemoveAll(dir)
	}
	mkdir(dir)
	return dir
}

func startsWithLower(s string) bool {
	if s == "" {
		return false
	}
	for _, char := range s {
		if unicode.IsLower(char) {
			return true
		} else {
			return false
		}
	}
	return false
}

type Buffer struct {
	buff *bytes.Buffer
}

func NewBuffer() *Buffer {
	return &Buffer{buff: &bytes.Buffer{}}
}

func (b *Buffer) String() string {
	return b.buff.String()
}

func (b *Buffer) Writeln(args ...string) {
	for i, arg := range args {
		if i > 0 {
			b.buff.WriteRune(' ')
		}
		b.buff.WriteString(arg)
	}
	b.buff.WriteRune('\n')
}
