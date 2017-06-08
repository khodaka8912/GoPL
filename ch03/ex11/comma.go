package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Printf("%s -> %s\n", arg, comma(arg))
	}
}

func comma(s string) string {
	buf := bytes.Buffer{}
	if first := s[0]; first == '+' || first == '-' {
		buf.WriteByte(first)
		s = s[1:]
	}
	period := strings.Index(s, ".")
	if period == -1 {
		appendComma(s, &buf)
	} else {
		appendComma(s[:period], &buf)
		buf.WriteRune('.')
		appendCommaFraction(s[period+1:], &buf)
	}
	return buf.String()
}

func appendComma(s string, buf *bytes.Buffer) {
	n := len(s)
	if n <= 3 {
		buf.WriteString(s)
		return
	}
	i := (n-1)%3 + 1
	buf.WriteString(s[:i])
	for s = s[i:]; len(s) > 0; s = s[3:] {
		buf.WriteRune(',')
		buf.WriteString(s[:3])
	}
}

func appendCommaFraction(s string, buf *bytes.Buffer) {
	n := len(s)
	if n <= 3 {
		buf.WriteString(s)
		return
	}
	for ; len(s) > 3; s = s[3:] {
		buf.WriteString(s[:3])
		buf.WriteRune(',')
	}
	buf.WriteString(s)
}
