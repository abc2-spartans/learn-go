/*
*	Go has the following basic data types:

bool
string
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr
byte // alias for uint8
rune // alias for int32

	// represents a Unicode code point

float32 float64
complex64 complex128

*
*/
package basics

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe bool = false
	// Create a number by shifting a 1 bit left 64 places.
	// In other words, the binary number that is 1 followed by 64 zeroes.
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func DataTypes() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
