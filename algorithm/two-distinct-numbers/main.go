package main

import "fmt"

func main() {
	arr := []int{44, 10, 10, 22, 3, 33, 44, 100, 22, 0, 7, 3, 100, 0, 111}
	distinct_numbers := find_multi_distinct_numbers_sort(arr)
	fmt.Printf("distinct_numbers: %v\n", distinct_numbers) //  distinct_numbers: [7 33 111]

	arr = []int{44, 10, 10, 22, 3, 44, 100, 22, 0, 3, 100, 0, 111}
	n := find_one_distinct_numbers_XOR(arr)
	fmt.Printf("distinct_number: %d\n", n) // distinct_number: 111

	arr = []int{44, 10, 10, 22, 3, 33, 44, 100, 22, 0, 7, 3, 100, 0}
	two_numbers := find_two_distinct_numbers_XOR(arr)
	fmt.Printf("two_distinct_number: %v\n", two_numbers) // two_distinct_number: [33 7]

}

// 使用排序的办法，支持找到超过两个只出现一次的数字
func find_multi_distinct_numbers_sort(arr []int) []int {
	quicksort(arr, 0, len(arr)-1)
	result := []int{}
	for i := 0; i < len(arr); i++ {
		if i == 0 {
			if arr[0] != arr[1] {
				result = append(result, arr[0])
			}
			continue
		}
		if i == len(arr)-1 {
			if arr[len(arr)-1] != arr[len(arr)-2] {
				result = append(result, arr[len(arr)-1])
			}
			continue
		}
		if arr[i] != arr[i-1] && arr[i] != arr[i+1] {
			result = append(result, arr[i])
		}
	}
	return result
}

// 使用异或特性，仅支持找到一个只出现一次的数字
func find_one_distinct_numbers_XOR(arr []int) int {
	sum := arr[0]
	for i := 1; i < len(arr); i++ {
		sum = sum ^ arr[i]
	}
	return sum
}

// 使用异或特性，仅支持找到两个只出现一次的数字
func find_two_distinct_numbers_XOR(arr []int) []int {
	sum := arr[0]
	for i := 1; i < len(arr); i++ {
		sum = sum ^ arr[i]
	}

	// 异或结果中，右起找到第一个为1的位position，则代表两个只出现一次的数字里，其中一个position位为0，另一个为1
	// 这样能保证所有相同的数都被放到同一个数组，也能保证两个只出现一次的数分别放入到不同的数组
	position := get1position(sum)

	// 原数组拆分成两个数组，分别是所有position位为0的，以及所有position位为1的
	arr0 := []int{}
	arr1 := []int{}
	for _, item := range arr {
		if getpositionvalue(item, position) == 0 {
			arr0 = append(arr0, item)
		}
		if getpositionvalue(item, position) == 1 {
			arr1 = append(arr1, item)
		}
	}
	num1 := find_one_distinct_numbers_XOR(arr0)
	num2 := find_one_distinct_numbers_XOR(arr1)
	return []int{num1, num2}
}

func quicksort(arr []int, low, high int) {
	if low >= high {
		return
	}
	index := partition(arr, low, high)
	//fmt.Printf("recursion: %v\n", arr)
	quicksort(arr, low, index-1)
	quicksort(arr, index+1, high)
}

func partition(arr []int, low, high int) int {
	pivot := arr[low]

	for low < high {
		for low < high && arr[high] >= pivot {
			high--
		}
		arr[low] = arr[high]

		for low < high && arr[low] <= pivot {
			low++
		}
		arr[high] = arr[low]
		//fmt.Printf("partition-for: %v\n", arr)
	}
	arr[low] = pivot
	//fmt.Printf("partition-end: %v\n", arr)
	return low
}

func get1position(n int) int {
	position := 0
	for n&1 == 0 {
		n >>= 1
		position++
	}
	return position
}

func getpositionvalue(n, position int) int {
	i := position
	for i > 0 {
		n >>= 1
		i--
	}
	return n & 1
}
