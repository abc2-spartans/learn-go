// An array has a fixed size. A slice, on the other hand, is a dynamically-sized,
// flexible view into the elements of an array. In practice, slices are much more common than arrays.

// A slice is formed by specifying two indices, a low and high bound, separated by a colon:
// a[low : high]
// This selects a half-open range which includes the first element, but excludes the last one.

// When slicing, you may omit the high or low bounds to use their defaults instead.
// For the array var a [10]int
// these slice expressions are equivalent:
// a[0:10]
// a[:10]
// a[0:]
// a[:]

package advance

import (
	"fmt"
	"strings"

	"github.com/gauravsinghaec/learn-go/basics"
)

func Slices() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// A slice does not store any data, it just describes a section of an underlying array.
	// Changing the elements of a slice modifies the corresponding elements of its underlying array.
	// The type []T is a slice with elements of type T
	// The following expression creates a slice which includes elements 1 through 3 of a:
	fmt.Println("Array: ", primes)
	var s []int = primes[1:4]
	fmt.Println("Slice s:", s)
	t := primes[3:]
	// A slice does not store any data, it just describes a section of an underlying array.
	// Changing the elements of a slice modifies the corresponding elements of its underlying array.
	// Other slices that share the same underlying array will see those changes.
	primes[3] = 100
	fmt.Println("Slice values after changing the 4th element of the underlying array: ")

	fmt.Println("Slice s: ", s, "Slice t: ", t)

	basics.Divider("Slices literals")
	// Slice literals
	// A slice literal is like an array literal without the length.
	primesSliceLiterals := []int{2, 3, 5, 7, 11, 13}

	fmt.Println(primesSliceLiterals)

	boolSliceLiterals := []bool{true, false, true, true, false, true}
	fmt.Println(boolSliceLiterals)

	structSliceLiterals := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(structSliceLiterals)

	basics.Divider("capacity and length of a slice")
	// The length of a slice is the number of elements it contains.
	// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

	// The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).

	fmt.Printf("Length of s: %d, Capacity of s: %d\n", len(s), cap(s))

	basics.Divider("Nil slices")
	// The zero value of a slice is nil. A nil slice has a length and capacity of 0 and has no underlying array.
	var mySlice []int
	fmt.Println("Nil slice: ", mySlice)
	fmt.Printf("Length of mySlice: %d, Capacity of mySlice: %d\n", len(mySlice), cap(mySlice))

	basics.Divider("Creating a slice with make")
	// Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
	// The make function allocates a zeroed array and returns a slice that refers to that array:
	// a := make([]int, 5)  // len(a)=5
	// To specify a capacity, pass a third argument to make:
	// b := make([]int, 3, 5) // len(b)=3, cap(b)=5
	nums := make([]int, 5, 10)
	fmt.Println("Slice nums: ", nums)
	fmt.Printf("Length of nums: %d, Capacity of nums: %d\n", len(nums), cap(nums))

	basics.Divider("Nested slice")
	// Slices can contain any type, including other slices.
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	fmt.Println("Tic-tac-toe board: ", board)
	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	fmt.Println("Tic-tac-toe board after players take turns: ")
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	basics.Divider("Appending to a slice")
	// It is common to append new elements to a slice, and so Go provides a built-in append function.
	// func append(s []T, vs ...T) []T
	// The first parameter s of append is a slice of type T, and the rest are T values to append to the slice.
	// The resulting value of append is a slice containing all the elements of the original slice plus the provided values.
	// The returned slice will point to the newly allocated array.
	var mySlice2 []int
	fmt.Println("Nil slice: ", mySlice2)
	fmt.Printf("Length of mySlice2: %d, Capacity of mySlice2: %d\n", len(mySlice2), cap(mySlice2))
	mySlice2 = append(mySlice2, 0, 5, 1)
	fmt.Println("Slice after appending: ", mySlice2)
	fmt.Printf("Length of mySlice2: %d, Capacity of mySlice2: %d\n", len(mySlice2), cap(mySlice2))
}
