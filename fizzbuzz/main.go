package main

import (
	"fizzbuzz/fizzbuzz"
	"fmt"
)

func main() {
	for i := 1; i < 100; i++ {
		fmt.Println(fizzbuzz.Answer(i))
	}
}
