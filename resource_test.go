package gogen

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ResourceSuite struct {
	suite.Suite
}

func (s *ResourceSuite) SetupTest() {}

func (s *ResourceSuite) TestResourceAdd() {
	rc := ResourceContainer{}
	rc.Add(42)
	assert.Equal(s.T(), 42, rc[0])
	rc.Add(1.25)
	rc.Add("wololo")
	assert.Equal(s.T(), "wololo", rc[2])
}

func (s *ResourceSuite) TestSearch() {
	type tstruct struct {
		A int
		B string
	}

	rc := ResourceContainer{}
	rc.Add(42)
	rc.Add(&tstruct{1, "Hello"})
	rc.Add("string")
	rc.Add("foo")

	res := rc.Search(&tstruct{})
	assert.Equal(s.T(), 1, len(res))
	for _, val := range res {
		assert.Equal(s.T(), reflect.TypeOf(val), reflect.TypeOf(&tstruct{}))
	}

	res = rc.Search("str")
	assert.Equal(s.T(), 2, len(res))
}

func TestResourceSuite(t *testing.T) {
}
	suite.Run(t, &ResourceSuite{})
