package main

import "fmt"

func main() {
	ids := []int{33, 76, 54, 23, 11, 2}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index. Use the under score if you want to leave out the index.
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum: ", sum)

	// Range with map
	phoneNumbers := map[string]string{"Tim": "0707770077", "Tom": "0770007700", "Tam": "0730731234"}

	for k, v := range phoneNumbers {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range phoneNumbers {
		fmt.Printf("Name: %s\n", k)
	}
}
