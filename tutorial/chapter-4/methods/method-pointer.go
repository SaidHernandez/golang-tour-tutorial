package methods

import (
	"fmt"
	"math"
)

// type structure
type VertexV2 struct {
	X, Y float64
}

func (v VertexV2) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//rasanos para usar un pointer receiver
/*
	1. modificar el valor
	2. evitar estar copiando el valor en cada llamada
*/
// De esta manera afectamos el valor de la estrutura.
func (v *VertexV2) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

/* otra manera de pasar por parametro la referencia de la estrutuctura
func Scale(v *VertexV2, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}*/

func CallPointerMethod() {
	v := VertexV2{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
