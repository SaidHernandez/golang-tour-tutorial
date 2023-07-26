package generics

import "fmt"

/*
[T comparable]: Es una construcción de tipo genérico en Go.

	Indica que la función Index es genérica y puede trabajar con cualquier tipo T que sea comparable.
	T es un nombre de tipo genérico que se puede sustituir por cualquier tipo concreto
*/
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}

	return -1
}

type List[T any] struct {
	next *List[T]
	val  T
}

func CallIndex() {
	fmt.Println("Index...")
	si := []int{1, 24, 56, 78}
	fmt.Println(Index(si, 24))

	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))
}
