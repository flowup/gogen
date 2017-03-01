package fixture

import (
	"math"
	"time"
)

// @pi
const Pi = 3.14
const StringConstant = "qwer"

// @dao --asdf "val poi" --qwer 654
// @test -r="q w e r"
type X struct {
	Val      int // @field
	SliceVal []string `gorm:"index"`
	MapVal   map[string]int `json:"map_val"`
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

type I int

type S string

type A []int

type Model struct {
	Ptr *int
	DeletedAt *time.Time
}