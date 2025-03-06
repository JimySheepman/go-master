package algorithm

import "fmt"

type search func(arr []int, x int) int

var searchAlgorithms = []search{
	linearSearch,
	binarySearch,
}

var searchAlgorithmsName = map[int]string{
	0: "Linear Search",
	1: "Binary Search",
}

func linearSearch(arr []int, x int) int {
	for i, v := range arr {
		if v == x {
			return i
		}
	}
	return -1
}

func binarySearch(arr []int, x int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == x {
			return mid
		}
		if arr[mid] < x {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func PrintSearchAlgorithm() {
	for i, searchFunc := range searchAlgorithms {
		arr := []int{2, 3, 4, 10, 40}
		x := 10
		fmt.Println("Algorithm name:", searchAlgorithmsName[i])
		result := searchFunc(arr, x)
		if result != -1 {
			fmt.Printf("Element %d is present at index %d\n", x, result)
		} else {
			fmt.Printf("Element %d is not present in array\n", x)
		}
	}
}
