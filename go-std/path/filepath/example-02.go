package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println("On Unix:")
	fmt.Println(filepath.Dir("/foo/bar/baz.js"))
	fmt.Println(filepath.Dir("/foo/bar/baz"))
	fmt.Println(filepath.Dir("/foo/bar/baz/"))
	fmt.Println(filepath.Dir("/dirty//path///"))
	fmt.Println(filepath.Dir("dev.txt"))
	fmt.Println(filepath.Dir("../todo.txt"))
	fmt.Println(filepath.Dir(".."))
	fmt.Println(filepath.Dir("."))
	fmt.Println(filepath.Dir("/"))
	fmt.Println(filepath.Dir(""))
}
