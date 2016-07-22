package gogen

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"github.com/stretchr/testify/assert"
	"path/filepath"
)

// Please note that this test suite refers to the
// test_fixtures/simple.go test file.
type ParseTypeSuite struct {
	suite.Suite
	build *Build
	file *File

	st *Structure
	in *Interface
}

func (s *ParseTypeSuite) SetupTest() {
	var err error
	s.build, err = ParseFile(SimpleFilePath)
	assert.Equal(s.T(), nil, err)
	s.file = s.build.Files[filepath.Base(SimpleFilePath)]

	s.st = s.file.Struct("X")
	assert.NotEqual(s.T(), (*Structure)(nil), s.st)

	s.in = s.file.Interface("Y")
	assert.NotEqual(s.T(), (*Interface)(nil), s.in)
}

// the parsing capability is already tested by the
// compiler, we test only the results of the ParseStruct
// that is already called by the ParseFile function in the
// test setup
func (s *ParseTypeSuite) TestParseStruct() {
	assert.Equal(s.T(), "X", s.st.Name())
}

func (s *ParseTypeSuite) TestParseInterface() {
	assert.Equal(s.T(), "Y", s.in.Name())
}

func (s *ParseTypeSuite) TestStructureFields() {
	assert.Equal(s.T(), 3, len(s.st.Fields()))

	assert.Equal(s.T(), "Val", s.st.Fields()[0].Name())
	intType, intSubType := s.st.Fields()[0].Type()
	assert.Equal(s.T(), "int", intType)
	assert.Equal(s.T(), PrimitiveType, intSubType)

	assert.Equal(s.T(), "SliceVal", s.st.Fields()[1].Name())
	sliceType, sliceSubtype := s.st.Fields()[1].Type()
	assert.Equal(s.T(), "string", sliceType)
	assert.Equal(s.T(), SliceType, sliceSubtype)

	assert.Equal(s.T(), "MapVal", s.st.Fields()[2].Name())
	mapType, mapSubtype := s.st.Fields()[2].Type()
	assert.Equal(s.T(), "[string]int", mapType)
	assert.Equal(s.T(), MapType, mapSubtype)
}

func TestParseTypeSuite(t *testing.T) {
	suite.Run(t, &ParseTypeSuite{})
}
