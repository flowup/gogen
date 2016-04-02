package architect

import "go/ast"

type Function interface {
	Name() string
	Exported() bool
}

// Function is an entity that provides basic information
// about the functions that are stored inside the first
// layer or struct methods
type FunctionImpl struct {
	name        string `json:"name"`
	exported    bool `json:"exported"`

	// back-links (ignored by the json)
	astFunction *ast.FuncDecl `json:"-"`
}

// NewFunction returns new instance of the FunctionImpl
func NewFunction() *FunctionImpl {
	return new(FunctionImpl)
}

// Name returns the name field of the function. This is
// the identifier name of the function, or empty in case
// of anonymous function
func (f *FunctionImpl) Name() string {
	return f.name
}

// Exported returns true if the function is exported, or
// false in case it is not exported ouside the package
func (f *FunctionImpl) Exported() bool {
	return f.exported
}