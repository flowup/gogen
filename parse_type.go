package gogen

import "go/ast"

// Structure represents the struct type of a
// given build
type Structure struct {

}

// Interface represents the interface type of a
// given build
type Interface struct {

}

func ParseStruct(tree *ast.StructType) *Structure {
	s := &Structure{}

	return s
}

func ParseInterface(tree *ast.InterfaceType) *Interface {
	i := &Interface{}

	return i
}