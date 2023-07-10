package interfaces

import "fmt"

func describeV2(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func CallEmptyInterface() {
	var ife interface{}
	describeV2(ife)

	ife = 422
	describeV2(ife)

	ife = "hello"
	describeV2(ife)

}
