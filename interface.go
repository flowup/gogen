package gogen

import (
	"go/ast"
)

// Interface represents the interface type of a
// given build
type Interface struct {
	BaseType
	parent    *ast.InterfaceType
	spec      *ast.TypeSpec
	methods   []*ast.Field
}

// NewInterface creates a new Interface type and returns
// it with the provided parent and spec.
func NewInterface(parent *ast.InterfaceType, spec *ast.TypeSpec, annotationMap *AnnotationMap) *Interface {

	return &Interface{
		BaseType: BaseType{
			name:        spec.Name.Name,
			annotations: annotationMap,
		},
		parent:  parent,
		spec:    spec,
		methods: parent.Methods.List,
	}
}

// ParseInterface will create an interface
// with parameters given and return it
func ParseInterface(spec *ast.TypeSpec, parent *ast.InterfaceType, comments ast.CommentMap) *Interface {
	return NewInterface(parent, spec, ParseAnnotations(comments))
}
