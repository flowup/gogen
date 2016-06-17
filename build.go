package gogen

import "go/ast"

// File encapsulates the physical File and it's ast
type File struct {
	name string // name of the file
	parent *ast.File
}

// NewFile creates a new File instance with provided
// file name and it's parent ast.File.
func NewFile(name string, parent *ast.File) *File {
	return &File{name, parent}
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