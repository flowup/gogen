package architect

import "go/ast"

// Interface defines methods above the interface
// type for the interface build organization
type Interface interface {
	Name() string

	AddMethod(method Function)
	Method(name string) Function
}

// InterfaceImpl is a implmentation of the Interface
type InterfaceImpl struct {
	name    string
	methods []Function

	astInterface *ast.InterfaceType
}

// Name returns the name of the Interface
func (i *InterfaceImpl) Name() string {
	return i.name
}

// AddMethod adds given method into the existing set of
// methods
func (i *InterfaceImpl) AddMethod(method Function) {
	i.methods = append(i.methods, method)
}

// Method searches for the method by the name in
// the current structure instance
func (i *InterfaceImpl) Method(name string) Function {
	for _, m := range i.methods {
		if m.Name() == name {
			return m
		}
	}

	return nil
}