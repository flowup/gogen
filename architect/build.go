package architect

import (
	"go/ast"
	"unicode"
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
	Name     string
	Exported bool
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
	pack      string      // name of the package
	imports   []string    // list of import packages
	structs   []*Struct   // list of structures in the build
	functions []*Function // list of functions in the build
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

// GetFunctions searches only first layer functions of the
// build. This does not include any methods or anonymous
// functions inside other functions
func (b *Build) GetFunction(name string) *Function {
	for _, fn := range b.functions {
		if fn.Name == name {
			return fn
		}
	}

	return nil
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
					b.handleTypeSpec(specValue)
				case *ast.ImportSpec:
					b.handleImportSpec(specValue)
				case *ast.ValueSpec:
					b.handleValueSpec(specValue)
				}
			}
		// catch function declarations
		case *ast.FuncDecl:
			// parse only functions that are not bound
			// to struct field list. Methods are parsed
			// within the struct parsing
			if declValue.Recv == nil {
				b.parseFunction(declValue)
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

func (b *Build) parseFunction(f *ast.FuncDecl) {
	fun := &Function{
		Name: f.Name.Name,
		// exported functions have first letter upper cased
		Exported: unicode.IsUpper(rune(f.Name.Name[0])),
	}

	// add function to the build
	b.functions = append(b.functions, fun)
}

func (b *Build) handleTypeSpec(spec *ast.TypeSpec) {

	// pre-save spec name - this may include validations
	// in the future
	specName := spec.Name.Name

	switch specValue := spec.Type.(type) {
	case *ast.StructType:
		//iterate over fields
		b.parseStructure(specName, specValue)
	}
}

func (b *Build) parseStructure(name string, st *ast.StructType) {
	stru := &Struct{
		Name: name,
	}

	// iterate over fields
	for _, field := range st.Fields.List {
		f := &StructField{
			Name: field.Names[0].Name,
		}

		stru.Fields = append(stru.Fields, f)
	}

	// add to the list of the structs
	b.structs = append(b.structs, stru)
}

func (b *Build) handleImportSpec(spec *ast.ImportSpec) {
}

func (b *Build) handleValueSpec(spec *ast.ValueSpec) {

}