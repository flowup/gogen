package unk

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SourceSuite struct {
	suite.Suite

	tfile, tcontent string
}

func (s *SourceSuite) SetupTest() {
	s.tfile = "source_test.go"
	bytes, err := ioutil.ReadFile(s.tfile)
	assert.Equal(s.T(), err, nil)

	s.tcontent = string(bytes)
}

func (s *SourceSuite) TestNewSource() {
	src := NewSource(s.tfile)
	assert.Equal(s.T(), s.tfile, src.Path)
	assert.Equal(s.T(), "", src.Name)
	assert.Equal(s.T(), 0, src.Type)
}

func (s *SourceSuite) TestResolve() {
	src := NewSource(s.tfile)
	err := src.Resolve()
	assert.Equal(s.T(), nil, err)
	assert.Equal(s.T(), s.tfile, src.Path)
	assert.Equal(s.T(), s.tfile, src.Name)
	assert.Equal(s.T(), FileSource, src.Type)

	src = NewSource("../gogen.go")
	err = src.Resolve()
	assert.Equal(s.T(), nil, err)
	assert.Equal(s.T(), "../gogen.go", src.Path)
	assert.Equal(s.T(), "gogen.go", src.Name)
	assert.Equal(s.T(), FileSource, src.Type)
}

func (s *SourceSuite) TestContent() {
	src := NewSource(s.tfile)
	err := src.Resolve()
	assert.Equal(s.T(), nil, err)

	src.Content()
}

func (s *SourceSuite) TestFetch() {
	src := NewSource(s.tfile)
	err := src.Resolve()
	assert.Equal(s.T(), nil, err)

	content, err := src.Fetch()
	assert.Equal(s.T(), nil, err)

	assert.Equal(s.T(), s.tcontent, content)
}

func TestSourceSuite(t *testing.T) {
	suite.Run(t, &SourceSuite{})
}
