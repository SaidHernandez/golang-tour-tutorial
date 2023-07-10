package arraies

import (
	"fmt"
	"strings"
)

func ArraiesDeclaration() {
	var a [2]string //fixed size
	a[0] = "Hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13} //slice
	var s []int = primes[0:4]            //exclude the last one
	fmt.Println(primes)
	fmt.Println(s)
}

// slice is a array but wothout length
func ArraisSliceDeclaration() {
	names := [4]string{
		"jhon",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX" // modified the base array or underlying array
	fmt.Println(a, b)
	fmt.Println(names)
}

// El valor por defecto de un slice es nill
func SliceDefinition() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{{2, true}, {3, false}, {5, true}, {7, true}}

	fmt.Println(s)

}

func SliceLengthCapacity() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice("", s)

	// Slice the slice to give it zero length
	s = s[:0]
	printSlice("", s)

	// Extend its length.
	s = s[:4]
	printSlice("", s)

	// Drop its first two values.
	s = s[2:]
	printSlice("", s)
}

func SliceDefinitionByMake() {
	fmt.Println()
	a := make([]int, 5) //len(5) cap(5)
	printSlice("a", a)

	b := make([]int, 0, 5) //len(0) cap(5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c) //len(2) cap(5)

	d := c[2:5]
	printSlice("d", d) //len(3) cap (3)

}

func SliceTicTacToe() {
	fmt.Println()
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

}

func SliceAppend() {
	var s []int
	printSlice("S", s)

	s = append(s, 0) // va colocando en la position que tiene disponible si no se llena el slice lo incrementa
	printSlice("Appending 0", s)

	s = append(s, 1)
	printSlice("Appending 1", s)

	s = append(s, 2, 3, 4)
	printSlice("Appending more one than element", s)
}

func SliceIteration() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	/*

		Omiting algunos de los dos index o value
		for i := range pow {
			pow[i] = 1 << uint(i) // 2**i
		}

		for _, value := range pow {
			fmt.Printf("%d\n", value)
		}
	*/
}

func printSlice(s string, x []int) {
	fmt.Println()
	fmt.Printf("string %s len =%d cap=%d %v", s, len(x), cap(x), x)
}
