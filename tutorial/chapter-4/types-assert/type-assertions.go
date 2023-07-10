package typesassert

import "fmt"

func CallTypeAsser() {
	var i interface{} = "hello" // creamos una iterface con un concreto type
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	//f = i.(float64) // panic: exection no manejado en tiempo de ejecucion
	//fmt.Println(f)
}
