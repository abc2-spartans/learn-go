// The var statement declares a list of variables;
// as in function argument lists, the type is last.

// If an initializer is present, the type can be omitted;
// the variable will take the type of the initializer.

// Inside a function, the := short assignment statement can be used
// in place of a var declaration with implicit type.

// Outside a function, every statement begins with a keyword
// (var, func, and so on) and so the := construct is not available.

package basics

import "fmt"

var isTrue, isGood, isBad, j = true, true, false, 20

func Basics() {
	var i int
	// We can remove var keyword and use := to declare and
	// initialize a variable inside a function.
	var c, python, java = true, false, "no!"
	fmt.Println(i, isTrue, isGood, isBad)
	fmt.Println(i, j, c, python, java)
}

func Basics2() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	// var is removed and := is used to declare and initialize
	c, python, java := true, false, "no!"
	fmt.Println(i, isTrue, isGood, isBad)
	fmt.Println(i, j, c, python, java)
}
