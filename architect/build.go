package architect

import (
	"go/ast"
	"fmt"
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
func (b *Build) Make(tree *ast.File) {
	b.pack = tree.Name.Name

	for _, decl := range tree.Decls {
		switch declValue := decl.(type) {
		case *ast.GenDecl:
			for _, spec := range declValue.Specs {
				switch specValue := spec.(type) {
				case *ast.TypeSpec:
					b.makeTypeSpec(specValue)
				case *ast.ImportSpec:
					b.makeImportSpec(specValue)
				case *ast.ValueSpec:
					b.makeValueSpec(specValue)
				}
			}
		}
	}
}

func (b *Build) makeTypeSpec(spec *ast.TypeSpec) {

}

func (b *Build) makeImportSpec(spec *ast.ImportSpec) {

}

func (b *Build) makeValueSpec(spec *ast.ValueSpec) {

}