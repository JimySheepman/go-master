package creational

import (
	"fmt"
	"sync"
)

type singleton struct {
	data string
}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{data: "Singleton Instance"}
	})
	return instance
}

func Singleton() {
	s1 := GetInstance()
	fmt.Println(s1.data)

	s2 := GetInstance()
	fmt.Println(s2.data)

	fmt.Println(s1 == s2)
}
