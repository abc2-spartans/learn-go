package advance

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func FuncAsValue() {
	// Functions are values too. They can be passed around just like other values.
	// Function values may be used as function arguments and return values.
	// The zero value of a function type is nil.
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println("Hyptenuse of 5 and 12: ", hypot(5, 12))

	fmt.Println("Compute hypotenuse using custom compute func: ", compute(hypot))
	fmt.Println("Compute pow using math.Pow: ", compute(math.Pow))
}
