/*
An interface type is defined as a set of method signatures.

A value of interface type can hold any value that implements those methods.

Note: There is an error in the example code on line 22. Vertex (the value type) doesn't implement Abser because the Abs method is defined only on *Vertex (the pointer type).
*/

package main

import (
	"fmt"
	"math"
)

type MyFloat float64
type Abser interface {
	Abs() float64
}
type Vertex struct {
	X, Y float64
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{ 3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// a = v // ERROR!
	fmt.Println(a.Abs())
}
