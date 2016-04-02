package architect

import (
	"go/ast"
)

type Field interface {
	Name() string
}

// StructField defines a field that is present
// within the structure in the build
type FieldImpl struct {
	name string
}

// Name returns name of the current Field instance
func (f *FieldImpl) Name() string {
	return f.name
}

type Struct interface {
	Name() string
	// getters
	Field(name string)
	Method(name string) Function
}

// Struct is declaration of any structure type
// defined within the Build
type StructImpl struct {
	name      string
	fields    []Field
	methods   []Function

	// Back-links to the AST
	astStruct *ast.StructType
}

// NewStruct returns new StructImpl object populated
// with the given name
func NewStruct(name string) *StructImpl {
	return &StructImpl{
		name: name,
	}
}

// Name returns name of the current struct instance
func (s *StructImpl) Name() string {
	return s.name
}

// Field searches for the field by the given name
func (s *StructImpl) Field(name string) Field {
	for _, f := range s.fields {
		if f.Name() == name {
			return f
		}
	}

	return nil
}

// Method searches for the method by the name in
// the current structure instance
func (s *StructImpl) Method(name string) Function {
	for _, m := range s.methods {
		if m.Name() == name {
			return m
		}
	}

	return nil
}