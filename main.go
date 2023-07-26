package main

import ( // Factored import statement
	"fmt"
	"golang-tour-tutorial/tutorial/chapter-1/constants"
	"golang-tour-tutorial/tutorial/chapter-1/functions"
	"golang-tour-tutorial/tutorial/chapter-1/variables"
	"golang-tour-tutorial/tutorial/chapter-2/condition"
	"golang-tour-tutorial/tutorial/chapter-2/iterables"
	"golang-tour-tutorial/tutorial/chapter-3/arraies"
	functionvalue "golang-tour-tutorial/tutorial/chapter-3/function-value"
	"golang-tour-tutorial/tutorial/chapter-3/maps"
	"golang-tour-tutorial/tutorial/chapter-3/pointers"
	"golang-tour-tutorial/tutorial/chapter-3/structure"
	"golang-tour-tutorial/tutorial/chapter-4/errors"
	"golang-tour-tutorial/tutorial/chapter-4/images"
	"golang-tour-tutorial/tutorial/chapter-4/interfaces"
	"golang-tour-tutorial/tutorial/chapter-4/methods"
	typesassert "golang-tour-tutorial/tutorial/chapter-4/types-assert"
	"golang-tour-tutorial/tutorial/chapter-5/generics"
	"golang-tour-tutorial/tutorial/chapter-6/channels"
	"golang-tour-tutorial/tutorial/chapter-6/goroutines"
	"math"
	"math/rand"
	"time"
)

//other way no so god
// import "fmt"
// import "time"

func main() {
	fmt.Println("Hello again!!")
	fmt.Println("The time is", time.Now())
	fmt.Println("My favority number is", rand.Intn(10))      // Genera un numero aletorio
	fmt.Printf("Now you have %g problems. \n", math.Sqrt(7)) //Square root of x = 7
	//fmt.Println(math.pi) Lanza error cuz no exportable error
	fmt.Println(math.Pi)                            //accediendo a un constante exportable
	fmt.Println(functions.Add(42, 13))              // llamando a una funcion en otro paquete
	fmt.Println(functions.AddFuncContinued(45, 50)) // llamando a una funcion continua
	fmt.Println(functions.Swap("World", "Hello"))

	// En go esta es una de las maneras de declarar una variable
	// sin declarar el tipado
	a, b := functions.Swap("Caballo", "rojo")
	fmt.Println(functions.Swap(a, b))
	fmt.Println(functions.Split(17))

	variables.DeclareVariables()
	variables.DeclareVariablesInitializers()
	variables.DeclareVariablesShortInitializers()
	variables.ConvertVariables()
	variables.InferenceTypeVariables()
	variables.ConstVariables()

	//fmt.Println("Big %v", constants.Big) //  overflows

	fmt.Println(constants.NeedInt(constants.Small))
	fmt.Println(constants.NeedFloat(constants.Small))
	fmt.Println(constants.NeedFloat(constants.Big))

	// TODO: dejar de ejemplo u codigo con errores

	/* Other word to print a  Big = 1 << 100
	Big := new(big.Int)
	Big.SetBit(Big, 100, 1)
	fmt.Println(Big)
	*/

	//Flow Control Stetements:
	// For
	iterables.IterableForBasic(10)
	iterables.IterableForBasic(20)

	//IF Condition
	fmt.Println(condition.Sqrt(2), condition.Sqrt(-4))
	fmt.Println(condition.Pow(3, 2, 10))
	fmt.Println(condition.Pow(3, 2, 20))

	//SWITCH condition
	condition.SwitchConditional()
	condition.SwitchConditionalDay()
	condition.SwitchTrueth()

	//defer
	condition.DeferCondition()
	condition.DeferConditionOntoStack()

	//pointer
	pointers.PointerDefinition()

	//struct
	structure.StructDefinition()
	structure.StructAccess()
	structure.StructAccessByPointer()

	//array
	arraies.ArraiesDeclaration()
	arraies.ArraisSliceDeclaration()
	arraies.SliceDefinition()
	arraies.SliceLengthCapacity()
	arraies.SliceDefinitionByMake()
	arraies.SliceTicTacToe()
	arraies.SliceAppend()
	arraies.SliceIteration()

	//array
	maps.MapsDeclaration()
	maps.MapsDeclarationUnderlying()
	maps.MapsMutating()

	//func
	functionvalue.FuncValue()
	functionvalue.FunctionClosures()

	//method
	methods.CallTypeMethod()
	methods.CallPointerMethod()

	//interfaces
	interfaces.UseInterface()
	interfaces.CallInterfaceValue()
	interfaces.CallInterfaceValueNill()
	interfaces.CallEmptyInterface()
	interfaces.CallStringer()

	//types assertion con interfaces
	typesassert.CallTypeAsser()
	typesassert.CallDoInterface()

	//errors
	errors.CallErrors()

	//images
	images.CallImages()

	//index
	generics.CallIndex()

	//goroutines
	goroutines.CallSay()

	//channel
	channels.CallSum()
	channels.BufferedChannels()
	channels.CallFibonacci()
}
