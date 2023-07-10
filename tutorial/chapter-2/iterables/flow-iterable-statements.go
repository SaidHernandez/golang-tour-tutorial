package iterables

import "fmt"

/*
En Golan solo se tiene un solo iterable flow statament, y puede ser declarado de muchas maneras.

 1. for i := 0; i < 10; i++ init  statement, condition expressio, post statement
    The init and post statements are optional.

 2. for i, v := range {{array}}

 3. for {}
*/
func IterableForBasic(n int) {
	sum := 0
	for i := 0; i < n; i++ {
		sum += i
	}
	fmt.Println(sum)
}

/*
*

	Optional statments
	for  {
		//infitity loop
	}
*/
func IterableForOptionalDefinition(n int) {
	sum := 1
	for sum < n {
		sum += sum
	}

	fmt.Println(sum)

}
