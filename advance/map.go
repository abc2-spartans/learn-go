package advance

import (
	"fmt"

	"github.com/gauravsinghaec/learn-go/basics"
)

// In Go, a map (map[keyType]valueType) is a built-in data type that represents a key-value store,
// similar to dictionaries in Python or objects in JavaScript.
// The zero value of a map is nil. A nil map has no keys, nor can keys be added.

func Maps() {
	var grades map[string]int
	fmt.Println(grades)
	// grades["Math"] = 70 // This will throw an error because grades is nil.

	basics.Divider("Map literals")
	// Map literals are like struct literals, but the keys are required.
	// If the top-level type is just a type name, you can omit it from the elements of the literal.
	grades2 := map[string]int{
		"Math":    95,
		"Science": 80,
		"English": 70,
	}
	fmt.Println("Grades of math grades2: ", grades2["Math"])

	// If a key doesn't exist, Go returns the zero value of the value type.
	fmt.Println("Grades of history in grades2: ", grades2["History"])

	basics.Divider("Creating a map using make")
	// If you don't know the values beforehand, you can use make.
	// make(map[keyType]valueType) initializes an empty map.
	grades3 := make(map[string]int)
	grades3["Math"] = 90

	fmt.Println("Grades of math in grades3: ", grades3["Math"])

	basics.Divider("Checking if a key exists in a map")
	// This avoids confusion when a missing key returns a default zero value.
	score, exists := grades3["History"]
	if exists {
		fmt.Println("History's marks:", score) // 0
	} else {
		fmt.Println("History not found")
	}
	score1, exists1 := grades2["Science"]
	if exists1 {
		fmt.Println("Science's marks:", score1) // 80
	} else {
		fmt.Println("Science not found")
	}

	basics.Divider("Length of a map")
	// The number of items in a map is obtained using the len function.
	fmt.Println("Length of grades2: ", len(grades2))

	basics.Divider("Deleting a key from a map")
	// To delete a key from a map, use the delete function.
	delete(grades2, "Math")
	fmt.Println("Grades2 after deleting Math: ", grades2)

	basics.Divider("Iterating over a map")
	// The first is the key, and the second is the value.
	for subject, marks := range grades2 {
		fmt.Printf("Subject: %s, Marks: %d\n", subject, marks)
	}

}
