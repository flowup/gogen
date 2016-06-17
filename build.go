package gogen

import "go/ast"

type File struct {
	parent *ast.File
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