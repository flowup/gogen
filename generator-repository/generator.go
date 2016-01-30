package repository

import (
	"github.com/flowup/gogen"
	"github.com/flowup/gogen/generator-model"
	"github.com/flowup/gogen/generator-repository/tmpl"
)

var (
	// Generator is global registration of the generator
	Generator = &generator{}
)

// Repository types that can be used
const (
	Mongo = iota
	Postgres
	Redis
)

// generator encapsulates the logic behind
// generating of models
type generator struct {
	gogen.Generator

	repositoryType int
}

// Name returns name of the generator
func (g *generator) Name() string {
	return "RepositoryGenerator"
}

// SetRepositoryType will set the type of the generated
// repository. Defaults to Mongo
func (g *generator) SetRepositoryType(t int) {
	g.repositoryType = t
}

// Generate will call the generator to generate
// results
func (g *generator) Generate() error {
	err := g.Prepare()
	if err != nil {
		return err
	}

	// temporary template variable
	var templ string
	switch g.repositoryType {
	case Mongo:
		templ = repositorytmpl.MongoRepositoryTemplate
	}

	// search for schemas in the resources
	schemas := g.Resources.Search(&model.Schema{})

	for _, schema := range schemas {
		entity := schema.(*model.Schema)
		g.Log.Info("Generating repository for model %s", entity.Name)
		err = g.ExecuteTemplate(entity.Name+"Repository", templ, struct {
			Model       *model.Schema
			PackageName string
		}{
			Model:       entity,
			PackageName: g.PackageName(),
		})
		if err != nil {
			return err
		}
	}

	return nil
}
