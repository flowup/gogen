package gogen

import "go/ast"

// StructField encapsulates one field of the Structure
type StructField struct {
	parent *ast.Field

	name string
}

// Name returns the name of the field
func (f *StructField) Name() string {
	return f.name
}

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

// Fields returns fields that are associated with the
// given Structure. Note that this function builds
// the field list every time it is called, so cache
// the results to improve the performance.
func (s *Structure) Fields() []*StructField {
	fields := []*StructField{}

	for _, field := range s.parent.Fields.List {
		for _, fieldName := range field.Names {
			fields = append(fields, &StructField{
				field, fieldName.Name,
			})
		}
	}

	return fields
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