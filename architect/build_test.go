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

	structsBuild *Build
	interfaceBuild *Build
}

const fixturePackageName = "fixture"
const fixtureSimpleStructsFilePath = "./_fixture/simple_structs.go"
const fixtureSimpleInterfaceFilePath = "./_fixture/simple_interface.go"

func (s *BuildSuite) SetupTest() {
	// Structs testing file parse
	var structsSet token.FileSet

	// parse the ast from the given file
	structuresAst, err := parser.ParseFile(&structsSet, fixtureSimpleStructsFilePath, nil, parser.AllErrors)
	assert.Equal(s.T(), nil, err)

	// create build and make it from the AST
	s.structsBuild = NewBuild()
	s.structsBuild.Make(structuresAst)

	// Interface testing file parse
	var interfaceSet token.FileSet

	interfaceAst, err := parser.ParseFile(&interfaceSet, fixtureSimpleInterfaceFilePath, nil, parser.AllErrors)
	assert.Equal(s.T(), nil, err)

	s.interfaceBuild = NewBuild()
	s.interfaceBuild.Make(interfaceAst)
}

func (s *BuildSuite) TestPackageName() {
	assert.Equal(s.T(), fixturePackageName, s.structsBuild.Package())
}

func (s *BuildSuite) TestStructDefinitions() {
	// check for struct length
	assert.Equal(s.T(), 2, len(s.structsBuild.structs))

	// check struct properties
	assert.Equal(s.T(), "X_Test", s.structsBuild.structs[0].Name())
	assert.Equal(s.T(), "Y_Test", s.structsBuild.structs[1].Name())
}

func (s *BuildSuite) TestImports() {
	assert.Equal(s.T(), 2, len(s.structsBuild.imports))

	assert.Equal(s.T(), true, s.structsBuild.HasImport("fmt"))
	assert.Equal(s.T(), true, s.structsBuild.HasImport("math"))
}

func (s *BuildSuite) TestFunctions() {
	assert.Equal(s.T(), 4, len(s.structsBuild.functions))

	fn := s.structsBuild.FindFunction("Function_Test")
	assert.NotEqual(s.T(), nil, fn)
	assert.Equal(s.T(), true, fn.Exported())

	unexpFn := s.structsBuild.FindFunction("unexported_Test")
	assert.NotEqual(s.T(), nil, unexpFn)
	assert.Equal(s.T(), false, unexpFn.Exported())
}

func (s *BuildSuite) TestFindStructAndMethods() {
	assert.NotEqual(s.T(), nil, s.structsBuild.FindStruct("X_Test"))
	assert.NotEqual(s.T(), nil, s.structsBuild.FindStruct("Y_Test"))

	xTest := s.structsBuild.FindStruct("X_Test")
	yTest := s.structsBuild.FindStruct("Y_Test")

	assert.NotEqual(s.T(), nil, xTest.Method("X_Test_Method"))
	assert.Equal(s.T(), nil, yTest.Method("Y_Test_Unexisting_Method"))
}

func (s *BuildSuite) TestStructures() {
	assert.Equal(s.T(), 2, len(s.structsBuild.structs))
}

func TestBuildSuite(t *testing.T) {
	suite.Run(t, &BuildSuite{})
}