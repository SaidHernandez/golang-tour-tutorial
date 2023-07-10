package functionvalue

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// Func closures
func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

func FuncValue() {
	fmt.Println("Function value")
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func FunctionClosures() {
	fmt.Println("Function Closures")
	post, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			post(i),
			neg(-2*i),
		)

	}

}
