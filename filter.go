package gogen

import "regexp"

// Source filter types
const (
	NameFitler = iota
	ContentFilter
)

// Filter holds data about resource filtering of gogen.
// It is used by the Importers subsystem to match the
// remote and local unknown resources.
// Filter must have known Type that should be parsed.
// While some rousrces can be determined by the name of
// the file or url (e.g. source files like Go), others are
// only determinable by their content (e.g. JSON)
type Filter struct {
	Type int
	Exp  *regexp.Regexp
}

var (
	// sourceFilters is a map of source filter arrays.
	// This provides a container of all available filters
	// and their types to the filtering system.
	sourceFilters = map[string][]Filter{}
)

// RegisterFilter adds filters to the existing set of filters.
// RegisterFilter accepts more more filters, this means there
// can be more filters by the filter name. When registering
// filters to the already registered name, these filters will
// be appended to the existing set, rather than rewritten
func RegisterFilter(name string, expType int, expressions ...string) {
	for _, exp := range expressions {
		sourceFilters[name] = append(sourceFilters[name], Filter{
			Type: expType,
			Exp:  regexp.MustCompile(exp),
		})
	}
}

// GetFilters returns filters stored in the map of filters
// by the given name.
func GetFilters(name string) []Filter {
	return sourceFilters[name]
}
