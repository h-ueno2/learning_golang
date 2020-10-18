package main

import "fmt"

func main() {
  i, j := 42, 2701

  p := &i
  fmt.Println(*p)
  *p = 21
  fmt.Println(i)
  fmt.Println(p) // メモリアドレスが表示される

  p = &j
  *p = *p / 37
  fmt.Println(j)
}
