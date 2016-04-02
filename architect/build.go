package architect

import (
	"go/ast"
	"unicode"
	"fmt"
)

// StructField defines a field that is present
// within the structure in the build
type StructField struct {
	Name string
}

type Function interface {
	Name() string
	Exported() bool
}

// Function is an entity that provides basic information
// about the functions that are stored inside the first
// layer or struct methods
type FunctionImpl struct {
	name     string
	exported bool
}

func (f *FunctionImpl) Name() string {
	return f.name
}

func (f *FunctionImpl) Exported() bool {
	return f.exported
}

type Struct interface {
	FindMethod(name string) Function
}

// Struct is declaration of any structure type
// defined within the Build
type StructImpl struct {
	Name      string
	Fields    []*StructField
	Methods   []Function

	// Back-links to the AST
	astStruct *ast.StructType
}

// FindMethod searches for the method by the name in
// the current structure instance
func (s *StructImpl) FindMethod(name string) Function {
	for _, m := range s.Methods {
		if m.Name() == name {
			return m
		}
	}

	return nil
}

// Build stores symbols that are available
// in the given package or file.
type Build struct {
	pack      string          // name of the package
	imports   []string        // list of import packages
	structs   []*StructImpl   // list of structures in the build
	functions []*FunctionImpl // list of functions in the build
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

// GetFunction searches only first layer functions of the
// build. This does not include any methods or anonymous
// functions inside other functions
func (b *Build) FindFunction(name string) *FunctionImpl {
	for _, fn := range b.functions {
		if fn.name == name {
			return fn
		}
	}

	return nil
}

// FindStruct searches for the given structure name in the
// build and returns the first match
func (b *Build) FindStruct(name string) *StructImpl {
	for _, st := range b.structs {
		if st.Name == name {
			return st
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
			b.parseFunction(declValue)
		}
	}
}

func (b *Build) parseImports(imports []*ast.ImportSpec) {
	// iterate over imports and save paths
	for _, imp := range imports {
		b.imports = append(b.imports, imp.Path.Value)
	}
}

// parseFunction will parse the given function declaration.
// If receiver parameter is nil, it is a function, else,
// it belongs to the field set, which is to be found.
func (b *Build) parseFunction(f *ast.FuncDecl) {

	fun := &FunctionImpl{
		name: f.Name.Name,
		// exported functions have first letter upper cased
		exported: unicode.IsUpper(rune(f.Name.Name[0])),
	}

	if f.Recv == nil {
		// add function to the build
		b.functions = append(b.functions, fun)
	} else {
		// find structure to which the field-set belongs
		// and add method to the structure methods
		found := false
		for _, s := range b.structs {
			if s.astStruct == f.Recv.List[0].Type.(*ast.StarExpr).X.(*ast.Ident).Obj.Decl.(*ast.TypeSpec).Type.(*ast.StructType) {
				s.Methods = append(s.Methods, fun)
				found = true
				break
			}
		}

		if !found {
			fmt.Println("Struct not found for the method:", fun.name)
		}
	}
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
	stru := &StructImpl{
		Name: name,
		// back-links
		astStruct: st,
	}

	// iterate over fields
	for _, field := range st.Fields.List {
		f := &StructField{
			Name: field.Names[0].Name,
		}

		stru.Fields = append(stru.Fields, f)
	}

	// add to the list of the structures
	b.structs = append(b.structs, stru)
}

func (b *Build) handleImportSpec(spec *ast.ImportSpec) {
}

func (b *Build) handleValueSpec(spec *ast.ValueSpec) {

}