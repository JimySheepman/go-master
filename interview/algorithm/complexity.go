package algorithm

import "fmt"

// constantTimeAlgorithm - O(1)
func constantTimeAlgorithm(arr []int) int {
	return arr[0]
}

// linearTimeAlgorithm - O(n)
func linearTimeAlgorithm(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

// quadraticTimeAlgorithm - O(n²)
func quadraticTimeAlgorithm(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Println(i, j)
		}
	}
}

// logarithmicTimeAlgorithm - O(log n)
func logarithmicTimeAlgorithm(n int) int {
	result := 1
	for i := 1; i <= n; i *= 2 {
		result *= i
	}
	return result
}

// exponentialTimeAlgorithm - O(2ⁿ)
func exponentialTimeAlgorithm(n int) int {
	if n <= 1 {
		return n
	}
	return exponentialTimeAlgorithm(n-1) + exponentialTimeAlgorithm(n-2) // Fibonacci sayısını hesaplama
}

func PrintTimeComplexity() {
	arr := []int{1, 2, 3, 4, 5}
	n := len(arr)

	constantTimeAlgorithm(arr)
	linearTimeAlgorithm(arr)
	quadraticTimeAlgorithm(n)
	logarithmicTimeAlgorithm(n)
	exponentialTimeAlgorithm(n)

}
