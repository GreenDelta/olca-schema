package main

import "bytes"

type buffer interface {
	buffer() *bytes.Buffer
	indent() string
}

func wln(buff buffer, xs ...string) {
	text := buff.buffer()
	for _, x := range xs {
		text.WriteString(x)
	}
	text.WriteRune('\n')
}

func wlni(buff buffer, i int, xs ...string) {
	text := buff.buffer()
	indent := buff.indent()
	for j := 0; j < i; j++ {
		text.WriteString(indent)
	}
	wln(buff, xs...)
}
