package architect

import (
	"go/ast"
	"encoding/json"
)

type Field interface {
	Name() string
}

// StructField defines a field that is present
// within the structure in the build
type FieldImpl struct {
	name string `json:"name"`
}

// Name returns name of the current Field instance
func (f *FieldImpl) Name() string {
	return f.name
}

type Struct interface {
	Name() string
	// getters
	Field(name string) Field

	AddMethod(method Function)
	Method(name string) Function
	// json marshaler
	MarshalJSON() ([]byte, error)
	// AST type getter
	AST() *ast.StructType
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

func (s *StructImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Name    string `json:"name"`
		Fields  []Field `json:"fields"`
		Methods []Function `json:"methods"`
	}{
		s.name, s.fields, s.methods,
	})
}

func (s *StructImpl) AST() *ast.StructType {
	return s.astStruct
}