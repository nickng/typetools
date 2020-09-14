package rtypes

import (
	"encoding/json"
	"reflect"
	"testing"

	"go.nickng.io/typetools"
)

func TestIsEqual(t *testing.T) {
	tests := [...]struct {
		name  string
		rtype reflect.Type
		typ   string
	}{
		{
			name:  "Primitive int",
			rtype: reflect.TypeOf(int(0)),
			typ:   `package main; var x int`,
		},
		{
			name:  "stdlib json.Encoder",
			rtype: reflect.TypeOf(json.Encoder{}),
			typ:   `package main; import "encoding/json"; var x json.Encoder`,
		},
		{
			name:  "external type",
			rtype: reflect.TypeOf(TypeComponents{}),
			typ:   `package main; import "go.nickng.io/typetools/rtypes"; var x rtypes.TypeComponents`,
		},
		{
			name:  "renamed external type",
			rtype: reflect.TypeOf(TypeComponents{}),
			typ:   `package main; import reflecttype "go.nickng.io/typetools/rtypes"; var x reflecttype.TypeComponents`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, err := typetools.ParseFromPkg(test.typ)
			if err != nil {
				t.Fatal(err)
			}
			IsEqual(test.rtype, ts["x"])
		})
	}
}
