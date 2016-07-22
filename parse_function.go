package gogen

import "go/ast"

// Function represents a function of a given build
type Function struct {
	parent *ast.FuncDecl

	tags *TagMap
}

// Name returns the name of the function
func (f *Function) Name() string {
	return f.parent.Name.Name
}

// Tags returns the tags of the function
func (f *Function) Tags() *TagMap {
  return f.tags
}

// IsMethod returns true if the function has receiver
func (f *Function) IsMethod() bool {
	return f.parent.Recv != nil
}

// ParseFunction will create and return a structure
// for a function in a build
func ParseFunction(parent *ast.FuncDecl, comments ast.CommentMap) *Function {
	return &Function{
		parent: parent,
		tags: ParseTags(comments),
	}
}