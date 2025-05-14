package advance

// An array's length is part of its type, so arrays cannot be resized. This seems limiting,
// but don't worry; Go provides a convenient way of working with arrays.
// We can overcome this limitation by using slices's literals.
import "fmt"

func Array() {
	var a = [2]int{1, 2}
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var strArr [2]string
	strArr[0] = "Hello"
	strArr[1] = "World"
	fmt.Println(strArr, strArr[1])
	fmt.Println(a[0], a[1])
	fmt.Println(primes)
}
