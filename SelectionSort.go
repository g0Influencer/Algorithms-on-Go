package main

import "fmt"

func findSmallest(arr []int) int {
	smallest := arr[0]
	smallest_index := 0

	for i := 1; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallest_index = i
		}
	}
	return smallest_index
}

func selection_sort(arr []int) []int {
	newArr := make([]int, len(arr))

	for i := range arr {
		smallest := findSmallest(arr)
		newArr[i] = arr[smallest]
		arr = append(arr[:smallest], arr[smallest+1:]...)
	}
	return newArr
}

func main() {
	arr := []int{5, 10, 1, 3, 8}
	fmt.Println(selection_sort(arr))
}
