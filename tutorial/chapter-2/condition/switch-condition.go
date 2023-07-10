package condition

import (
	"fmt"
	"runtime"
	"time"
)

/*
Los Switch son muy parecidos a otros lenguaje pero con la particularidad que:

 1. No necesarimente debe ser constante la expreció a evaliuar

 2. Los break no es una expresión obligatoria.

 3. Las condicion no deben ser siempre numericos.

    Se evalua siempre de arriba a bajo hasta encontrar la condición verdadera ¡
*/
func SwitchConditional() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println(("Linux. "))
	default:
		fmt.Printf("%s \n", os)

	}
}

func SwitchConditionalDay() {
	fmt.Println("When is Satuday")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In Two Days")
	default:
		fmt.Println("Too far way")

	}
}

// La condición a evaluar puede ser siempre verdadera usado para escribir un if-then-else demasiado extenso
func SwitchTrueth() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening")

	}

}
