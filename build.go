package gogen

import "go/ast"

// File encapsulates the physical File and it's ast
type File struct {
	name string // name of the file
	parent *ast.File
	// types
	structures map[string]*Structure
	interfaces map[string]*Interface
	functions map[string]*Function
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

// FilteredStructs is a type of map that can be filtered
// by its tags.
type FilteredStructs map[string]*Structure

// Filter can be called on a map of structures
// It will filter only those that have tag with name
// given by parameter
func (f FilteredStructs) Filter(name string) map[string]*Structure {
  newMap := make(map[string]*Structure)
  for it := range f {
    if f[it].Tags().Has(name) {
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

type FilteredFunctions map[string]*Function

func (f FilteredFunctions) Filter (name string) map[string]*Function {
  newMap := make(map[string]*Function)
  for it := range f {
    if f[it].Tags().Has(name) {
      newMap[it] = f[it]
    }
  }

  return newMap
}

// Functions will return a full map of functions provided
// by the file
func (f *File) Functions() map[string]*Function {
	return f.functions
}

// Build is an entity that holds the file map. It can be
// dynamically adjusted (files can be added or removed).
type Build struct {
	Files map[string]*File;
}

// NewBuild creates a new Build object that encapsulates
// files being processed in the build.
func NewBuild() *Build {
	return &Build{
		Files: make(map[string]*File),
	}
}

// AddFile adds a file with the given name and File reference
// into an existing map of files.
func (b *Build) AddFile(name string, f *File) {
	b.Files[name] = f
}