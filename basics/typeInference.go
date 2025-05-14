/*
*
When declaring a variable without specifying an explicit type
(either by using the := syntax or var = expression syntax),
the variable's type is inferred from the value on the right hand side.
*
*/
package basics

import (
	"fmt"
)

func TypeInf() {
	var j int
	k := j            // j is an int
	i := 42           // int
	f := 3.142        // float64
	g := 0.867 + 0.5i // complex128
	fmt.Printf("k is of type %T\n", k)
	fmt.Printf("i is of type %T\n", i)
	fmt.Printf("f is of type %T\n", f)
	fmt.Printf("g is of type %T\n", g)
}
