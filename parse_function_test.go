package gogen

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"path/filepath"
	"github.com/stretchr/testify/assert"
)

type ParseFunctionSuite struct {
	suite.Suite

	build *Build
	file *File
}

func (s *ParseFunctionSuite) SetupTest() {
	var err error
	s.build, err = ParseFile(SimpleFilePath)
	assert.Equal(s.T(), nil, err)
	assert.NotEqual(s.T(), (*Build)(nil), s.build)
	s.file = s.build.Files[filepath.Base(SimpleFilePath)]
}

func (s *ParseFunctionSuite) TestParseFunction() {
	// there is only one function inside the simple.go
	assert.Equal(s.T(), 1, len(s.file.Functions()))
	// retrieve the Add function
	fun := s.file.Function("AddTwo")
	assert.NotEqual(s.T(), nil, fun)
	assert.Equal(s.T(), "AddTwo", fun.Name())
	assert.Equal(s.T(), false, fun.IsMethod())
}

func (s *ParseFunctionSuite) TestParseMethod() {
	st := s.file.Struct("X")
	assert.NotEqual(s.T(), nil, st)
	assert.Equal(s.T(), 2, len(st.Methods()))
	assert.Equal(s.T(), "Add", st.Method("Add").Name())
	assert.Equal(s.T(), "Copy", st.Method("Copy").Name())
}

func TestParseFunctionSuite(t *testing.T) {
	suite.Run(t, &ParseFunctionSuite{})
}