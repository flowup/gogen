package gogen

import (
	"fmt"
	"go/ast"
)

// BaseType is a base structure of all declarations
// It can return a name and tag map.
type BaseType struct {
	name        string
	baseType    string
	annotations *AnnotationMap
}

// Name will return a name of a declaration
func (b *BaseType) Name() string {
	return b.name
}

// Type will return a type of a declaration
func (b *BaseType) Type() string {
	return b.baseType
}

// Annotations will return a tag map associated to a declaration
func (b *BaseType) Annotations() *AnnotationMap {
	return b.annotations
}

// NewBaseType creates a new baseType type and returns
// it with the provided parent and spec.
func NewBaseType(parent *ast.Expr, spec *ast.TypeSpec, annotationMap *AnnotationMap) *BaseType {
	return &BaseType{
		name:        spec.Name.Name,
		annotations: annotationMap,
		baseType:    fmt.Sprintf("%s", spec.Type),
	}
}

// ParseBaseType will create an basetype
// with parameters given and return it
func ParseBaseType(spec *ast.TypeSpec, parent ast.Expr, comments ast.CommentMap) *BaseType {
	return NewBaseType(&parent, spec, ParseAnnotations(comments))
}
