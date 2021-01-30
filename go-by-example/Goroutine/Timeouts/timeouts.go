package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	go dummy(c1)

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		// 1秒経ったらこっち
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go dummy(c2)

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}

func dummy(msg chan<- string) {
	time.Sleep(2 * time.Second)
	msg <- "result"
}
