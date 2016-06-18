package gogen

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"github.com/stretchr/testify/assert"
	"path/filepath"
)

type ParseTypeSuite struct {
	suite.Suite
	build *Build
	file *File
}

func (s *ParseTypeSuite) SetupTest() {
	var err error
	s.build, err = ParseFile(SimpleFilePath)
	assert.Equal(s.T(), nil, err)
	s.file = s.build.Files[filepath.Base(SimpleFilePath)]
}

// the parsing capability is already tested by the
// compiler, we test only the results of the ParseStruct
// that is already called by the ParseFile function in the
// test setup
func (s *ParseTypeSuite) TestParseStruct() {
	st := s.file.Struct("X")
	assert.NotEqual(s.T(), (*Structure)(nil), st)
}

func (s *ParseTypeSuite) TestParseInterface() {
	in := s.file.Interface("Y")
	assert.NotEqual(s.T(), (*Interface)(nil), in)
}

func TestParseTypeSuite(t *testing.T) {
	suite.Run(t, &ParseTypeSuite{})
}
