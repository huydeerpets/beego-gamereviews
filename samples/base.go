package main

import (
	"fmt"
)

func init() {
	fmt.Println("init.")
}

func main() {
	defer fmt.Println("end.")

	str := "foo"
	num := 1
	arrays := []string{"a", "b", "c"}
	nums := []int{1, 2, 3}
	slice1 := []int{}
	slice2 := make([]int, 0, 10)
	map1 := map[string]int{"a": 1, "b": 2}

	fmt.Println("string:", str)
	fmt.Println("int:", num)
	fmt.Println("arrays:", arrays)
	fmt.Println("nums:", nums)
	fmt.Println("nums(1:3):", nums[1:3])

	switch {
	case num == 1:
		fmt.Println("num is 1")
	case num == 2:
		fmt.Println("num is 2")
	case num == 3:
		fmt.Println("num is 3")
	}

	for _, item := range nums {
		fmt.Println(item)
	}

	for pos, char := range "日本語" {
		fmt.Printf("character %c starts at byte position %d\n", char, pos)
	}

	for i := 0; i < 10; i++ {
		slice1 = append(slice1, i)
		fmt.Println(slice1, len(slice1), cap(slice1))
	}

	for i := 0; i < 10; i++ {
		slice2 = append(slice2, i)
		fmt.Println(slice2, len(slice2), cap(slice2))
	}

	fmt.Println("map1.a =", map1["a"])
	for k, v := range map1 {
		fmt.Println(k, v)
	}
}
