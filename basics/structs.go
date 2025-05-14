package basics

import "fmt"

type Vertex struct {
	X int
	Y int
}

func Structs() {
	v := Vertex{1, 2}
	v.X = 4 // Struct fields are set and accessed using a dot.
	fmt.Println(v.X, v.Y)
	// Struct literals
	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		p  = &Vertex{1, 2} // has type *Vertex
	)
	fmt.Println(v1, p, v2, v3)
}
