package fixture

import "math"

// @dao --asdf "val poi" --qwer 654
// @test -r="q w e r"
type X struct {
	Val int
	SliceVal []string
	MapVal map[string]int
}

// @func --name add
func (x *X) Add(y *X) {
	x.Val += int(math.Abs(float64(y.Val)))
}

// @func --name copy
func (x X) Copy(y *X) {
	y.Val = x.Val
}

// @func --name addtwo --nomethod
func AddTwo(a X, b X) X {
	return X{a.Val + b.Val, []string{}, make(map[string]int)}
}

// @interface --name y
type Y interface {
	quack()
}