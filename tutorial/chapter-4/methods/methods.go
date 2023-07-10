package methods

import (
	"fmt"
	"math"
)

//Go no tiene clase. pero se puede definir metodos de tipo

type Vertex struct {
	X, Y float64
}

type MyFloat float64 // se puede definir una type sin estrutura y asociarle un metodo

func (v Vertex) Abs() float64 { // creado una función que pertenece a la estructura Vertex
	return math.Sqrt((v.X*v.X + v.Y*v.Y))
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

func CallTypeMethod() {
	fmt.Println("Methods....")
	v := Vertex{3, 4}
	fmt.Println(v.Abs()) // llamando a Abs a través de la estrutura.

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
