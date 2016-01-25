package main

import (
	"github.com/flowup/gogen"
	"github.com/flowup/gogen/generator-generate"
)

func main() {
	// Define pipes that should be run
	gogen.Pipe(
		generate.Generator,
	)

	// start the generator
	gogen.Generate()
}
