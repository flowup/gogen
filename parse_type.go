package gogen

import "go/ast"

const (
	PrimitiveType = 0
	SliceType = 1
	MapType = 2
)

// StructField encapsulates one field of the Structure
type StructField struct {
	parent *ast.Field

	name string
}

// Name returns the name of the field
func (f *StructField) Name() string {
	return f.name
}

// Type returns type of the field as a string and
// FieldType such as Slice, Map or Primitive
func (f *StructField) Type() (string, int) {
	switch t := f.parent.Type.(type) {
	case *ast.Ident:
		return t.Name, PrimitiveType
	case *ast.ArrayType:
		return t.Elt.(*ast.Ident).Name, SliceType
	case *ast.MapType:
		return "[" + t.Key.(*ast.Ident).Name + "]" + t.Value.(*ast.Ident).Name, MapType
	default:
		panic("StructField type not recognized! Please report this issue.")
	}
}

// Structure represents the struct type of a
// given build
type Structure struct {
	parent *ast.StructType
	spec *ast.TypeSpec

	// map of methods
	methods map[string]*Function
}

// NewStructure returns new Instance of the structure type
// with the provided parent and type spec.
func NewStructure(parent *ast.StructType, spec *ast.TypeSpec) *Structure {
	return &Structure{
		parent: parent,
		spec: spec,
		methods: make(map[string]*Function),
	}
}

// Name returns the name of the given structure
func (s *Structure) Name() string {
	return s.spec.Name.Name
}

// Fields returns fields that are associated with the
// given Structure. Note that this function builds
// the field list every time it is called, so cache
// the results to improve the performance.
func (s *Structure) Fields() []*StructField {
	fields := []*StructField{}

	for _, field := range s.parent.Fields.List {
		for _, fieldName := range field.Names {
			fields = append(fields, &StructField{
				field, fieldName.Name,
			})
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

// Interface represents the interface type of a
// given build
type Interface struct {
	parent *ast.InterfaceType
	spec *ast.TypeSpec
}

// NewInterface creates a new Interface type and returns
// it with the provided parent and spec.
func NewInterface(parent *ast.InterfaceType, spec *ast.TypeSpec) *Interface {
	return &Interface{
		parent: parent,
		spec: spec,
	}
}

// Name returns the name of the interface type
func (i *Interface) Name()  string {
	return i.spec.Name.Name
}

func ParseStruct(spec *ast.TypeSpec, parent *ast.StructType) *Structure {
	s := NewStructure(parent, spec)

	return s
}

func ParseInterface(spec *ast.TypeSpec, parent *ast.InterfaceType) *Interface {
	i := NewInterface(parent, spec)

	return i
}