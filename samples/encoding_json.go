package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Str   string
	Num   int
	Array []string
	Map   map[string]int
}

func main() {
	data := Data{Str: "foo", Num: 123, Array: []string{"a", "b"}, Map: map[string]int{"a": 1, "b": 2}}
	content, _ := json.Marshal(&data)
	fmt.Println(string(content))

	data.Str = "bar"
	fmt.Println(data.Str)

	json.Unmarshal(content, &data)
	fmt.Println(data.Str)
	fmt.Println(data.Num)
	fmt.Println(data.Array)
	fmt.Println(data.Map)
}
