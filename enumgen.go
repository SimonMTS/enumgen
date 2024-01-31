package main

import (
	"strings"
	"unsafe"
)

func enumgen(g *Generator, typeName string, values []Value) {
	g.Printf("\n")
	g.Printf("\n")
	list(g, typeName, values)
	g.Printf("\n")
	g.Printf("\n")
	fromString(g, typeName, values)
	g.Printf("\n")
	g.Printf("\n")
	fromInt(g, typeName, values)
}

func list(g *Generator, typeName string, values []Value) {
	names := make([]string, 0, len(values))
	for _, value := range values {
		if *(*int64)(unsafe.Pointer(&value.value)) != -1 {
			names = append(names, value.originalName)
		}
	}

	var prefix string
	if *prefixFunctions {
		prefix = typeName
	}
	g.Printf("func %sList() []%s {", prefix, typeName)
	g.Printf("return []%s{%s}", typeName, strings.Join(names, ","))
	g.Printf("}")
}

func fromString(g *Generator, typeName string, values []Value) {
	var prefix string
	if *prefixFunctions {
		prefix = typeName
	}

	g.Printf("func %sFromString(s string) (%s, bool) {", prefix, typeName)
	g.Printf("switch s {")

	for _, value := range values {
		g.Printf(`case "%s": return %s, true;`, value.name, value.originalName)
	}

	if values[0].signed {
		g.Printf("default: return %s(-1), false;", typeName)
	} else {
		g.Printf("default: return %s(0), false;", typeName)
	}

	g.Printf("}")
	g.Printf("}")
}

func fromInt(g *Generator, typeName string, values []Value) {
	var prefix string
	if *prefixFunctions {
		prefix = typeName
	}

	g.Printf("func %sFromInt(n int) (%s, bool) {", prefix, typeName)
	g.Printf("switch n {")

	for _, value := range values {
		g.Printf(`case %s: return %s, true;`, &value, value.originalName)
	}

	if values[0].signed {
		g.Printf("default: return %s(-1), false;", typeName)
	} else {
		g.Printf("default: return %s(0), false;", typeName)
	}

	g.Printf("}")
	g.Printf("}")
}
