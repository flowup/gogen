package gogen

import (
	"go/parser"
	"go/token"
	"go/ast"
	"path/filepath"
	"fmt"
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
			baseName := filepath.Base(name)

			fileAST, err := ParseFileAST(baseName, ast)
			if err != nil {
				return nil, err
			}
			build.AddFile(baseName, fileAST)
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
func ParseFileAST(name string, tree *ast.File) (*File, error) {
	f := NewFile(name, tree)

	for _, declaration := range tree.Decls {
		switch decValue := declaration.(type) {
		// catch only generic declarations
		case *ast.GenDecl:
			for _, spec := range decValue.Specs {
				switch specValue := spec.(type) {
				case *ast.TypeSpec:
					// all cases should pass in also specValue as
					// it is the underlying spec
					switch typeValue := specValue.Type.(type) {
					case *ast.StructType:
						f.AddStruct(ParseStruct(specValue, typeValue))
					case *ast.InterfaceType:
						f.AddInterface(ParseInterface(specValue, typeValue))
					}
				case *ast.ImportSpec:
				case *ast.ValueSpec:
				default:
					fmt.Println("Generic value not recognized: ", specValue)
				}
			}
		// catch function declarations
		case *ast.FuncDecl:
			fun := ParseFunction(decValue)
			if !fun.IsMethod() {
				// add the function to the top level map
				f.AddFunction(fun)
			} else {
				// add the function to the structure it belongs to
			}
		}
	}

	return f, nil
}