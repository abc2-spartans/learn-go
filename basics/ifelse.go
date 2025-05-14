package basics

import (
	"fmt"
	"math"
)

func IfElse(x float64) string {
	if x < 0 {
		return IfElse(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func IfElseShort(x, n, lim float64) float64 {
	// Like for, the if statement can start with a
	// short statement to execute before the condition.
	// Variables declared by the statement are only in scope until the end of the if.

	// Variables declared inside an if short statement
	// are also available inside any of the else blocks.
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}
