package main

import (
	"fmt"
	"github.com/JimySheepman/go-master/go-algorithm/algorithm"
	"github.com/JimySheepman/go-master/go-algorithm/designpatterns"
	"github.com/JimySheepman/go-master/go-algorithm/errorhandling"
	"github.com/JimySheepman/go-master/go-algorithm/goroutine"
	"github.com/JimySheepman/go-master/go-algorithm/oop"
	"github.com/JimySheepman/go-master/go-algorithm/solids"
)

func main() {
	algorithm.PrintSortAlgorithm()
	fmt.Println()

	algorithm.PrintSearchAlgorithm()
	fmt.Println()

	algorithm.PrintGraphAlgorithm()
	fmt.Println()

	algorithm.PrintDataStructures()
	fmt.Println()

	algorithm.PrintTimeComplexity()
	fmt.Println()

	solids.PrintSolid()
	fmt.Println()

	oop.PrintOop()
	fmt.Println()

	designpatterns.PrintDesignPatterns()
	fmt.Println()

	goroutine.PrintGoroutine()
	fmt.Println()

	errorhandling.PrintErrorHandling()
	fmt.Println()
}
