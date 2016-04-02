package architect

// Interface defines methods above the interface
// type for the interface build organization
type Interface interface {
	Name() string

	AddMethod(method Function)
	Method(name string) Function
}

// InterfaceImpl is a implmentation of the Interface
type InterfaceImpl struct {
	name string
}

// Name returns the name of the Interface
func (i *InterfaceImpl) Name() string {
	return i.name
}