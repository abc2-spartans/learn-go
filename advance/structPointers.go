package advance

// Struct fields can be accessed through a struct pointer.

// To access the field X of a struct when we have the struct pointer p
// we could write (*p).X. However, that notation is cumbersome,
// so the language permits us instead to write just p.X,
// without the explicit dereference.

import (
	"fmt"
)

func StructPointers() {
	type Vertex struct {
		X int
		Y int
	}
	v := Vertex{1, 2}
	pv := &v
	pv.X = 35
	fmt.Println(v) // {35 2}
}
