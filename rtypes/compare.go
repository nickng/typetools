package rtypes

import (
	"go/types"
	"reflect"
	"strings"
)

// TypeComponents represents different components of a type.
type TypeComponents struct {
	PkgPath string // Package path of type
	Name    string // Name of type in package
}

// IsPrimitive returns true if the type is a built-in promitive (no import path).
func (c TypeComponents) IsPrimitive() bool {
	return c.PkgPath == ""
}

// IsEqual returns true if the reflect and static type are nominally the same.
//
// i.e. same import package and same name within package.
// For most Go programs it's sufficient to consider them equivalent without a
// deeper comparison of type structure.
func IsEqual(rtype reflect.Type, typ types.Type) bool {
	rt := TypeComponents{
		PkgPath: rtype.PkgPath(),
		Name:    rtype.Name(),
	}
	t := splitTypeName(typ.String())
	return rt.PkgPath == t.PkgPath && rt.Name == t.Name
}

func splitTypeName(typ string) TypeComponents {
	if index := strings.LastIndex(typ, "."); index != -1 {
		return TypeComponents{
			PkgPath: typ[:index],
			Name:    typ[index+1:],
		}
	}
	return TypeComponents{Name: typ}
}
