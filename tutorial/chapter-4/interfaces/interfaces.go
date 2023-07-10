package interfaces

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

type VertexV3 struct {
	X, Y float64
}

func (v *VertexV3) Abs() float64 {
	return math.Sqrt(v.X*v.Y + v.Y*v.Y)
}

func UseInterface() {
	fmt.Println("Interfaces.....")
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := VertexV3{3, 4}
	a = f  // implementa Abser
	a = &v // implementa abser como puntero
	// No funciona porque la estrutura no implementa la inferface
	// a = v

	fmt.Println(a.Abs())
}
