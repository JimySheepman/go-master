package main

import (
	"fmt"
	"strings"

	"github.com/fatih/camelcase"
)

func CamelCaseSplits(words []string) {
	for _, word := range words {
		splitted := camelcase.Split(word)
		str := strings.Join(splitted, "_")
		fmt.Print(strings.ToLower(str) + "_mock.go ")
	}
}

func CamelCaseSplit(word string) string {
	splitted := camelcase.Split(word)
	str := strings.Join(splitted, "_")
	fmt.Print()
	return strings.ToLower(str) + "_mock.go"
}
