package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")

	channel()

	channelBuffuring()

}

func channel() {
	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}

func channelBuffuring() {
	messages := make(chan string, 2) // バッファが2
	messages <- "bufferd"
	messages <- "channel"
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
