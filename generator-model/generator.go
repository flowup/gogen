package model

import (
	"bytes"
	"text/template"

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

	// compile model template
	tmpl, err := template.New("model").Parse(modelTemplate)
	if err != nil {
		return err
	}

	for _, resource := range *g.Resources {
		if model, ok := resource.(*Schema); ok {
			genlog.Info("Generating model for %s", model.Name)
			content := bytes.Buffer{}
			tmpl.Execute(&content, struct {
				Model       *Schema
				PackageName string
			}{
				Model:       model,
				PackageName: g.PackageName(),
			})
			g.SaveFile(model.Name, content)

			// set other meta to model
			model.Package = g.PackageName()
		}
	}

	return nil
}

// Templates
var (
	modelTemplate = `
		package {{.PackageName}}

		// {{.Model.Name}} is model representing the entity
		type {{.Model.Name}} struct {
		  {{range .Fields}}{{.Model.Name}} {{.Type.Name}}
		  {{end}}
		}`
)
