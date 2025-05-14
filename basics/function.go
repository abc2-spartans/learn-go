package basics

import (
	"fmt"
)

func Swap(x, y string) (string, string) {
	return y, x
}

type Rectangle struct {
	width, height float64
}

func area(v Rectangle) float64 {
	return v.width * v.height
}

func scale(r *Rectangle, f float64) {
	r.width *= f
	r.height *= f
}

func FunctionWithPointers() {
	r := Rectangle{3, 4}
	fmt.Println("Width: ", r.width, "Height: ", r.height)
	fmt.Println("Area before scaling: ", area(r))
	scale(&r, 10)
	fmt.Println("Area after scaling: ", area(r))
}
