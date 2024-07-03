package main

import "bytes"

type buffer interface {
	buffer() *bytes.Buffer
	indent() string
}

type writer struct {
	buff *bytes.Buffer
}

func newWriter() *writer {
	return &writer{buff: &bytes.Buffer{}}
}

func (w *writer) buffer() *bytes.Buffer {
	return w.buff
}

func (w *writer) indent() string {
	return "  "
}

// ln concatenates the string arguments and writes them as a line into the
// given buffer.
func ln(buff buffer, xs ...string) {
	text := buff.buffer()
	for _, x := range xs {
		text.WriteString(x)
	}
	text.WriteRune('\n')
}

// lni concatenates the string arguments and writes them as an i-indented line
// into the given buffer where i is the number of indentations.
func lni(buff buffer, i int, xs ...string) {
	text := buff.buffer()
	indent := buff.indent()
	for j := 0; j < i; j++ {
		text.WriteString(indent)
	}
	ln(buff, xs...)
}
