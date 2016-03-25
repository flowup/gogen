package architect

import (
	"fmt"
	"go/ast"
)

// Build stores symbols that are available
// in the given package or file.
type Build struct {
	pack string // name of the package
}

// NewBuild will return new Build
func NewBuild() *Build {
	return &Build{}
}

// Package returns name of the package which
// is built in the Build instance
func (b *Build) Package() string {
	return b.pack
}

// Make will start the build from the given
// ast file
func (b *Build) Make(ast *ast.File) {
	fmt.Println(ast.Name.Name)
}
