package gogen

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	SimpleFilePath    = "test_fixtures/simple.go"
	InterfaceFilePath = "test_fixtures/interfaces_go"
	ComplexFilePath   = "test_fixtures/complex.go"
	SimpleDirPath     = "test_fixtures/dir"
)

var (
	SimpleDirFiles    = [...]string{"1.go", "2.go"}
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
	// there should be only one file, as only one was
	// parsed by the build
	assert.Equal(s.T(), 1, len(build.Files()))
	assert.NotEqual(s.T(), nil, build.File(fileName))
}

func (s *ParseSuite) TestParseDir() {
	build, err := ParseDir(SimpleDirPath)
	assert.Equal(s.T(), nil, err)
	assert.NotEqual(s.T(), (*Build)(nil), build)
	assert.Equal(s.T(), len(SimpleDirFiles), len(build.Files()))

	for i, name := range SimpleDirFiles {
		assert.NotEqual(s.T(), (*File)(nil), build.File(name))
		assert.Equal(s.T(), SimpleDirPackages[i], build.File(name).Package())
	}
}

/*func (s *ParseSuite) TestParseInterface() {
	fileName := filepath.Base(InterfaceFilePath)

	build, err := ParseFile(InterfaceFilePath)
	s.Nil(err)

	file := build.File(fileName)

	for _, s := range file.Interfaces() {
		fmt.Println(s.name)

		for _, r := range s.methods {
			for _, y := range r.Names {
				fmt.Println(y.Obj.Kind, " ", y.Obj.Name)
			}
		}
	}

	//assert.Equal(s.T(), nil, err)
	//assert.NotEqual(s.T(), (*Build)(nil), build)
	//// there should be only one file, as only one was
	//// parsed by the build
	//assert.Equal(s.T(), 1, len(build.Files()))
	//assert.NotEqual(s.T(), nil, build.File(fileName))
}*/

func (s *ParseSuite) TestParseFileAST() {

}

func TestParseSuite(t *testing.T) {
	suite.Run(t, new(ParseSuite))
}
