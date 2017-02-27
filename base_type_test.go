package gogen

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// Please note that this test suite refers to the
// test_fixtures/simple.go test file.
type ParseBaseTypeSuite struct {
	suite.Suite
	build        *Build
	file         *File
	complexBuild *Build
	complexFile  *File

	iba *BaseType
	sba *BaseType
}

func (s *ParseBaseTypeSuite) SetupTest() {
	var err error
	s.build, err = ParseFile(SimpleFilePath)
	assert.Equal(s.T(), nil, err)
	s.file = s.build.File(filepath.Base(SimpleFilePath))


	s.iba = s.file.BaseType("I")
	assert.NotEqual(s.T(), (*BaseType)(nil), s.iba)

	s.sba = s.file.BaseType("S")
	assert.NotEqual(s.T(), (*BaseType)(nil), s.sba)

	s.complexBuild, err = ParseFile(ComplexFilePath)
	assert.Equal(s.T(), nil, err)
	s.complexFile = s.complexBuild.File(filepath.Base(ComplexFilePath))
	assert.NotEqual(s.T(), (*File)(nil), s.complexFile)
}

// the parsing capability is already tested by the
// compiler, we test only the results of the ParseStruct
// that is already called by the ParseFile function in the
// test setup


func (s *ParseBaseTypeSuite) TestParseInt() {
	assert.Equal(s.T(), "I", s.iba.Name())
	assert.Equal(s.T(), "int", s.iba.Type())
}

func (s *ParseBaseTypeSuite) TestParseString() {
	assert.Equal(s.T(), "S", s.sba.Name())
	assert.Equal(s.T(), "string", s.sba.Type())
}

func TestParseBaseTypeSuite(t *testing.T) {
	suite.Run(t, &ParseBaseTypeSuite{})
}
