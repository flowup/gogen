package gogen

import (
	"go/parser"
	"go/token"
	"go/ast"
	"path/filepath"
)

// ParseDir will create a Build from the directory that
// was passed into the function.
func ParseDir(path string) (*Build, error) {
	var fileSet token.FileSet

	packages, err := parser.ParseDir(&fileSet, path, nil, parser.AllErrors)
	if err != nil {
		return nil, err
	}

	// create new build for the file set
	build := NewBuild()

	// iterate over all packages in the directory
	for _, pkg := range packages {
		// iterate over all files within the package
		for name, ast := range pkg.Files {
			fileAST, err := ParseFileAST(name, ast)
			if err != nil {
				return nil, err
			}
			build.AddFile(name, fileAST)
		}
	}

	return build, nil
}

// ParseFile will create a Build from the file path that
// was passed. FileSet of the Build will only contain a
// single file.
func ParseFile(path string) (*Build, error) {
	var fileSet token.FileSet

	ast, err := parser.ParseFile(&fileSet, path, nil, parser.AllErrors)
	if err != nil {
		return nil, err
	}

	fileName := filepath.Base(path)

	// create new build for the file
	build := NewBuild()
	fileAST, err := ParseFileAST(fileName, ast)
	if err != nil {
		return nil, err
	}

	// add parsed file to the build file set
	build.AddFile(fileName, fileAST)

	return build, nil
}

// ParseFileAST creates a File parse with all necessary
// structures.
func ParseFileAST(name string, ast *ast.File) (*File, error) {
	f := NewFile(name, ast)

	return f, nil
}