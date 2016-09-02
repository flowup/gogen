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
  fields map[string]*StructField
	// map of methods
	methods map[string]*Function
}

// NewStructure returns new Instance of the structure type
// with the provided parent and type spec.
func NewStructure(parent *ast.StructType, spec *ast.TypeSpec, fMap map[string]*StructField, tagMap *TagMap) *Structure {
	s := &Structure{
    BaseType: BaseType{
      name: spec.Name,
      tags: tagMap,
    },
    parent : parent,
    spec: spec,
    fields: fMap,
    methods: make(map[string]*Function),
  }

  return s
}

// FilteredStructFields is a map of struct fields
// that can be filtered by their tags
type FilteredStructFields map[string]*StructField

// Filter will filter struct fields map and return
// all those that have tag with name given by parameter
func(f FilteredStructFields) Filter(name string) map[string]*StructField {
  newMap := make(map[string]*StructField)
  for it := range f {
    if f[it].Tags().Has(name) {
      newMap[it] = f[it]
    }
  }

  return newMap
}

// Fields returns fields that are associated with the
// given Structure.
func (s *Structure) Fields() FilteredStructFields {
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
  fMap := make(map[string]*StructField)
  for _, field := range parent.Fields.List {
    if len(field.Names) == 0 {
      continue
    }
    fMap[field.Names[0].Name] = ParseStructField(field, comments.Filter(field))
  }

  s := NewStructure(parent, spec, fMap, ParseTags(comments))

	return s
}
