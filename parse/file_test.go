package parse

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type FileSuite struct {
	suite.Suite
}

func (s *FileSuite) SetupTest() {}

func (s *FileSuite) TestFileParse() {
	build, err := File("./file.go")
	assert.Equal(s.T(), nil, err)
	assert.NotEqual(s.T(), nil, build)
	assert.Equal(s.T(), "parse", build.Package())
}

func TestFileSuite(t *testing.T) {
	suite.Run(t, new(FileSuite))
}
