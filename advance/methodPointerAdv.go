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

	"github.com/gauravsinghaec/learn-go/basics"
)

func MethodPointers() {
	// Go automatically handles conversion between values and pointers for method calls.
	basics.Divider("Indirection 1")
	// Methods with value receivers take either a value or a pointer as the receiver when they are called:

	r := Rect{width: 10, height: 5}
	fmt.Println("Width: ", r.width, "Height: ", r.height)
	area := r.area()   // This is equivalent to (&r).area()
	perim := r.perim() //	This is equivalent to (&r).perim()

	fmt.Println("Area: ", area)
	fmt.Println("Perimeter: ", perim)

	basics.Divider("Indirection 2")
	// Methods with pointer receivers take either a value or a pointer as the receiver when they are called:

	// Methods can be declared with pointer receivers.
	// If the receiver is a value, Go will automatically take the address of the value.
	// If the receiver is a pointer, Go will use the receiver as is.

	r.scale(10) // This is equivalent to (&r).scale(10)
	// Go interprets the statement r.Scale(10) as (&r).Scale(10) since the Scale method has a pointer receiver.

	fmt.Println("Area with pointers: ", r.areaPointer())
}
