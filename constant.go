package gogen

import (
  "go/ast"
  "go/token"
)

const (
  IntType       = 1024
  FloatType     = 1025
  CharType      = 1026
  StringType    = 1027
)

// Constant is a structure for a constant in Go program
type Constant struct {
  BaseType
  parent *ast.ValueSpec
}

// NewConstant is a factory method for constants
func NewConstant(parent *ast.ValueSpec, tagMap *TagMap) *Constant {
  c := &Constant{
    BaseType: BaseType{
      name: parent.Names[0],
      tags: tagMap,
    },
    parent: parent,
  }

  return c
}

// Value will return a sting representation of a value of a constant
func (c *Constant) Value() string{
  return c.parent.Values[0].(*ast.BasicLit).Value
}

// Type will return a type of a constant
func (c *Constant) Type() int {
  switch t := c.parent.Values[0].(type) {
  case *ast.BasicLit:
    switch t.Kind {
    case token.INT:
      return IntType
    case token.CHAR:
      return CharType
    case token.STRING:
      return StringType
    case token.FLOAT:
      return FloatType
    default:
      return Unrecognized
    }
  default:
    return Unrecognized
  }
}


// ParseConstant will create a constant
// with parameters given and return it
func ParseConstant(parent *ast.ValueSpec, comments ast.CommentMap) *Constant {
  c := NewConstant(parent, ParseTags(comments))

  return c
}