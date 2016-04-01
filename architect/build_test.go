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
	s.build = &Build{}
	s.build.Make(ast)
}

func (s *BuildSuite) TestPackageName() {
	assert.Equal(s.T(), "architect", s.build.Package())
}

func TestBuildSuite(t *testing.T) {
	suite.Run(t, &BuildSuite{})
}