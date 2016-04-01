package architect

import (
	"go/ast"
)

// StructField defines a field that is present
// within the structure in the build
type StructField struct {
	Name string
}

// Function is an entity that provides basic information
// about the functions that are stored inside the first
// layer or struct methods
type Function struct {
	Name string
}

// Struct is declaration of any structure type
// defined within the Build
type Struct struct {
	Name    string
	Fields  []*StructField
	Methods []*Function
}

// Build stores symbols that are available
// in the given package or file.
type Build struct {
	pack    string // name of the package
	imports []string // list of import packages
	structs []*Struct
}

// NewBuild will return new Build
func NewBuild() *Build {
	return &Build{
		pack: "",
		imports: []string{},
		structs: []*Struct{},
	}
}

// Package returns name of the package which
// is built in the Build instance
func (b *Build) Package() string {
	return b.pack
}

// HasImport will check if given import is included
// in the build requirement
func (b *Build) HasImport(check string) bool {
	// correct import if no " is present
	if check[0] != '"' {
		check = "\"" + check + "\""
	}

	for _, imp := range b.imports {
		if imp == check {
			return true
		}
	}

	return false
}

// Make will start the build from the given
// ast file
func (b *Build) Make(tree *ast.File) {
	b.pack = tree.Name.Name

	// parse import paths requested by the build
	b.parseImports(tree.Imports)

	// iterate over all declarations in the ast
	for _, decl := range tree.Decls {
		switch declValue := decl.(type) {
		// catch only generic declarations
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

func (b *Build) parseImports(imports []*ast.ImportSpec) {
	// iterate over imports and save paths
	for _, imp := range imports {
		b.imports = append(b.imports, imp.Path.Value)
	}
}

func (b *Build) makeTypeSpec(spec *ast.TypeSpec) {
	stru := &Struct{
		Name: spec.Name.Name,
	}

	//switch specValue := spec.Type.(type) {
	//case *ast.StructType:
	// iterate over fields
	//}

	b.structs = append(b.structs, stru)
}

func (b *Build) makeImportSpec(spec *ast.ImportSpec) {
}

func (b *Build) makeValueSpec(spec *ast.ValueSpec) {

}