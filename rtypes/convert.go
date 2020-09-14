package rtypes

import (
	"fmt"
	"go/types"
	"reflect"
	"strings"

	"go.nickng.io/typetools"
)

// Convert reads the type from reflection and
// convert it to a known public static type.
func Convert(rtype reflect.Type) (types.Type, error) {
	rt := TypeComponents{
		PkgPath: rtype.PkgPath(),
		Name:    rtype.Name(),
	}
	var src string
	if rt.IsPrimitive() {
		src = fmt.Sprintf("package main; var x %s", rt.Name)
	} else {
		var pkgName string
		if index := strings.LastIndex(rt.PkgPath, "/"); index == -1 {
			pkgName = rt.PkgPath
		} else {
			pkgName = rt.PkgPath[index+1:]
		}
		src = fmt.Sprintf("package main; import \"%s\"; var x %s.%s", rt.PkgPath, pkgName, rt.Name)
	}
	parsed, err := typetools.ParseFromPkg(src)
	if err != nil {
		return nil, err
	}
	return parsed["x"], nil
}
