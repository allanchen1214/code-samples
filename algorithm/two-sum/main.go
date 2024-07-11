package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7777, 7779, 222}
	sum := 15556
	sum = 7
	result := two_sum_map(arr, sum)
	fmt.Printf("result: %v\n", result)
}

func two_sum_map(arr []int, sum int) (result [][]int) {
	m := make(map[int]int)
	for i, v := range arr {
		m[v] = i
	}
	tmp := make(map[int]int)
	for i, v := range arr {
		if _, ok := tmp[i]; ok {
			continue
		}
		if j, ok := m[sum-v]; ok {
			tmp[j] = 1
			result = append(result, []int{i, j})
		}
	}
	return
}
