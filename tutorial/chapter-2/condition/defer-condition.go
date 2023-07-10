package condition

import "fmt"

// La sentencia defer hace que una funcion sea llamada cuando la funcion principal retorna ya sea en su flujo normal o derrepente
func DeferCondition() {
	defer fmt.Println("world") // es executado cuando DeferCondition termina. En el flujo la funcion es evaluada pero no llamada hasta el final

	fmt.Println("hello")
}

func DeferConditionOntoStack() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i) // esto se almacena en un stack que es llamada en el order last-in-first-out
	}

	fmt.Println("done")
}
