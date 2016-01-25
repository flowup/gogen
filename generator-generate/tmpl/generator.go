package tmpl

import "text/template"

var (
	// GeneratorTemplate is template for generic generator
	GeneratorTemplate, _ = template.New("generator template").Parse(`
  package {{.PackageName}}

  var (
		// Generator is instance of current generator
    Generator = &generator{}

    genlog = logging.MustGetLogger("gogen")
  )

	type generator struct {
		gogen.Generator
	}

	func (g *Generator) Generate() error {
		// buffer for output
		content := bytes.Buffer{}

		return nil
	}
  `)
)
