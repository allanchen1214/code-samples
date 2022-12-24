package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Convert String to Integer
	ss := "42"
	ans, err := strconv.Atoi(ss)
	if err != nil {
		fmt.Println("Error convert string to integer", err)
	}
	fmt.Printf("%d\n", ans)

	// Convert Integer to a String
	ss2 := strconv.Itoa(42)
	fmt.Printf("%s\n", ss2)

	// Convert Integer to a Float
	i := 42   // int type
	y := 42.0 // float type
	x := float64(i)
	fmt.Printf("y:%f, x:%f\n", y, x)

	// Convert a String to a Float
	pi := "3.14159"
	m, err := strconv.ParseFloat(pi, 64)
	if err != nil {
		fmt.Println("Error convert string to float", err)
	}
	fmt.Printf("%f\n", m)
}
