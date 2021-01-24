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

	synchronization()

	channelDirections()
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

func synchronization() {
	done := make(chan bool, 1)
	go worker(done)

	<-done // これがないとdoneが返ってくるまで待たないで終了してしまう。
}

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

// channelは送信専用、受信専用を明記することができる
// 受信専用channel
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// 送信専用channelを使う
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func channelDirections() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
