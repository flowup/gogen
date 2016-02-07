package unk

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// Possible sources that can be recognized
const (
	None       = iota
	FileSource // local file
	HTTPSource // http link
	GistSource // gist link
)

// Source stores data about any object that
// is importable (or unimportable).
type Source struct {
	// Path to source. Filepath, url, etc.
	Path string
	// Name of the source. File name, folder name etc.
	Name string
	// Type from the list of possible sources
	Type int
	// content is cached content. If Fetch() is called
	// second time, it will directly return this conent
	content string
}

// NewSource returns a new instance of the source
func NewSource(path string) *Source {
	return &Source{
		Path: path,
	}
}

// Resolve will cause the source to get it's name
// and type. These are automatically determined
func (s *Source) Resolve() error {
	// default type is FileSource
	s.Type = FileSource

	// try to determine other type
	// TODO

	switch s.Type {
	case FileSource:
		if strings.Index(s.Path, "/") == -1 {
			// name is the path itself
			s.Name = s.Path
		} else {
			chain := strings.Split(s.Path, "/")
			s.Name = chain[len(chain)-1]
		}
	}
	return nil
}

// Content will return content from the source.
// See the type definitions for what will be
// returned by the Content
func (s *Source) Content() (string, error) {
	if s.content != "" {
		// ReFetch will also store content
		return s.content, nil
	}

	return s.Fetch()
}

// Fetch will get content of the source. Even
// if the content is already loaded, this will
// cause it to reload. Any changes made to the
// configuration of the Source will appear
func (s *Source) Fetch() (string, error) {
	// error that will be returned at the end
	var err error
	// bytes that will be stored to the content
	var bytes []byte

	switch s.Type {
	// For the FIleSource, we just reading the file
	// and returning its content
	case FileSource:
		bytes, err = ioutil.ReadFile(s.Path)

	case HTTPSource:
		var resp *http.Response
		resp, err = http.Get(s.Path)
		if err != nil {
			break
		}
		defer resp.Body.Close()
		bytes, err = ioutil.ReadAll(resp.Body)

	case GistSource:
		panic("Gist source not implemented")
	}

	// store bytes
	s.content = string(bytes)

	return s.content, err
}
