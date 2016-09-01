package gogen

import "go/ast"


// Interface represents the interface type of a
// given build
type Interface struct {
  BaseType
  parent *ast.InterfaceType
  spec   *ast.TypeSpec
}

// NewInterface creates a new Interface type and returns
// it with the provided parent and spec.
func NewInterface(parent *ast.InterfaceType, spec *ast.TypeSpec, tagMap *TagMap) *Interface {
  i := &Interface{
    BaseType: BaseType{
      name: spec.Name,
      tags: tagMap,
    },
    parent: parent,
    spec:   spec,
  }

  return i
}


// ParseInterface will create an interface
// with parameters given and return it
func ParseInterface(spec *ast.TypeSpec, parent *ast.InterfaceType, comments ast.CommentMap) *Interface {
  i := NewInterface(parent, spec, ParseTags(comments))

  return i
}