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
	assert.Equal(s.T(), "X_Test", s.build.structs[0].name)
	assert.Equal(s.T(), "Y_Test", s.build.structs[1].name)
}

func (s *BuildSuite) TestImports() {
	assert.Equal(s.T(), 2, len(s.build.imports))

	assert.Equal(s.T(), true, s.build.HasImport("fmt"))
	assert.Equal(s.T(), true, s.build.HasImport("math"))
}

func (s *BuildSuite) TestFunctions() {
	assert.Equal(s.T(), 4, len(s.build.functions))

	fn := s.build.FindFunction("Function_Test")
	assert.NotEqual(s.T(), nil, fn)
	assert.Equal(s.T(), true, fn.exported)

	unexpFn := s.build.FindFunction("unexported_Test")
	assert.NotEqual(s.T(), nil, unexpFn)
	assert.Equal(s.T(), false, unexpFn.exported)
}

func (s *BuildSuite) TestFindStructAndMethods() {
	assert.NotEqual(s.T(), nil, s.build.FindStruct("X_Test"))
	assert.NotEqual(s.T(), nil, s.build.FindStruct("Y_Test"))

	xTest := s.build.FindStruct("X_Test")
	yTest := s.build.FindStruct("Y_Test")

	assert.NotEqual(s.T(), nil, xTest.Method("X_Test_Method"))
	assert.Equal(s.T(), nil, yTest.Method("Y_Test_Unexisting_Method"))
}

func (s *BuildSuite) TestStructures() {
	assert.Equal(s.T(), 2, len(s.build.structs))
}

func TestBuildSuite(t *testing.T) {
	suite.Run(t, &BuildSuite{})
}