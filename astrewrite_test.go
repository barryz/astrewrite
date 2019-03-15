package astrewrite

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestWriteFile(t *testing.T) {
	fset := token.NewFileSet()
	filename := "astrewrite_example_test.go"

	file, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	rewriteFunc := func(n ast.Node) (ast.Node, bool) {
		return n, true
	}

	rn := Walk(file, rewriteFunc)

	if err := WriteFile(fset, rn, filename); err != nil {
		t.Fatal(err)
	}
}
