package main

import (
	"github.com/flowup/gogen"
	"github.com/flowup/gogen/generator-model"
	"github.com/flowup/gogen/generator-repository"
	"github.com/flowup/gogen/resource"
)

func main() {
	// Define resources that should be available to
	// generators.
	gogen.Define(resource.User)

	// set output for the generator
	model.Generator.SetOutputDir("./model")
	repository.Generator.SetOutputDir("./repository")

	// Define pipes that should be run
	gogen.Pipe(
		model.Generator,
		repository.Generator,
	)

	// start the generator
	if err := gogen.Generate(); err != nil {
		panic(err)
	}
}
