package gogen

import "go/ast"

// File encapsulates the physical File and it's ast
type File struct {
	name string // name of the file
	parent *ast.File

	// imports
	imports map[string]*Import

	// types
	structures map[string]*Structure
	interfaces map[string]*Interface
	functions map[string]*Function
	constants map[string]*Constant
}

// NewFile creates a new File instance with provided
// file name and it's parent ast.File.
func NewFile(name string, parent *ast.File) *File {
	return &File{
		name: name,
		parent: parent,
		structures: make(map[string]*Structure),
		interfaces: make(map[string]*Interface),
		functions: make(map[string]*Function),
    constants: make(map[string]*Constant),
		imports: make(map[string]*Import),
	}
}

// Name returns the base name of the file. This will
// not return the full path of the file
func (f *File) Name() string {
	return f.name
}

// Package returns name of the package that is being
// referenced by the file
func (f *File) Package() string {
	return f.parent.Name.Name
}

// AddImport adds passed import to imports in the file
func (f *File) AddImport(i *Import) {
	f.imports[i.Name()] = i
}

// Import will return an import with given name
// or nil if file does not contain such import
func (f *File) Import(name string) *Import{
	return f.imports[name]
}

func (f *File) Imports() map[string]*Import {
	return f.imports
}

// AddStruct adds passed structure type into the
// structures described by the File
func (f *File) AddStruct(s *Structure) {
	f.structures[s.Name()] = s
}

// Struct returns a structure with the given name.
// If the structure could not be found, returns nil
func (f *File) Struct(name string) *Structure {
	return f.structures[name]
}

// FilteredStructs is a type of map of structures
// that can be filtered by its annotations.
type FilteredStructs map[string]*Structure

// Filter can be called on a map of structures
// It will filter only those that have annotation with name
// given by parameter
func (f FilteredStructs) Filter(name string) map[string]*Structure {
  newMap := make(map[string]*Structure)
  for it := range f {
    if f[it].Annotations().Has(name) {
      newMap[it] = f[it]
    }
  }

  return newMap
}

// Structs returns a map of structures contained
// within the requested file
func (f *File) Structs() FilteredStructs {
	return f.structures
}

// AddInterface adds passed interface type into the
// interfaces described by the File
func (f *File) AddInterface(i *Interface) {
	f.interfaces[i.Name()] = i
}

// Interface returns an interface with the given name.
// If no interface with the name was found, returns nil
func (f *File) Interface(name string) *Interface {
	return f.interfaces[name]
}

// AddFunction will add the passed function into the
// map of functions described by the file
func (f *File) AddFunction(fun *Function) {
	f.functions[fun.Name()] = fun
}

// Function will return a function by it's name. If no
// function is found, this will return nil
func (f *File) Function(name string) *Function {
	return f.functions[name]
}

// AddConstant will add a constant to map of constants.
func (f *File) AddConstant(c *Constant) {
	f.constants[c.Name()] = c
}

// Constant will return a constant by it's name. If no
// constant is found, this will return nil
func (f *File) Constant(name string) *Constant{
	return f.constants[name]
}

// Constants will return full map of constants
func (f *File) Constants() FilteredConstants{
	return f.constants
}

// FilteredConstants is a type of map of constants
// that can be filtered by its annotations.
type FilteredConstants map[string]*Constant

// Filter will filter a map of constants by their annotations.
func (f FilteredConstants) Filter (name string) map[string]*Constant {
	newMap := make(map[string]*Constant)
	for it := range f {
		if f[it].Annotations().Has(name) {
			newMap[it] = f[it]
		}
	}

	return newMap
}

// FilteredFunctions is a type of map of functions
// that can be filtered by its annotations.
type FilteredFunctions map[string]*Function

// Filter will filter a map of functions by their annotations
func (f FilteredFunctions) Filter (name string) map[string]*Function {
  newMap := make(map[string]*Function)
  for it := range f {
    if f[it].Annotations().Has(name) {
      newMap[it] = f[it]
    }
  }

  return newMap
}

// Functions will return a full map of functions provided
// by the file
func (f *File) Functions() FilteredFunctions {
	return f.functions
}

// Build is an entity that holds the file map. It can be
// dynamically adjusted (files can be added or removed).
type Build struct {
	files map[string]*File;
}

// Files will return all files contained in a build
func (b *Build) Files() map[string]*File {
	return b.files
}

// File will return a file in a build by its name
func (b *Build) File(name string) *File{
	return b.files[name]
}

// NewBuild creates a new Build object that encapsulates
// files being processed in the build.
func NewBuild() *Build {
	return &Build{
		files: make(map[string]*File),
	}
}

// AddFile adds a file with the given name and File reference
// into an existing map of files.
func (b *Build) AddFile(name string, f *File) {
	b.files[name] = f
}