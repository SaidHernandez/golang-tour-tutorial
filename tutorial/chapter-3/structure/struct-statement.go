package structure

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      //X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func StructDefinition() {
	fmt.Println("Struct: ", Vertex{1, 2})
}

func StructAccess() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println("acceding by dot", v.X)
}

func StructAccessByPointer() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println("Struct access by pointer", v)
}
