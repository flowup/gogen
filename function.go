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

func NewFunction(parent *ast.FuncDecl, tags *TagMap) *Function {
	return &Function{
		BaseType: BaseType{
			name: parent.Name.Name,
			tags: tags,
		},
		parent: parent,
	}
}

// ParseFunction will create and return a structure
// for a function in a build
func ParseFunction(parent *ast.FuncDecl, comments ast.CommentMap) *Function {
	f := NewFunction(parent, ParseTags(comments))
	return f
}