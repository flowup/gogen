package parse

import "github.com/flowup/gogen/architect"

// Any will determine if file, folder or any
// other resource is specifiec by the path
// and will try to parse it
func Any(path string) *architect.Build {
	build := architect.NewBuild()

	return build
}

// File will parse file at the given destination
func File(path string) *architect.Build {
	build := architect.NewBuild()

	return build
}

// Folder will parse the whole folder on the
// given destination
func Folder(path string) *architect.Build {
	build := architect.NewBuild()

	return build
}
