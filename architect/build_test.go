package architect

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"github.com/stretchr/testify/assert"
	"go/token"
	"go/parser"
)

type BuildSuite struct {
	suite.Suite

	build *Build
}

func (s *BuildSuite) SetupTest() {
	var fset token.FileSet

	// parse the ast from the given file
	ast, err := parser.ParseFile(&fset, "./fixture.go", nil, parser.AllErrors)
	assert.Equal(s.T(), nil, err)

	// create build and make it from the AST
	s.build = NewBuild()
	s.build.Make(ast)
}

func (s *BuildSuite) TestPackageName() {
	assert.Equal(s.T(), "architect", s.build.Package())
}

func (s *BuildSuite) TestStructDefinitions() {
	// check for struct length
	assert.Equal(s.T(), 2, len(s.build.structs))

	// check struct properties
	assert.Equal(s.T(), "X_Test", s.build.structs[0].Name)
	assert.Equal(s.T(), "Y_Test", s.build.structs[1].Name)
}

func (s *BuildSuite) TestImports() {
	assert.Equal(s.T(), 2, len(s.build.imports))

	assert.Equal(s.T(), true, s.build.HasImport("fmt"))
	assert.Equal(s.T(), true, s.build.HasImport("math"))
}

func TestBuildSuite(t *testing.T) {
	suite.Run(t, &BuildSuite{})
}