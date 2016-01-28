package gogen

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"
)

var (
	// ErrTemplateAlreadyExists error is returned every time
	// duplicate template is found when saving template
	ErrTemplateAlreadyExists = errors.New("Template could not be saved due to duplicate name")
)

// Generable interface provides interface definition
// for any generator that can be used within gogen
type Generable interface {
	// Initialize is called just before generate to
	// pass the generator resources that should be used
	Initialize(resources *ResourceContainer)
	// Generate is entry point to the generator. This
	// method is called only once, when the generator
	// is invoked
	Generate() error
	// Name returns the name of the generator. This
	// method should only be used for the debugging
	// purposes, as there may be duplicit community
	// generators with the same name.
	Name() string
	// SetOutputDir will change the output directory
	// from the default to the given. This is strongly
	// recommended every time user something wants to
	// generate.
	SetOutputDir(dir string)
	// Output will make generator output all files that
	// were temporarily created during its run. Note that
	// generator may output any files during it's generation,
	// this should allow generator to check for dependencies
	// or bad files if necessary.
	Output() error
}

// SavePlate is structure that encapsulates template
// that is ready to save. Valid SavePlate consists
// of output path, name of the template
// and its buffer that will be used every time execute is
// issued more times
//
// OutputDir is optional for the template, This means,
// it may not always be populated
type SavePlate struct {
	Content   *bytes.Buffer
	OutputDir string
	FileName  string
}

// Generator is base class that should be used
// as a composite to any other created generator context.
// It supports basic data flow and provides helpers.
//
// This type should encapsulate all widely used methods
// that are needed by the generators, thus may be extended
// by the time.
type Generator struct {
	// directory to which should all outputs go
	OutputDir string
	// Resources stores all Resources that were passed
	// to the current generator context
	Resources *ResourceContainer
	// Templates is map of maps of templates, where first
	// key of the map is name of the template
	Templates map[string]SavePlate
}

// Initialize accepts resources that should be used by
// the current generator context
func (g *Generator) Initialize(resources *ResourceContainer) {
	g.Resources = resources
	g.Templates = make(map[string]SavePlate)
}

// SetOutputDir will set the output dir of the generator
// to the specified value, which should result in code
// generated to the destination
func (g *Generator) SetOutputDir(dir string) {
	g.OutputDir = dir
}

// Name is virtual method that should return the
// name of the generator. This is used for the debugging
// purpose
func (g *Generator) Name() string {
	return "Generator"
}

// PackageName returns the name of the package based on the
// last directory from the OutputDir
func (g *Generator) PackageName() string {
	// get package chain from the output dir
	packChain := strings.Split(g.OutputDir, "/")

	if len(packChain) == 0 || packChain[len(packChain)-1] == "." {
		// return current working dir
		wd, _ := os.Getwd()
		pack := strings.Split(wd, "/")
		return pack[len(pack)-1]
	}

	// get the package (last in the chain)
	return packChain[len(packChain)-1]
}

// ImportPath returns import path to the current output dir.
func (g *Generator) ImportPath() string {
	cwd, _ := os.Getwd()
	outPath := path.Join(cwd, g.OutputDir)
	gopath := path.Join(os.Getenv("GOPATH"), "src")

	// remove gopath from the output directory
	return strings.TrimLeft(strings.TrimPrefix(outPath, gopath), "/")
}

// Prepare will ensure, that output directory exists
// and all needed values are correctly set
func (g *Generator) Prepare() error {
	var err error

	// if no output dir was
	if g.OutputDir == "" {
		g.SetOutputDir(".")
	}

	// create directories that are needed
	err = os.MkdirAll(g.OutputDir, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// ExecuteTemplate is same as ExecuteTemplateExt, but
// this function is using .go as a default extension
func (g *Generator) ExecuteTemplate(name, content string, schema interface{}) error {
	return g.ExecuteTemplateExt(name, content, schema, ".go")
}

// ExecuteTemplateExt saves template that will be generated once
// generator finishes run
func (g *Generator) ExecuteTemplateExt(name, content string, schema interface{}, ext string) error {
	if _, ok := g.Templates[name]; !ok {
		g.Templates[name] = SavePlate{
			Content:  new(bytes.Buffer),
			FileName: name + ext,
		}
	}

	tmpl, err := template.New(name).Parse(content)
	if err != nil {
		return err
	}

	err = tmpl.Execute(g.Templates[name].Content, schema)
	if err != nil {
		return err
	}

	return nil
}

// SaveFile will save provided content into the
// specified file with extension .gen.go and output
// directory previously set
func (g *Generator) SaveFile(name string, content bytes.Buffer) error {
	// calculate path to the file
	filePath := path.Join(g.OutputDir, name+".go")
	// save file
	return ioutil.WriteFile(filePath, content.Bytes(), os.ModePerm)
}

// Output will generate all accumulated instances
// of SavePlate in the generator. This allows the logic
// of the generator to execute first and possibly depend
// on templates that will be generated
func (g *Generator) Output() error {
	for name, plate := range g.Templates {
		genlog.Info("Generating template [%s], file [%s]", name, plate.FileName)
		err := g.SaveFile(path.Join(plate.OutputDir, plate.FileName), *plate.Content)
		if err != nil {
			return err
		}
	}

	return nil
}
