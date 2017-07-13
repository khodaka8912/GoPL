package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Node interface{} // CharData or *Element

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func main() {
	tree := parseXmlTree(os.Stdin)
	for _, elm := range tree {
		fmt.Println(elm)
	}
}

func (e *Element) String() string {
	buf := bytes.NewBufferString("<")
	buf.WriteString(e.Type.Local)
	for _, attr := range e.Attr {
		fmt.Fprintf(buf, " %s=%q", attr.Name.Local, attr.Value)
	}
	if len(e.Children) == 0 {
		buf.WriteString("/>")
		return buf.String()
	}
	buf.WriteString(">\n")
	for _, node := range e.Children {
		switch node := node.(type) {
		case CharData:
			buf.WriteString(string(node))
			buf.WriteRune('\n')
		case *Element:
			buf.WriteString(node.String())
		}
	}
	fmt.Fprintf(buf, "</%s>\n", e.Type.Local)
	return buf.String()
}

func parseXmlTree(in io.Reader) (roots []*Element) {
	dec := xml.NewDecoder(in)
	var stack []*Element
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			elm := &Element{Type: tok.Name, Attr: tok.Attr}
			if len(stack) == 0 {
				roots = append(roots, elm)
			} else {
				parent := len(stack) - 1
				stack[parent].Children = append(stack[parent].Children, elm)
			}
			stack = append(stack, elm)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if parent := len(stack) - 1; parent >= 0 {
				stack[parent].Children = append(stack[parent].Children, CharData(tok))
			}
		}
	}
	return roots
}
