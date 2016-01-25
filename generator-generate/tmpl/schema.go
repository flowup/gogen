package tmpl

import "text/template"

var (
	// SchemaTemplate is template for the schema file
	SchemaTemplate, _ = template.New("schema").Parse(`
  package {{.PackageName}}

	// Schema is main resource for the generator
  type Schema struct {

  }
  `)
)
