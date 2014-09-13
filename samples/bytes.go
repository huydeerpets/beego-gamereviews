package main

import (
	"bytes"
	"fmt"
)

func main() {
	buffer := bytes.Buffer{}
	for i := 0; i < 1000; i++ {
		buffer.WriteString("a")
	}
	fmt.Println(buffer.String())
	fmt.Println([]byte(buffer.String()))
}
