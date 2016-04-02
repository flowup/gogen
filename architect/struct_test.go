package architect

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"github.com/stretchr/testify/assert"
)

const defaultName = "test"
const changedName = "changed"

type StructSuite struct {
	suite.Suite
}

func (s *StructSuite) TestNew() {
	assert.NotEqual(s.T(), nil, NewStruct(defaultName))
}

func (s *StructSuite) TestName() {
	st := NewStruct(defaultName)
	assert.Equal(s.T(), defaultName, st.Name())

	st.name = changedName
	assert.Equal(s.T(), changedName, st.Name())
}

func (s *StructSuite) TestField() {
}

func (s *StructSuite) TestMethod() {
}

func TestStructTest(t *testing.T) {
	suite.Run(t, new(StructSuite))
}
