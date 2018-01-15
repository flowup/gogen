package gogen

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// Please note that this test suite refers to the
// test_fixtures/simple.go test file.
type ParseArraySuite struct {
	suite.Suite
	build        *Build
	file         *File
	complexBuild *Build
	complexFile  *File

	a *Array
}


func (s *ParseArraySuite) SetupTest() {
	var err error
	s.build, err = ParseFile(SimpleFilePath)
	assert.Equal(s.T(), nil, err)
	s.file = s.build.File(filepath.Base(SimpleFilePath))

	s.a = s.file.Array("A")
	assert.NotEqual(s.T(), (*Array)(nil), s.a)

	s.complexBuild, err = ParseFile(ComplexFilePath)
	assert.Equal(s.T(), nil, err)
	s.complexFile = s.complexBuild.File(filepath.Base(ComplexFilePath))
	assert.NotEqual(s.T(), (*File)(nil), s.complexFile)
}

// the parsing capability is already tested by the
// compiler, we test only the results of the ParseStruct
// that is already called by the ParseFile function in the
// test setup
func (s *ParseArraySuite) TestParseInt() {
	assert.Equal(s.T(), "A", s.a.Name())
	assert.Equal(s.T(), "int", s.a.baseType)
}

func TestParseArraySuite(t *testing.T) {
	suite.Run(t, &ParseArraySuite{})
}
