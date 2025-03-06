package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	body, err := os.ReadFile("./a.html")
	if err != nil {
		log.Fatalln("read file error")
	}

	var newList []string

	re := regexp.MustCompile(`.*alt=.*`)
	comments := re.FindAllString(string(body), -1)
	for _, comment := range comments {
		newList = append(newList, clearAnswers(comment))
	}

	reverse(newList)

	for i, l := range newList {
		fmt.Printf("%d: %s\n\n", i+1, l)
	}

}

func clearAnswers(str string) string {
	re := regexp.MustCompile(`"([^"]+)"`)
	matches := re.FindStringSubmatch(str)

	if len(matches) < 1 {
		return ""
	}

	return matches[1]
}

func reverse(ss []string) {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
}
