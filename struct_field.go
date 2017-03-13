package gogen

import (
	"go/ast"
	"fmt"
)

// These value are used in determining type of struct field
const (
	PrimitiveType = 0
	SliceType     = 1
	MapType       = 2
	StructType    = 3
	SelectorType  = 4
	Unrecognized  = 5
	PointerType   = 6
)

// StructField encapsulates one field of the Structure
type StructField struct {
	BaseType

	parent *ast.Field
}

// NewStructField will return a new struct field with
// given parent and tag map
func NewStructField(parent *ast.Field, annotations *AnnotationMap) *StructField {
	return &StructField{
		BaseType: BaseType{
			annotations: annotations,
		},
		parent: parent,
	}
}

// Type returns type of the field as a string and
// FieldType such as Slice, Map or Primitive
func (f *StructField) Type() (string, int) {
	if f == nil {
		return "", Unrecognized
	}

	switch t := f.parent.Type.(type) {
	case *ast.Ident:
		if t.Obj != nil {
			return t.Name, StructType
		}
		return t.Name, PrimitiveType
	case *ast.ArrayType:
		return t.Elt.(*ast.Ident).Name, SliceType
	case *ast.MapType:
		return "[" + t.Key.(*ast.Ident).Name + "]" + t.Value.(*ast.Ident).Name, MapType
	// imported types
	case *ast.SelectorExpr:
		return t.X.(*ast.Ident).Name + "." + t.Sel.Name, SelectorType
	case *ast.StarExpr:
		if ident, ok := t.X.(*ast.Ident); ok {
			return ident.Name, PointerType
		}

		if expr, ok := t.X.(*ast.SelectorExpr); ok {
			return fmt.Sprintf("%s.%s", expr.X, expr.Sel.Name), PointerType
		}
		return "", PointerType
	default:
		panic("StructField type not recognized! Please report this issue.")
	}
}

// Tag returns tag of the field as a string
func (f *StructField) Tag() (string) {
	if f.parent.Tag != nil {
		return string([]byte(f.parent.Tag.Value)[1:len(f.parent.Tag.Value)-1])
	}
	return ""
}

// ParseStructField will create a struct field
// with given parameters and return it
func ParseStructField(parent *ast.Field, comments ast.CommentMap) *StructField {
	sf := NewStructField(parent, ParseAnnotations(comments))

	return sf
}
