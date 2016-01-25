package main

import (
	"github.com/flowup/gogen"
	"github.com/flowup/gogen/generator-model"
	"github.com/flowup/gogen/resource"
)

func main() {
	// Define resources that should be available to
	// generators.
	gogen.Define(resource.User)

	// set output for the generator
	model.Generator.SetOutputDir("./model")

	// Define pipes that should be run
	gogen.Pipe(
		model.Generator,
	)

	// start the generator
	if err := gogen.Generate(); err != nil {
		panic(err)
	}
}
