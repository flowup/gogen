package gogen

import (
	"go/ast"
	"fmt"
)

// Array represents the interface type of a
// given build
type Array struct {
	BaseType
}

// Name will return a name of a declaration
func (b *Array) Name() string {
	return b.name
}

// Type will return a type of a declaration
func (b *Array) Type() string {
	return b.baseType
}

// NewArray creates a new Array type and returns
// it with the provided parent and spec.
func NewArray(parent *ast.ArrayType, spec *ast.TypeSpec, annotationMap *AnnotationMap) *Array {
	return &Array{
		BaseType: BaseType{
			name:        spec.Name.Name,
			annotations: annotationMap,
			baseType:    fmt.Sprintf("%s", parent.Elt),
		},
	}
}

// ParseArray will create an array
// with parameters given and return it
func ParseArray(spec *ast.TypeSpec, parent *ast.ArrayType, comments ast.CommentMap) *Array {
	return NewArray(parent, spec, ParseAnnotations(comments))
}
