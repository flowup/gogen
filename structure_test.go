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

	st  *Structure
	in  *Interface
	mdl *Structure
	z   *Structure
	w   *Structure
	v   *Structure
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

	s.z = s.file.Struct("Z")
	assert.NotEqual(s.T(), (*Structure)(nil), s.z)

	s.w = s.file.Struct("W")
	assert.NotEqual(s.T(), (*Structure)(nil), s.w)

	s.v = s.file.Struct("V")
	assert.NotEqual(s.T(), (*Structure)(nil), s.v)

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

func (s *ParseTypeSuite) TestParseZ() {
	assert.Equal(s.T(), "Z", s.z.Name())
}

func (s *ParseTypeSuite) TestStructureFields() {
	assert.Equal(s.T(), 4, len(s.st.Fields()))

	intType, intSubType := s.st.Fields()["Val"].Type()
	assert.Equal(s.T(), "int", intType)
	assert.Equal(s.T(), PrimitiveType, intSubType)

	sliceType, sliceSubtype := s.st.Fields()["SliceVal"].Type()
	assert.Equal(s.T(), "string", sliceType)
	assert.Equal(s.T(), SliceType, sliceSubtype)

	sliceTag := s.st.Fields()["SliceVal"].Tag()
	assert.Equal(s.T(), `gorm:"index"`, sliceTag)

	valTag := s.st.Fields()["Val"].Tag()
	assert.Equal(s.T(), ``, valTag)

	mapType, mapSubtype := s.st.Fields()["MapVal"].Type()
	assert.Equal(s.T(), "[string]int", mapType)
	assert.Equal(s.T(), MapType, mapSubtype)

	mapTag := s.st.Fields()["MapVal"].Tag()
	assert.Equal(s.T(), `json:"map_val"`, mapTag)

	interType, interSubtype := s.st.Fields()["Inter"].Type()
	interName := s.st.Fields()["Inter"].Name()
	assert.Equal(s.T(), "Inter", interName)
	assert.Equal(s.T(), "interface", interType)
	assert.Equal(s.T(), InterfaceType, interSubtype)

	interTag := s.st.Fields()["Inter"].Tag()
	assert.Equal(s.T(), ``, interTag)

	deleteType, timeSubType := s.mdl.Fields()["DeletedAt"].Type()
	assert.Equal(s.T(), "time.Time", deleteType)
	assert.Equal(s.T(), PointerType, timeSubType)

	ptrType, ptrSubType := s.mdl.Fields()["Ptr"].Type()
	assert.Equal(s.T(), "int", ptrType)
	assert.Equal(s.T(), PointerType, ptrSubType)

	modelType, _ /*modelSubType */ := s.z.Fields()["_Model"].Type()
	assert.Equal(s.T(), "Model", modelType)

	wType, _ /*wSubType */ := s.w.Fields()["_Z"].Type()
	assert.Equal(s.T(), "Z", wType)

	vWType, _ /*vWSubType */ := s.v.Fields()["_W"].Type()
	assert.Equal(s.T(), "W", vWType)

	vZType, _ /*vZSubType */ := s.v.Fields()["_Z"].Type()
	assert.Equal(s.T(), "Z", vZType)

}

func (s *ParseTypeSuite) TestStructureComplexFields() {
	str := s.complexFile.Struct("MyStruct")
	fType, _ := str.Fields()["MyTime"].Type()
	assert.Equal(s.T(), "time.Time", fType)
}

func TestParseTypeSuite(t *testing.T) {
	suite.Run(t, &ParseTypeSuite{})
}
