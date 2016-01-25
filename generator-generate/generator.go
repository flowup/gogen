package generate

import (
	"bytes"

	"github.com/flowup/gogen"
	"github.com/flowup/gogen/generator-generate/tmpl"
	"github.com/op/go-logging"
)

var (
	// Generator is instance of current generator
	Generator = &generator{}

	genlog = logging.MustGetLogger("gogen")
)

type generator struct {
	gogen.Generator
}

func (g *generator) Generate() error {
	genlog.Info("Generating generator")

	// execute generator template
	gencontent := bytes.Buffer{}
	err := tmpl.GeneratorTemplate.Execute(&gencontent, g)
	if err != nil {
		return err
	}
	g.SaveFile("generator", gencontent)

	// execute schema
	schemacontent := bytes.Buffer{}
	err = tmpl.SchemaTemplate.Execute(&schemacontent, g)
	if err != nil {
		return err
	}
	g.SaveFile("schema", schemacontent)

	return nil
}
