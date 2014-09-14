package main

import (
	"fmt"
	"sync"
)

type Result struct {
	Message string
}

func UseChannel() {
	fmt.Println("UseChannel start!")

	ch := make(chan Result)
	go func(ch chan Result) {
		fmt.Println("goroutine!")
		ch <- Result{Message: "success"}
	}(ch)

	result := <-ch
	fmt.Println(result.Message)
}

func UseWaitGroup() {
	fmt.Println("UseWaitGroup start!")
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func UseReceiver() <-chan string {
	fmt.Println("UseReceiver start!")
	var wg sync.WaitGroup
	receiver := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			wg.Add(1)
			go func(i int) {
				msg := fmt.Sprintf("%d done", i)
				receiver <- msg
				wg.Done()
			}(i)
		}
		wg.Wait()
		close(receiver)
	}()
	return receiver
}

func caller(f func(msg string)) {
	f("test")
}

func main() {

	caller(func(msg string) {
		fmt.Println(msg)
	})

	UseChannel()
	UseWaitGroup()

	receiver := UseReceiver()
	for {
		receive, ok := <-receiver
		if !ok {
			fmt.Println("closed")
			return
		}
		fmt.Println(receive)
	}
}

// http://jxck.hatenablog.com/entry/20130414/1365960707
