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
		for name, astTree := range pkg.Files {
			baseName := filepath.Base(name)

			// create a comment map from file
			commentMap := ast.NewCommentMap(&fileSet, astTree, astTree.Comments)

			fileAST, err := ParseFileAST(baseName, astTree, commentMap)
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

	astTree, err := parser.ParseFile(&fileSet, path, nil, parser.AllErrors|parser.ParseComments)
	if err != nil {
		return nil, err
	}

	fileName := filepath.Base(path)

	// create a comment map from file
	commentMap := ast.NewCommentMap(&fileSet, astTree, astTree.Comments)

	// create new build for the file
	build := NewBuild()
	fileAST, err := ParseFileAST(fileName, astTree, commentMap)
	if err != nil {
		return nil, err
	}

	// add parsed file to the build file set
	build.AddFile(fileName, fileAST)

	return build, nil
}

// ParseFileAST creates a File parse with all necessary
// structures.
func ParseFileAST(name string, tree *ast.File, commentMap ast.CommentMap) (*File, error) {
	f := NewFile(name, tree)

	for _, i := range tree.Imports {
		f.AddImport(ParseImport(i, commentMap.Filter(i)))
	}

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
						f.AddStruct(ParseStruct(specValue, typeValue, commentMap.Filter(declaration)))
					case *ast.InterfaceType:
						f.AddInterface(ParseInterface(specValue, typeValue, commentMap.Filter(declaration)))
					case *ast.FuncType:
						fmt.Println("Generic value not recognized: ", specValue)
					case *ast.ArrayType:
						f.AddArray(ParseArray(specValue, typeValue, commentMap.Filter(declaration)))
					case *ast.MapType:
						fmt.Println("Generic value not recognized: ", specValue)
					case *ast.ChanType:
						fmt.Println("Generic value not recognized: ", specValue)
					default:
						f.AddBaseType(ParseBaseType(specValue, typeValue, commentMap.Filter(declaration)))
					}

				case *ast.ImportSpec:
					// just ignore for now
				case *ast.ValueSpec:
					f.AddConstant(ParseConstant(specValue, commentMap.Filter(declaration)))
				default:
					fmt.Println("Generic value not recognized: ", specValue)
				}
			}
			// catch function declarations
		case *ast.FuncDecl:
			fun := ParseFunction(decValue, commentMap.Filter(declaration))
			if !fun.IsMethod() {
				// add the function to the top level map
				f.AddFunction(fun)
			} else {
				// add the function to the structure it belongs to
				if len(fun.parent.Recv.List) <= 0 {
					// TODO: no receiver defined report?
					break
				}

				// struct that should be assigned the method
				var structType *ast.StructType

				switch receiver := fun.parent.Recv.List[0].Type.(type) {
				// pointer receiver
				case *ast.StarExpr:
					// if the receiver is defined append it to it,
					// otherwise register it as normal function
					if receiver.X.(*ast.Ident).Obj != nil {
						structType = receiver.X.(*ast.Ident).Obj.Decl.(*ast.TypeSpec).Type.(*ast.StructType)
					} else {
						f.AddFunction(fun)
					}
				// copy receiver
				case *ast.Ident:
					switch receiver.Obj.Decl.(*ast.TypeSpec).Type.(type) {
					case *ast.StructType:
						structType = receiver.Obj.Decl.(*ast.TypeSpec).Type.(*ast.StructType)
					}
				}

				// search for the structures that receive the method
				// and bind it
				for _, st := range f.structures {
					if st.parent == structType {
						st.AddMethod(fun)
					}
				}
			}
		}
	}

	return f, nil
}
