package typetools

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"io"
	"strings"
)

// ParseFromPkg parses the types of objects in a package
// from a given source code src. For example,
//
//     package main
//     var x int           // x    → int type
//     func main() { ... } // main → func() type
//
func ParseFromPkg(src string) (map[string]types.Type, error) {
	parsed, err := parseReader(strings.NewReader(src))
	if err != nil {
		return nil, fmt.Errorf("Cannot parse source: %w", err)
	}
	var pkgTypes map[string]types.Type
	for _, n := range parsed.Package.Scope().Names() {
		obj := parsed.Package.Scope().Lookup(n)
		if pkgTypes == nil {
			pkgTypes = make(map[string]types.Type)
		}
		pkgTypes[n] = obj.Type()
	}
	return pkgTypes, nil
}

// parsed represent a parsed program.
type parsed struct {
	Package *types.Package // Package is the top-level type of the program
	File    *ast.File      // File is the AST of the parsed 'file'
}

// parseReader parses the given io.Reader r
// and returns the parsed type and AST.
func parseReader(r io.Reader) (*parsed, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "file.go", r, 0)
	if err != nil {
		return nil, fmt.Errorf("Cannot parse file: %w", err)
	}
	conf := types.Config{Importer: importer.Default()}
	pkg, err := conf.Check("", fset, []*ast.File{f}, nil)
	if err != nil {
		return nil, fmt.Errorf("Cannot type-check package: %w", err)
	}
	return &parsed{Package: pkg, File: f}, nil
}
