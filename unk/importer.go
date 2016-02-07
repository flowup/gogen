package unk

// ImporterCollection is a collection of importers that is used for
// the unknown source resolving.
type ImporterCollection struct {
	Importers []Importer
}

// Register will register importer for the use.
func (i *ImporterCollection) Register(imp Importer) {
	i.Importers = append(i.Importers, imp)
}

// Importer is an interface, that defines the behavior
// of resource importer. Importer should be able to
// convert specific type of loadable resource into the
// specific object or objects. All resources that are
// imported are defined and can be used by the generators
type Importer interface {
	Import(imps []Source) error
}
