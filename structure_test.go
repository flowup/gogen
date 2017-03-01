package gogen

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// Please note that this test suite refers to the
// test_fixtures/simple.go test file.
type ParseTypeSuite struct {
	suite.Suite
	build        *Build
	file         *File
	complexBuild *Build
	complexFile  *File

	st *Structure
	in *Interface
	mdl *Structure
}

func (s *ParseTypeSuite) SetupTest() {
	var err error
	s.build, err = ParseFile(SimpleFilePath)
	assert.Equal(s.T(), nil, err)
	s.file = s.build.File(filepath.Base(SimpleFilePath))

	s.st = s.file.Struct("X")
	assert.NotEqual(s.T(), (*Structure)(nil), s.st)

	s.in = s.file.Interface("Y")
	assert.NotEqual(s.T(), (*Interface)(nil), s.in)

	s.mdl = s.file.Struct("Model")
	assert.NotEqual(s.T(), (*Structure)(nil), s.mdl)

	s.complexBuild, err = ParseFile(ComplexFilePath)
	assert.Equal(s.T(), nil, err)
	s.complexFile = s.complexBuild.File(filepath.Base(ComplexFilePath))
	assert.NotEqual(s.T(), (*File)(nil), s.complexFile)
}

// the parsing capability is already tested by the
// compiler, we test only the results of the ParseStruct
// that is already called by the ParseFile function in the
// test setup
func (s *ParseTypeSuite) TestParseStruct() {
	assert.Equal(s.T(), "X", s.st.Name())
}

func (s *ParseTypeSuite) TestParseInterface() {
	assert.Equal(s.T(), "Y", s.in.Name())
}

func (s *ParseTypeSuite) TestParseModel() {
	assert.Equal(s.T(), "Model", s.mdl.Name())
}

func (s *ParseTypeSuite) TestStructureFields() {
	assert.Equal(s.T(), 3, len(s.st.Fields()))

	intType, intSubType := s.st.Fields()["Val"].Type()
	assert.Equal(s.T(), "int", intType)
	assert.Equal(s.T(), PrimitiveType, intSubType)

	sliceType, sliceSubtype := s.st.Fields()["SliceVal"].Type()
	assert.Equal(s.T(), "string", sliceType)
	assert.Equal(s.T(), SliceType, sliceSubtype)

	sliceTag := s.st.Fields()["SliceVal"].Tag()
	assert.Equal(s.T(), `gorm:"index"`,sliceTag )

	valTag := s.st.Fields()["Val"].Tag()
	assert.Equal(s.T(), ``,valTag )

	mapType, mapSubtype:= s.st.Fields()["MapVal"].Type()
	assert.Equal(s.T(), "[string]int", mapType)
	assert.Equal(s.T(), MapType, mapSubtype)

	mapTag := s.st.Fields()["MapVal"].Tag()
	assert.Equal(s.T(), `json:"map_val"`,mapTag)

	deleteType, timeSubType := s.mdl.Fields()["DeletedAt"].Type()
	assert.Equal(s.T(), "time.Time", deleteType)
	assert.Equal(s.T(), PointerType, timeSubType)

	ptrType, ptrSubType := s.mdl.Fields()["Ptr"].Type()
	assert.Equal(s.T(), "int", ptrType)
	assert.Equal(s.T(), PointerType, ptrSubType)
}

func (s *ParseTypeSuite) TestStructureComplexFields() {
	str := s.complexFile.Struct("MyStruct")
	fType, _ := str.Fields()["MyTime"].Type()
	assert.Equal(s.T(), "time.Time", fType)
}

func TestParseTypeSuite(t *testing.T) {
	suite.Run(t, &ParseTypeSuite{})
}
