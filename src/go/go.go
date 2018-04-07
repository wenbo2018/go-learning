package main

import (
	"fmt"
	"time"
)

func hello() {
	for i := 1; i < 1000; i++ {
		fmt.Print(i)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Hello world goroutine")
}
func main() {
	go hello()
	for i := 1000; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
