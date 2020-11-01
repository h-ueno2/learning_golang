package main

import "fmt"

func main() {
  sum := 1
  // 他の言語でいうWhile
  for sum < 1000 {
    sum += sum
  }
  fmt.Println(sum)
}
