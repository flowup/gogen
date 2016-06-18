package gogen

import "go/ast"

type Function struct {
	parent *ast.FuncDecl
}

// Name returns the name of the function
func (f *Function) Name() string {
	return f.parent.Name.Name
}

// IsMethod returns true if the function has receiver
func (f *Function) IsMethod() bool {
	return f.parent.Recv != nil
}

func ParseFunction(parent *ast.FuncDecl) *Function {
	return &Function{
		parent,
	}
}