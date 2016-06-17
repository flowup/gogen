package gogen

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"github.com/stretchr/testify/assert"
	"path/filepath"
)

const (
	SimpleFilePath = "test_fixtures/simple.go"
	SimpleDirPath = "test_fixtures/dir"
)

type ParseSuite struct {
	suite.Suite
}

func (s *ParseSuite) TestParseFile() {
	fileName := filepath.Base(SimpleFilePath)

	build, err := ParseFile(SimpleFilePath)
	assert.Equal(s.T(), nil, err)
	assert.NotEqual(s.T(), nil, build)
	assert.Equal(s.T(), 1, len(build.Files))
	assert.NotEqual(s.T(), nil, build.Files[fileName])
}

func (s *ParseSuite) TestParseDir() {
	build, err := ParseDir(SimpleDirPath)
	assert.Equal(s.T(), nil, err)
	assert.NotEqual(s.T(), nil, build)
	assert.Equal(s.T(), 2, len(build.Files))
}

func TestParseSuite(t *testing.T) {
	suite.Run(t, new(ParseSuite))
}