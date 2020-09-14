package rtypes

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestConvert(t *testing.T) {
	tests := [...]struct {
		name  string
		rtype reflect.Type
		typ   string
	}{
		{
			name:  "Primitive int",
			rtype: reflect.TypeOf(int(0)),
		},
		{
			name:  "stdlib json.Encoder",
			rtype: reflect.TypeOf(json.Encoder{}),
		},
		{
			name:  "external type",
			rtype: reflect.TypeOf(TypeComponents{}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			typ, err := Convert(test.rtype)
			if err != nil {
				t.Fatal(err)
			}
			if !IsEqual(test.rtype, typ) {
				t.Fatalf("The reflect Type and static type is not (shallow) equal")
			}
		})
	}
}
