package gogen

import (
	"io/ioutil"
	"net/http"
)

// Importable is a resource that should be imported. This
// can be anything that has type and link.
type Importable interface {
	// GetSource will return path to source. This can be http
	// local file or anything else.
	GetSource() string
}

// Importer is an interface, that defines the behavior
// of resource importer. Importer should be able to
// convert specific type of loadable resource into the
// specific object or objects. All resources that are
// imported are defined and can be used by the generators
type Importer interface {
	Import(imps []Importable) error
}

// GetSourceFromUrl returns source from the given page in
// string data.
func GetSourceFromUrl(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}
