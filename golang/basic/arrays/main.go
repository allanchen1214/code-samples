package main

import "fmt"

func elemExists(s string, a []string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}
	return false
}

func main() {
	var empty []int
	fmt.Println(len(empty))

	var nums []int
	nums = append(nums, 203)
	nums = append(nums, 302)
	fmt.Println(nums)

	alphas := []string{"abc", "def", "ghi", "jkl"}
	alphas = append(alphas, "mno", "pqr", "stu")

	// Retrieve a single element from slice
	fmt.Println(alphas[3]) // jkl

	//Retrieve a sub slice of a slice
	alphas2 := alphas[2:5]
	fmt.Println(alphas2) // [ghi, jkl, mno]

	fmt.Println(elemExists("mno", alphas)) // true
	fmt.Println(elemExists("std", alphas)) // false
}
