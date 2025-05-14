package advance

import (
	"fmt"
)

func Pointers() {
	fmt.Println("Pointers")
	// Go has pointers. A pointer holds the memory address of a value.
	// The type *T is a pointer to a T value. Its zero value is nil.
	// var p *int
	// The & operator generates a pointer to its operand.
	// i := 42
	// p = &i
	// The * operator denotes the pointer's underlying value.
	// fmt.Println(*p) // read i through the pointer p
	// *p = 21         // set i through the pointer p
	// This is known as "dereferencing" or "indirecting".
	// Unlike C, Go has no pointer arithmetic.
	// Go has garbage collection, so it's safe to return a pointer to local variable.
	i := 42
	p := &i
	message := fmt.Sprintf("The value at memory location %d at pointer p is %d", p, *p) // 42
	fmt.Println(message)
	*p = 21
	fmt.Println(i) // 21

	// To access the field X of a struct when we have the struct pointer p
	// we could write (*p).X. However, that notation is cumbersome,
	// so the language permits us instead to write just p.X,
	// without the explicit dereference.
	type Vertex struct {
		X int
		Y int
	}
	v := Vertex{1, 2}
	pv := &v
	pv.X = 35
	fmt.Println(v) // {35 2}
}
