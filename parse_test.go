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

var (
	SimpleDirFiles = [...]string{"1.go", "2.go"}
	SimpleDirPackages = [...]string{"dir", "dir"}
)

type ParseSuite struct {
	suite.Suite
}

func (s *ParseSuite) TestParseFile() {
	fileName := filepath.Base(SimpleFilePath)

	build, err := ParseFile(SimpleFilePath)
	assert.Equal(s.T(), nil, err)
	assert.NotEqual(s.T(), (*Build)(nil), build)
	assert.Equal(s.T(), 1, len(build.Files))
	assert.NotEqual(s.T(), nil, build.Files[fileName])
}

func (s *ParseSuite) TestParseDir() {
	build, err := ParseDir(SimpleDirPath)
	assert.Equal(s.T(), nil, err)
	assert.NotEqual(s.T(), (*Build)(nil), build)
	assert.Equal(s.T(), len(SimpleDirFiles), len(build.Files))

	for i, name := range SimpleDirFiles {
		assert.NotEqual(s.T(), (*File)(nil), build.Files[name])
		assert.Equal(s.T(), SimpleDirPackages[i], build.Files[name].Package())
	}
}

func (s *ParseSuite) TestParseFileAST() {

}

func TestParseSuite(t *testing.T) {
	suite.Run(t, new(ParseSuite))
}