package main

import (
	"fmt"
)

func main() {
	// Define map
	emails := make(map[string]string)

	// Assign kv
	emails["Tim"] = "tim@gmail.com"
	emails["Tom"] = "tom@gmail.com"
	emails["Tam"] = "tam@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Tim"])

	// Delete key value
	delete(emails, "Tam")
	fmt.Println(emails)

	// Decalare map and add kv
	phoneNumbers := map[string]string{"Tim": "0707770077", "Tom": "0770007700", "Tam": "0730731234"}
	fmt.Println(phoneNumbers) 
}
