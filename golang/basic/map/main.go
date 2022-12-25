package main

import "fmt"

func main() {
	// Make a Map
	m := make(map[string]string)
	m["c"] = "Cyan"
	m["y"] = "Yellow"
	m["m"] = "Magenta"
	m["k"] = "Black"

	// Init
	var m2 = map[string]string{
		"c": "Cyan",
		"y": "Yellow",
		"m": "Magenta",
		"k": "Black",
	}

	// Iterate over
	for k, v := range m2 {
		fmt.Printf("Key: %s, Value: %s\n", k, v)
	}

	// Retrieve a Map Item
	k := m["k"]
	fmt.Printf("%s\n", k)

	// Delete a Map Item
	delete(m, "k")

	// Test Item Exists
	_, ok := m["k"]
	if !ok {
		fmt.Println("Key:k not exists")
	}
}
