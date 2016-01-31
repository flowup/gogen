package gosource

import "github.com/flowup/gogen"

// Importer is able to import golang files and create
// structure, function and constant resources
type Importer struct {
}

func (i *Importer) Import(imps []gogen.Importable) error {
	return nil
}
