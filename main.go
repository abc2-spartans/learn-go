package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/gauravsinghaec/learn-go/advance"
	"github.com/gauravsinghaec/learn-go/basics"
)

// func split(sum int) (x, y int) {
// 	x = sum * 4 / 9
// 	y = sum - x
// 	return
// }

func main() {
	basics.Divider("factored imports")
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println("My favorite number is", rand.Intn(100))

	basics.Divider("Exports")
	a := basics.Add(4, 5)
	b := basics.Subtract(14, 5)
	fmt.Println(a, b)
	basics.Divider("Functions")
	x, y := basics.Swap("hello", "world")
	fmt.Println(x, y)
	basics.FunctionWithPointers()

	basics.Divider("Naked Return or function with no return value")
	s, t := basics.Split(17)
	fmt.Println(s, t)

	basics.Divider("Variables")
	basics.Basics()
	basics.Basics2()

	basics.Divider("Templates")
	basics.TemplateString()

	basics.Divider("Templates with text package")
	basics.TemplateWithTextPackage()

	basics.Divider("Templates with html package")
	basics.TemplateWithHTMLPackage()

	basics.Divider("Data Types")
	basics.DataTypes()

	basics.Divider("Data Type Conversion")
	basics.TypeConv2()

	basics.Divider("Type Inference")
	basics.TypeInf()

	basics.Divider("Const Keyword")
	basics.Const()

	basics.Divider("Numeric constants")
	basics.NumericConst()

	basics.Divider("For Loop")
	basics.ForLoop()

	basics.Divider("If Else")
	fmt.Println(basics.IfElse(17))

	basics.Divider("If Else Short")
	fmt.Println(basics.IfElseShort(3, 3, 20))

	basics.Divider("Switch")
	basics.Switch()

	basics.Divider("Stacking Defers")
	advance.DeferredFunc()

	basics.Divider("Pointers")
	advance.Pointers()

	basics.Divider("Structs and Struct Literals")
	advance.StructPointers()

	basics.Divider("Arrays and Array Literal")
	advance.Array()

	basics.Divider("Slices")
	advance.Slices()

	basics.Divider("Maps")
	advance.Maps()

	basics.Divider("Function Values")
	advance.FuncAsValue()

	basics.Divider("Function Closures")
	advance.Closures()

	basics.Divider("Methods")
	advance.Methods()

	basics.Divider("Method Pointers Indirections")
	advance.MethodPointers()

	basics.Divider("Interfaces")
	advance.InterfaceBasic()
}
