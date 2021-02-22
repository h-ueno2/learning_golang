package fizzbuzz

import "fmt"

const (
	fizzbuzz string = "FizzBuzz"
	fizz     string = "Fizz"
	buzz     string = "Buzz"
)

func Answer(i int) string {
	if i%15 == 0 {
		return fizzbuzz
	} else if i%3 == 0 {
		return fizz
	} else if i%5 == 0 {
		return buzz
	}
	return fmt.Sprint(i)
}
