package main

import (
	"fmt"
)

func Regist(f func(msg string)) {
	f("hoge")
}

func main() {
	Regist(func(msg string) {
		fmt.Println(msg)
	})
}
