package architect

import (
	"go/ast"
	"encoding/json"
)

// Field is any field from the Struct field, to function
// parameter fields
type Field interface {
	Name() string
}

// FieldImpl defines an implementation of the Field interface
type FieldImpl struct {
	name string
}

// Name returns name of the current Field instance
func (f *FieldImpl) Name() string {
	return f.name
}

// Struct is a definition for the Struct entities
// present in the Build
type Struct interface {
	Name() string
	// getters
	Field(name string) Field

	AddMethod(method Function)
	Method(name string) Function
	// json marshaler
	// AST type getter
	AST() *ast.StructType
}

// StructImpl defined the implementation of the Struct
// interface, which can be used transparently with the
// architect package
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

// AddMethod adds given method into the existing set of
// methods
func (s *StructImpl) AddMethod(method Function) {
	s.methods = append(s.methods, method)
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

// MarshalJSON is a helper method for the json marshaller,
// which is needed as all fields are unexported
func (s *StructImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Name    string `json:"name"`
		Fields  []Field `json:"fields"`
		Methods []Function `json:"methods"`
	}{
		s.name, s.fields, s.methods,
	})
}

// AST returns the underlying ast node of the StructImpl
func (s *StructImpl) AST() *ast.StructType {
	return s.astStruct
}