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

	// map of methods
	methods map[string]*Function
}

// NewStructure returns new Instance of the structure type
// with the provided parent and type spec.
func NewStructure(parent *ast.StructType, spec *ast.TypeSpec, tagMap *TagMap) *Structure {
	s := &Structure{
    BaseType: BaseType{
      name: spec.Name,
      tags: tagMap,
    },
    parent : parent,
    spec: spec,
    methods: make(map[string]*Function),
  }

  return s
}

// Fields returns fields that are associated with the
// given Structure. Note that this function builds
// the field list every time it is called, so cache
// the results to improve the performance.
func (s *Structure) Fields() []*StructField {
	fields := []*StructField{}

  for _, field := range s.parent.Fields.List {
    for _, fieldName := range field.Names {
      newField := &StructField{
        BaseType: BaseType{
          name: fieldName,
          tags: nil,
        },
        parent:field,
      }

      fields = append(fields, newField)
    }
  }

	return fields
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
	s := NewStructure(parent, spec, ParseTags(comments))

	return s
}
