package gogen

import (
  "go/ast"
  "strings"
)

// Import is a structure representing import string in ast tree
type Import struct{
  BaseType
  parent *ast.ImportSpec

  importString string
}

// String will return import string
func (i *Import) String() string {
  return i.importString
}

// NewImport will construct a new import
func NewImport(parent *ast.ImportSpec, tags *TagMap) *Import{
  var n string
  if parent.Name != nil {
    n = parent.Name.Name
  } else {
    raw := strings.Replace(parent.Path.Value, "\"", "", -1)
    split := strings.Split(raw, "/")
    n = split[len(split) - 1]
  }
  return &Import{
    BaseType: BaseType{
      name: n,
      tags: tags,
    },
    parent: parent,
    importString: parent.Path.Value,
  }
}

// ParseImport will parse an import and return its structure
func ParseImport(parent *ast.ImportSpec, comments ast.CommentMap) *Import {
  i := NewImport(parent, ParseTags(comments))

  return i
}