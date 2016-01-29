package gogen

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GeneratorSuite struct {
	suite.Suite
}

// set of file paths that should output the
// package on the left
var testPackages = map[string]string{
	"abcd":                    "abcd",
	"/something":              "something",
	"./something":             "something",
	"/something/from/nothing": "nothing",
	"../aaa/bbb":              "bbb",
	"":                        "gogen",
}

// testGopaths key must be concatenated with
// the current gopath to give the result
var testGopaths = map[string]string{
	"resource": "github.com/flowup/gogen/resource",
	"abcd/ef":  "github.com/flowup/gogen/abcd/ef",
}

func (s *GeneratorSuite) TestSavePlateFullFileName() {
	p := SavePlate{
		FileName:  "foo",
		Extension: ".bar",
	}

	assert.Equal(s.T(), "foo.bar", p.FullFileName())
}

type testGenerator struct {
	Generator
}

func (g *testGenerator) Name() string {
	return "TestGenerator"
}

func (s *GeneratorSuite) TestInitialize() {
	// prepare resource container
	rc := &ResourceContainer{}
	rc.Add(5)
	rc.Add("hello")
	// prepare test generator
	g := testGenerator{}
	g.Initialize(rc)

	assert.Equal(s.T(), 2, len(*g.Resources))
}

func (s *GeneratorSuite) TestSetOutputDir() {
	g := testGenerator{}
	assert.Equal(s.T(), "", g.OutputDir)
	g.SetOutputDir("./output/dir")
	assert.Equal(s.T(), "output/dir", g.OutputDir)
}

func (s *GeneratorSuite) TestName() {
	g := testGenerator{}
	assert.Equal(s.T(), "TestGenerator", g.Name())

	base := Generator{}
	assert.Equal(s.T(), "Generator", base.Name())
}

func (s *GeneratorSuite) TestPackageName() {
	for dir, output := range testPackages {
		generator := Generator{
			OutputDir: dir,
		}
		assert.Equal(s.T(), output, generator.PackageName())
	}
}

func (s *GeneratorSuite) TestImportPath() {
	for relative, relativeToGopath := range testGopaths {
		generator := Generator{
			OutputDir: relative,
		}

		assert.Equal(s.T(), relativeToGopath, generator.ImportPath())
	}
}

func (s *GeneratorSuite) TestPrepare() {
	// test generated directory
	g := testGenerator{}
	g.SetOutputDir("_test/preparetest")
	err := g.Prepare()
	assert.Equal(s.T(), nil, err)
	// clear
	os.RemoveAll("_test/")

	// should not create anything
	g.SetOutputDir("")
	err = g.Prepare()
	assert.Equal(s.T(), nil, err)
}

func (s *GeneratorSuite) TestExecuteTemplate() {
	g := testGenerator{}
	g.Initialize(nil)
	g.ExecuteTemplate("testplate", "Something", struct{}{})

	tmpl, ok := g.Templates["testplate"]
	assert.Equal(s.T(), true, ok)
	assert.Equal(s.T(), "testplate", tmpl.FileName)
	// ExecuteTemplate should by default give .go extension
	assert.Equal(s.T(), ".go", tmpl.Extension)
	assert.Equal(s.T(), "", tmpl.OutputDir)
}

func TestGogenSuite(t *testing.T) {
	suite.Run(t, &GeneratorSuite{})
}
