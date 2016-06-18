package fixture

import "math"

type X struct {
	Val int
}

func (x *X) Add(y *X) {
	x.Val += int(math.Abs(float64(y.Val)))
}

type Y interface {
	quack()
}