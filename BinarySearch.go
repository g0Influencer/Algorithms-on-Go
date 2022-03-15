package main

import (
	"fmt"
	"sort"
)

func is_Sorted(list []int) bool {
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list)-1; j++ {
			if list[i] > list[j] {
				return false
			}
		}
	}
	return true
}

func binary_search(list []int, key int) interface{} {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := list[mid]
		if guess == key {
			return mid
		}
		if guess > key {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return nil
}

func main() {
	arr := []int{1, 5, 3, 2}
	if is_Sorted(arr) == false {
		sort.Ints(arr)
	}
	fmt.Println(binary_search(arr, 2))
}
