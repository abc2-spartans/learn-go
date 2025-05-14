package basics

import "fmt"

// Go has only one looping construct, the for loop.
func ForLoop() {
	sum := 0
	// The init and post statements are optional.
	for i := 0; i < 10; i++ {
		sum += i
	}
	message := fmt.Sprintf("The sum from 0 to 9 is = %d", sum)
	fmt.Println(message)

	s := 1
	// It is like while loop in other languages.
	for s < 10 {
		s += s
	}
	fmt.Println(s)

	// If you omit the loop condition it loops forever,
	// so an infinite loop is compactly expressed.
	// e.g for {}
	Divider("For Loop with range")
	// The range form of the for loop iterates over a slice or map.
	// When ranging over a slice, two values are returned for each iteration.
	// The first is the index, and the second is a copy of the element at that index.

	// You can skip the index or value by assigning to _.
	// If you only want the index, drop the ", value" entirely.
	// for i, _ := range pow
	// for _, value := range pow
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow2 := make([]int, 10)
	for i := range pow2 {
		pow2[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow2 {
		fmt.Printf("%d\n", value)
	}

}
