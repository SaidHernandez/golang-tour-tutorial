package condition

import (
	"fmt"
	"math"
)

func Sqrt(x float64) string {
	if x < 0 { // no necesita parentesis
		return Sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

/*
se pueden declarar variables y usarla alli mismo. Pero, Esto solo es en el scope del if, incluso en cualquier else
*/
func Pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else { //usando v en else scope
		fmt.Printf("%g >= %g\n", v, lim)
		// can't use v here, though
	}
	//return v error: esta variable solo esta en el scope del if
	return lim
}
