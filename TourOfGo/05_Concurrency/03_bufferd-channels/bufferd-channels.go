package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	// ch <- 4 // buffer上限を超えるとfatal errorが発生する
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
