package pointers

import "fmt"

/*
Pointer su valor por defecto es nill
The & asigana un puntero a la variable
*/
func PointerDefinition() {
	i, j := 42, 2701

	p := &i
	fmt.Println("pointer p", *p) //read i through the pointer
	*p = 21                      //set i through the pointer deferencing
	fmt.Println("pointer i", i)

	p = &j
	*p = *p / 37
	fmt.Println("pointer j", j)
}
