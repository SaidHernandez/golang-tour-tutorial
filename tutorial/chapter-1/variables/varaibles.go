package variables

import (
	"fmt"
	"math"
	"math/cmplx"
)

//Typos de variables
// bool =  valor de defecto false
// string = valor por defecto ''
// int int8 int16 int32 int64
// uint uint8 uint16 uint32 uint64 uintptr Solo son utilizados para almacenar numeros positivos, por ende su capacidad siempre comienza desde cero.  uint8 variable can store values from 0 to 255, while an int8 variable can store values from -128 to 127
// byte = uint8
//rune = int32
// float32 float64
// complex64 complex128

//Valor por defecto
//0 for numeric
//false para boleanos
//"" string

// Declaración de variables a nivel de paquete
// es una lista de declaración de tipo bool
var c, python, java bool // el valor de iniciación es false
var i, j int = 1, 2      // incialización de variables en orden de lo teclado.

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1            // Es la maxima asiganación numerica para un no unsigned de 64-bit integer, esto sirve para revisar que no se exceda la cantidad numerica maxima.
	z      complex128 = cmplx.Sqrt(-5 + 12i) // Representa numeros complejos con floating-point de una presición de 128 bits.. con real par y imaginaria parte
)

func DeclareVariables() {
	var i int // variables a nivel de función
	fmt.Println(i, c, python, java)
}

func DeclareVariablesInitializers() {
	var c, python, java = true, false, "no!" // Cuando se coloca la palavra var al incio se declara la variable, se puede omitir si se coloca :=, el tipe de la variables lo toma segun lo asignado.
	fmt.Println(i, j, c, python, java)
}

func DeclareVariablesShortInitializers() {
	k := 3 // iniciación rapida, por defecto se toma el tipo del valor asignado
	fmt.Println(k)
}

func ConvertVariables() {
	var x, y int = 3, 6
	var f float64 = math.Sqrt(float64((x * x) + (y * y)))
	fmt.Println(f)
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

func InferenceTypeVariables() {
	i := 42           // int
	f := 3.142        // float64
	g := 0.867 + 0.5i // complex128

	fmt.Println(i, f, g)
	fmt.Printf("g is of type %T\n", g) //%T imprime el tipo de la Variables
}

func ConstVariables() {
	const World = "Colombia" // declaración de un string
	fmt.Println("Hola", World)
}
