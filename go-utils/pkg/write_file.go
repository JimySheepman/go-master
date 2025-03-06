package main

import (
	"fmt"
	"log"
	"os"
)

func WriterFile(path, str, info string) {
	f, err := os.Create(path + str)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(info)
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("done")
}
