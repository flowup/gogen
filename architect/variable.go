package architect

// Variable is any field from the Struct field, to function
// parameter fields
type Variable interface {
	Name() string
}
