package gogen

import (
  "go/ast"
)

// Structure represents the struct type of a
// given build
type Structure struct {
  BaseType
  parent *ast.StructType
	spec   *ast.TypeSpec

  // slice of fields
  fields []*StructField
	// map of methods
	methods map[string]*Function
}

// NewStructure returns new Instance of the structure type
// with the provided parent and type spec.
func NewStructure(parent *ast.StructType, spec *ast.TypeSpec, fList []*StructField, tagMap *TagMap) *Structure {
	s := &Structure{
    BaseType: BaseType{
      name: spec.Name,
      tags: tagMap,
    },
    parent : parent,
    spec: spec,
    fields: fList,
    methods: make(map[string]*Function),
  }

  return s
}

// Fields returns fields that are associated with the
// given Structure.
func (s *Structure) Fields() []*StructField {
	return s.fields
}

// AddMethod binds a method into the current structure
func (s *Structure) AddMethod(fun *Function) {
	s.methods[fun.Name()] = fun
}

// Method returns a Function bound to the current structure
func (s *Structure) Method(name string) *Function {
	return s.methods[name]
}

// Methods returns all Function-s found to the current structure
func (s *Structure) Methods() map[string]*Function {
	return s.methods
}

// ParseStruct will create a structure
// with parameters given and return it
func ParseStruct(spec *ast.TypeSpec, parent *ast.StructType, comments ast.CommentMap) *Structure {
  var fList []*StructField
  for _, field := range parent.Fields.List {
    fList = append(fList, ParseStructField(field, comments.Filter(field)))
  }

  s := NewStructure(parent, spec, fList, ParseTags(comments))

	return s
}
