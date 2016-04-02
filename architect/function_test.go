package architect

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"github.com/stretchr/testify/assert"
)

type FunctionSuite struct {
	suite.Suite
}

func (s *FunctionSuite) TestNewFunction() {
	assert.NotEqual(s.T(), nil, NewFunction())
}

func (s *FunctionSuite) TestName() {
	const funcName = "Test"
	const funcNameChange = "Change"

	// create function and set the default name
	f := NewFunction()
	f.name = funcName

	assert.Equal(s.T(), funcName, f.Name())
	// change the name
	f.name = funcNameChange

	assert.Equal(s.T(), funcNameChange, f.Name())
}

func (s *FunctionSuite) TestExported() {
	const exported = true
	const unexported = false

	f := NewFunction()
	f.exported = exported

	assert.Equal(s.T(), exported, f.Exported())
	f.exported = unexported

	assert.Equal(s.T(), unexported, f.Exported())
}

func TestFunctionSuite(t *testing.T) {
	suite.Run(t, new(FunctionSuite))
}