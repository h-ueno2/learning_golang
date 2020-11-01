package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	// 引数として渡ってきた関数を実行しただけだけど、
	// 変数のように使えるよという例
	return fn(3, 4)
}

func main() {
	// 変数に関数を代入できるよ
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))

	// 引数となる関数の引数が合っていれば、
	// なんでも渡せる
	fmt.Println(compute(math.Pow))
}
