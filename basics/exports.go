// In Go, a name is exported if it begins with a capital letter.
// For example, Pizza is an exported name, as is Pi,
// which is exported from the math package.
package basics

import "fmt"

func Add(x int, y int) int {
	return x + y
}

// When two or more consecutive named function parameters share a type,
// you can omit the type from all but the last. x int, y int -> x , y int

func Subtract(x, y int) int {
	return x - y
}

func Divider(topic string) {
	message := fmt.Sprintf("\n-------------- %s --------------\n", topic)
	fmt.Println(message)
}
