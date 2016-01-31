package gogen

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PipelineSuite struct {
	suite.Suite
}

func (s *PipelineSuite) TestAdd() {
	p := Pipeline{}
	p.Add(&testGenerator{})
	p.Add(&testGenerator{})

	assert.Equal(s.T(), 2, len(p.generators))
}

func (s *PipelineSuite) TestSize() {
	p := Pipeline{}
	assert.Equal(s.T(), 0, p.Size())
	p.Add(&testGenerator{})
	assert.Equal(s.T(), 1, p.Size())
}

func TestPipelineSuite(t *testing.T) {
	suite.Run(t, &PipelineSuite{})
}
