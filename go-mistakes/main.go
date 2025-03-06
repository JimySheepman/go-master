package main

import "log"

func main() {
	for _, shade := range shades {
		if err := shade(); err != nil {
			log.Println(err)
		}
	}
}
