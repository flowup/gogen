package gogen

import "go/ast"

// Structure represents the struct type of a
// given build
type Structure struct {
	parent *ast.StructType
	spec *ast.TypeSpec
}

// Name returns the name of the given structure
func (s *Structure) Name() string {
	return s.spec.Name.Name
}

// Interface represents the interface type of a
// given build
type Interface struct {
	parent *ast.InterfaceType
	spec *ast.TypeSpec
}

// Name returns the name of the interface type
func (i *Interface) Name()  string {
	return i.spec.Name.Name
}

func ParseStruct(spec *ast.TypeSpec, parent *ast.StructType) *Structure {
	s := &Structure{parent, spec}

	return s
}

func ParseInterface(spec *ast.TypeSpec, parent *ast.InterfaceType) *Interface {
	i := &Interface{parent, spec}

	return i
}