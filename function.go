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

// ParseFunction will create and return a structure
// for a function in a build
func ParseFunction(parent *ast.FuncDecl, comments ast.CommentMap) *Function {
	f := &Function{
		parent: parent,
	}
	f.name = parent.Name
	f.tags = ParseTags(comments)
	return f
}