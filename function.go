package gogen

import "go/ast"

// Function represents a function of a given build
type Function struct {
	BaseType
	parent *ast.FuncDecl
}

// IsMethod returns true if the function has receiver
func (f *Function) IsMethod() bool {
	return f.parent.Recv != nil
}

// NewFunction return initialization Function
func NewFunction(parent *ast.FuncDecl, annotations *AnnotationMap) *Function {
	return &Function{
		BaseType: BaseType{
			name: parent.Name.Name,
			annotations: annotations,
		},
		parent: parent,
	}
}

// ParseFunction will create and return a structure
// for a function in a build
func ParseFunction(parent *ast.FuncDecl, comments ast.CommentMap) *Function {
	f := NewFunction(parent, ParseAnnotations(comments))
	return f
}