package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// Channelをcloseしても受信は可能
	for elem := range queue {
		fmt.Println(elem)
	}
}
