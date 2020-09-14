package typetools

import (
	"strings"
	"testing"
)

func TestParseFromPkg(t *testing.T) {
	tcs := [...]struct {
		name          string
		src           string
		expectedTypes map[string]string
		expectedError string
	}{
		{
			name: "Empty",
			src:  `package main`,
		},
		{
			name:          "Syntax error",
			src:           `var x int`, // Needs package to be syntactically correct
			expectedError: "Cannot parse source",
		},
		{
			name: "Primitive int",
			src: `package main
var x int`,
			expectedTypes: map[string]string{
				"x": "int",
			},
		},
		{
			name: "Primitive int with main function",
			src: `package main
var x int
func main() {}`,
			expectedTypes: map[string]string{
				"x":    "int",
				"main": "func()",
			},
		},
		{
			name: "Type definition",
			src: `package main
type Num int`,
			expectedTypes: map[string]string{
				"Num": "Num",
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			ts, err := ParseFromPkg(tc.src)
			if tc.expectedError == "" {
				if err != nil {
					t.Error(err)
				}
				if want, got := len(tc.expectedTypes), len(ts); want != got {
					t.Fatalf("Expecting %d items in map but got %d", want, got)
				}
				for name, typ := range tc.expectedTypes {
					if _, ok := ts[name]; !ok {
						t.Fatalf("Expected `%s` to have type `%s` but not parsed", name, typ)
					}
					if want, got := typ, ts[name].String(); want != got {
						t.Fatalf("Mismatched types\nWant `%v`\nGot  `%v`", want, got)
					}
				}
			} else {
				if err == nil {
					t.Fatalf("Expecting error but got none:\nWant `%s`", tc.expectedError)
				}
				if !strings.HasPrefix(err.Error(), tc.expectedError) {
					t.Fatalf("Expecting a different error:\nWant `%s`\nGot  `%v`", tc.expectedError, err)
				}
			}
		})
	}
}
