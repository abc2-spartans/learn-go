package advance

import "fmt"

// An interface type is defined as a set of method signatures.
// A value of interface type can hold any value that implements those methods.

// Define an interface
type Shape interface {
	Area() float64
}

// Define a struct
type Circle struct {
	radius float64
}

// Implement the interface method
func (c *Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

// Circle implements Shape because it has an Area() method.

func Interfaces() {
	var s Shape
	c := Circle{radius: 5}
	s = &c // *Circle implements Shape

	// Circle (the value type) doesn't implement Area because the Area method is defined only on *Circle (the pointer type).
	// s1 = c will throw an error.

	// var s1 Shape
	// s1 = c                // Circle does not implement Shape
	fmt.Println(s.Area()) // Output: 78.5
}
