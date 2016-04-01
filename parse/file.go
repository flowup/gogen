package parse

import (
	"github.com/flowup/gogen/architect"
	"go/token"
	"go/parser"
)

// Any will determine if file, folder or any
// other resource is specifiec by the path
// and will try to parse it
func Any(path string) *architect.Build {
	build := architect.NewBuild()

	return build
}

// File will parse file at the given destination
func File(path string) (*architect.Build, error) {
	build := architect.NewBuild()

	var fset token.FileSet

	// parse the ast from the given file
	ast, err := parser.ParseFile(&fset, path, nil, parser.AllErrors)
	if err != nil {
		return nil, err
	}

	// make build from the parsed file
	build.Make(ast)

	return build, nil
}

// Folder will parse the whole folder on the
// given destination
func Folder(path string) *architect.Build {
	build := architect.NewBuild()

	return build
}
