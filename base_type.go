package gogen

// BaseType is a base structure of all declarations
// It can return a name and tag map.
type BaseType struct {
  name string

  tags *TagMap
}

// Name will return a name of a declaration
func (b *BaseType) Name() string{
  return b.name
}

// Tags will return a tag map associated to a declaration
func (b *BaseType) Tags() *TagMap{
  return b.tags
}