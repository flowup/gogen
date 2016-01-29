package model

import (
	"github.com/flowup/gogen"
	"github.com/op/go-logging"
)

var (
	// Generator is global registration of the generator
	Generator = &generator{}

	genlog = logging.MustGetLogger("gogen")
)

// Types of models
const (
	Default = iota
	Mongo
)

// ModelGenerator encapsulates the logic behind
// generating of models
type generator struct {
	gogen.Generator

	modelType int
}

// Name returns name of the generator
func (g *generator) Name() string {
	return "ModelGenerator"
}

// Generate will call the generator to generate
// results
func (g *generator) Generate() error {
	err := g.Prepare()
	if err != nil {
		return err
	}

	schemas := g.Resources.Search(&Schema{})
	if len(schemas) == 0 {
		genlog.Warning("No Schemas for the generator ModelGenerator found")
	}

	for _, ischema := range schemas {
		schema := ischema.(*Schema)
		// meta data to the schema
		schema.Package = g.PackageName()

		g.ExecuteTemplate(schema.Name+"Model", modelTemplate, struct {
			Model       *Schema
			PackageName string
		}{
			Model:       schema,
			PackageName: g.PackageName(),
		})
	}

	return nil
}

// Templates
var (
	modelTemplate = `
		package {{.PackageName}}

		// {{.Model.Name}} is model representing the entity
		type {{.Model.Name}} struct {
		  {{range .Model.Fields}}{{.Name}} {{.Type.Name}}
		  {{end}}
		}`
)
