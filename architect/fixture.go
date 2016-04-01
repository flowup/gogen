package architect

import (
	"fmt"
	"math"
)

// This file is meant to have testing only
// purpose. Please don't use any types, or
// functions provided by it. It may be frequently
// changed or moved.

type X_Test struct {

}

func (x *X_Test) X_Test_Method() {

}

type Y_Test struct {

}

func Function_Test() {

}

func unexported_Test() {

}

func ParametricFunction_Test(a string, b, c int) {
	math.Pow(1, 2)
}

func ReturnFunction_Test() string {
	fmt.Println("here")

	return "test"
}