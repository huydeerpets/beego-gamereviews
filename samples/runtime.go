package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("GOROOT:", runtime.GOROOT())
	fmt.Println("NumCPU:", runtime.NumCPU())
	fmt.Println("NumCgoCall:", runtime.NumCgoCall())
	fmt.Println("NumGoroutine:", runtime.NumGoroutine())
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(runtime.NumCPU()))

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println("Memory", m.Alloc, m.StackInuse, m.HeapAlloc)
}
