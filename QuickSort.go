package main

import "fmt"

func quick_sort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	} else {
		pivot := arr[0]

		var less = []int{}
		var greater = []int{}

		for _, num := range arr[1:] {
			if num <= pivot {
				less = append(less, num)
			} else {
				greater = append(greater, num)
			}
		}
		less = append(quick_sort(less), pivot)
		greater = quick_sort(greater)
		return append(less, greater...)

	}
}

func main() {
	arr := []int{10, 2, 1, 8, 3}

	fmt.Println(quick_sort(arr))
}
