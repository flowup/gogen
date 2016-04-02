package architect

import (
	"go/ast"
	"encoding/json"
)

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
	// Struct extends the Interface interface
	Interface
	// getters
	Var(name string) Variable
	// json marshaler
	// AST type getter
	AST() *ast.StructType
}

// StructImpl defined the implementation of the Struct
// interface, which can be used transparently with the
// architect package
type StructImpl struct {
	// StructImpl extends the InterfaceImpl
	InterfaceImpl

	name      string
	fields    []Variable
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

// Var searches for the variable by the given name
func (s *StructImpl) Var(name string) Variable {
	for _, f := range s.fields {
		if f.Name() == name {
			return f
		}
	}

	return nil
}

// MarshalJSON is a helper method for the json marshaller,
// which is needed as all fields are unexported
func (s *StructImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Name    string `json:"name"`
		Fields  []Variable `json:"fields"`
		Methods []Function `json:"methods"`
	}{
		s.name, s.fields, s.methods,
	})
}

// AST returns the underlying ast node of the StructImpl
func (s *StructImpl) AST() *ast.StructType {
	return s.astStruct
}