package fixture

type X struct {
	Val int
}

func (x *X) Add(y *X) {
	x.Val += y.Val
}