package gogen

import "go/ast"

type File struct {
	name string // name of the file
	parent *ast.File
}

func NewFile(name string, parent *ast.File) *File {
	return &File{name, parent}
}

func (f *File) Name() string {
	return f.name
}

func (f *File) Package() string {
	return f.parent.Name.Name
}

type Build struct {
	Files map[string]*File;
}

func NewBuild() *Build {
	return &Build{
		Files: make(map[string]*File),
	}
}

func (b *Build) AddFile(name string, f *File) {
	b.Files[name] = f
}