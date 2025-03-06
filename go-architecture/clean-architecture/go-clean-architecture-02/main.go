package main

import (
	"go-transactions/infrastructure"
)

func main() {
	infrastructure.NewHTTPServer().Start()
}
