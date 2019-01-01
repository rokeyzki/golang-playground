package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	for i := 0; i < 5000; i++ {
		go printHelloWorld(i, ch)
	}

	// time.Sleep(time.Millisecond)
	for {
		msg := <-ch
		fmt.Println(msg)
	}
}

func printHelloWorld(i int, ch chan string) {
	// fmt.Println("Hello Conc from goroutine $d!", i)
	for {
		ch <- fmt.Sprintf("Hello Conc from " + "goroutine %d!\n", i)
	}
}
