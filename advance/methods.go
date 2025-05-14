package advance

// Methods are functions with a special receiver argument before the function name.

// func (receiver Type) methodName(params) returnType {
//     // Method body
// }

// receiver → The variable representing the instance (like this in other languages).
// Type → The struct or named type the method belongs to.
// methodName → The name of the method.
// params → The parameters the method accepts.
// returnType → The return type of the method.

import (
	"fmt"
	"math"

	"github.com/gauravsinghaec/learn-go/basics"
)

type Rect struct {
	width, height float64
}

func (r Rect) area() float64 {
	return r.width * r.height
}

func (r Rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// You can only declare a method with a receiver whose type is defined in the same package as the method.
// You cannot declare a method with a receiver whose type is defined in another package
// (which includes the built-in types such as int).

type MyFloat float64

func (f MyFloat) abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Methods can be declared with pointer receivers.
func (r *Rect) areaPointer() float64 {
	return r.width * r.height
}

func (r *Rect) scale(factor float64) {
	r.width *= factor
	r.height *= factor
}

func Methods() {
	basics.Divider("Value Receivers")
	r := Rect{width: 10, height: 5}
	fmt.Println("Width: ", r.width, "Height: ", r.height)
	area := r.area()
	perim := r.perim()
	fmt.Println("Area: ", area)
	fmt.Println("Perimeter: ", perim)

	f := MyFloat(-math.Sqrt2)
	fmt.Println("Absolute value of -√2: ", f.abs())

	basics.Divider("Pointer Receivers")
	fmt.Println("Area with pointers: ", r.areaPointer())
}
