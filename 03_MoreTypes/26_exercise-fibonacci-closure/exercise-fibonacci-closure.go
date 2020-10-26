package main

import "fmt"

func fibonacci() func() int {
	a, b := 1, 0
	return func() int {
		c := b
		a, b = b, a+b
		return c
	}
}

func main() {
	f := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
