package gogen

import (
  "github.com/stretchr/testify/suite"
  "testing"
  "github.com/stretchr/testify/assert"
  "path/filepath"
)

type ConstantSuite struct{
  suite.Suite

  file *File
}

func (s *ConstantSuite) SetupTest() {
  build, err := ParseFile(SimpleFilePath)
  assert.Nil(s.T(), err)
  s.file = build.File(filepath.Base(SimpleFilePath))
  assert.NotNil(s.T(), s.file)
}

func(s *ConstantSuite) TestParseConstant() {

  pi := s.file.Constant("Pi")
  assert.Equal(s.T(), "3.14", pi.Value())
  assert.Equal(s.T(), FloatType, pi.Type())

  strConst := s.file.Constant("StringConstant")
  assert.NotNil(s.T(), strConst)
  assert.Equal(s.T(), "\"qwer\"", strConst.Value())
  assert.Equal(s.T(), StringType, strConst.Type())

  consts := s.file.Constants()
  assert.NotNil(s.T(), consts)
  assert.Equal(s.T(), 2, len(consts))
}

func TestConstantSuite(t *testing.T) {
  suite.Run(t, &ConstantSuite{})
}
