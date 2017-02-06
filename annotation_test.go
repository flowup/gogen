package gogen

import (
  "github.com/stretchr/testify/suite"
  "github.com/stretchr/testify/assert"
  "path/filepath"
  "testing"
)

// Please note that this test suite refers to the
// test_fixtures/simple.go test file.
type ParseAnnotationsSuite struct {
  suite.Suite
  build *Build
  file *File

  st *Structure
  in *Interface
}

func (s *ParseAnnotationsSuite) SetupTest() {
  var err error
  s.build, err = ParseFile(SimpleFilePath)
  assert.Equal(s.T(), nil, err)
  s.file = s.build.File(filepath.Base(SimpleFilePath))

  s.st = s.file.Struct("X")
  assert.NotEqual(s.T(), (*Structure)(nil), s.st)

  s.in = s.file.Interface("Y")
  assert.NotEqual(s.T(), (*Interface)(nil), s.in)
}

func (s *ParseAnnotationsSuite) TestParseAnnotations() {
  //assert.Equal(s.T(), true, s.st.Annotations().HasAnnotation("@dao"))
  daoAnnotation, ok := s.st.Annotations().Get("@dao")
  assert.Equal(s.T(), true, ok)
  paramVal, ok := daoAnnotation.Get("asdf")
  assert.Equal(s.T(), true, ok)
  assert.Equal(s.T(), "val poi", paramVal)

  interAnnotation, ok := s.in.Annotations().Get("@interface")
  assert.Equal(s.T(), true, ok)
  name, ok := interAnnotation.Get("name")
  assert.Equal(s.T(), true, ok)
  assert.Equal(s.T(), "y", name)
}

func (s *ParseAnnotationsSuite) TestAnnotation() {
  tag := NewAnnotation("customAnnotation")
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

func (s *ParseAnnotationsSuite) TestAnnotationMap() {
  tm := NewAnnotationMap()
  tag := NewAnnotation("customAnnotation")
  tm.Set("customAnnotation", tag)

  assert.Equal(s.T(), true, tm.Has("customAnnotation"))

  assert.Equal(s.T(), 1, tm.Num())

  names := tm.GetAnnotationNames()
  assert.Equal(s.T(), "customAnnotation", names[0])

  tm.Delete("customAnnotation")
  assert.Equal(s.T(), 0, tm.Num())

  newAnnotation := NewAnnotation("newAnnotation")
  newAnnotation2 := NewAnnotation("newAnnotation2")

  tm.Set("newAnnotation", newAnnotation)
  tm.Set("newAnnotation2", newAnnotation2)
  tags := tm.GetAll()
  assert.Equal(s.T(), 2, len(tags))
}

func TestParseAnnotationsSuite(t *testing.T) {
  suite.Run(t, &ParseAnnotationsSuite{})
}