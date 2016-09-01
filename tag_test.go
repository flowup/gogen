package gogen

import (
  "github.com/stretchr/testify/suite"
  "github.com/stretchr/testify/assert"
  "path/filepath"
  "testing"
)

// Please note that this test suite refers to the
// test_fixtures/simple.go test file.
type ParseTagsSuite struct {
  suite.Suite
  build *Build
  file *File

  st *Structure
  in *Interface
}

func (s *ParseTagsSuite) SetupTest() {
  var err error
  s.build, err = ParseFile(SimpleFilePath)
  assert.Equal(s.T(), nil, err)
  s.file = s.build.Files[filepath.Base(SimpleFilePath)]

  s.st = s.file.Struct("X")
  assert.NotEqual(s.T(), (*Structure)(nil), s.st)

  s.in = s.file.Interface("Y")
  assert.NotEqual(s.T(), (*Interface)(nil), s.in)
}

func (s *ParseTagsSuite) TestParseTags() {
  //assert.Equal(s.T(), true, s.st.Tags().HasTag("@dao"))
  daoTag, ok := s.st.Tags().Get("@dao")
  assert.Equal(s.T(), true, ok)
  paramVal, ok := daoTag.Get("asdf")
  assert.Equal(s.T(), true, ok)
  assert.Equal(s.T(), "val poi", paramVal)

  interTag, ok := s.in.Tags().Get("@interface")
  assert.Equal(s.T(), true, ok)
  name, ok := interTag.Get("name")
  assert.Equal(s.T(), true, ok)
  assert.Equal(s.T(), "y", name)
}

func (s *ParseTagsSuite) TestTag() {
  tag := NewTag("customTag")
  tag.Set("customPar", "customVal")

  assert.Equal(s.T(), true, tag.Has("customPar"))

  assert.Equal(s.T(), 1, tag.Num())

  names := tag.GetParameterNames()
  assert.Equal(s.T(), "customPar", names[0])

  tag.Delete("customPar")
  assert.Equal(s.T(), 0, tag.Num())

  tag.Set("newPar", "newVal")
  tag.Set("newPar2", "newVal2")

  assert.Equal(s.T(), 2, len(tag.GetAll()))
}

func (s *ParseTagsSuite) TestTagMap() {
  tm := NewTagMap()
  tag := NewTag("customTag")
  tm.Set("customTag", tag)

  assert.Equal(s.T(), true, tm.Has("customTag"))

  assert.Equal(s.T(), 1, tm.Num())

  names := tm.GetTagNames()
  assert.Equal(s.T(), "customTag", names[0])

  tm.Delete("customTag")
  assert.Equal(s.T(), 0, tm.Num())

  newTag := NewTag("newTag")
  newTag2 := NewTag("newTag2")

  tm.Set("newTag", newTag)
  tm.Set("newTag2", newTag2)
  tags := tm.GetAll()
  assert.Equal(s.T(), 2, len(tags))
}

func TestParseTagsSuite(t *testing.T) {
  suite.Run(t, &ParseTagsSuite{})
}