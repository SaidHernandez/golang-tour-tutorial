package functions // nombre del paquete

// How to creating a func that add two number, the function is exportable
// and it has two parameters
func Add(x int, y int) int {
	return x + y
}

// Cuando los parametros tienen el mismo tipo se puede hacer una funcion continua
func AddFuncContinued(x, y int) int {
	return x + y
}

// Se puede retornar multiples resultados en el return de una funcion esta funcion cambia el orden de
// de los parametros
func Swap(x, y string) (string, string) {
	return y, x
}

// Named return: Es una funciòn que retorna las variables declarasas en la firma de la
// de la función.
// Al final se le coloca return sin mas nada
func Split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // naked return
}
