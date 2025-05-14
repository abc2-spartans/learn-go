/*
*
Unlike in C, in Go assignment between items of different type requires an explicit conversion.
*
*/
package basics

import (
	"fmt"
	"math"
)

func TypeConv() {
	i := 42
	f := float64(i)
	u := uint(f)
	fmt.Println(i, f, u)
}

// Try removing the float64 or uint conversions in the example and see what happens.
func TypeConv2() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	// var z uint = f // This will give compilation error

	fmt.Println(x, y, z)
}
