package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Vertex 型のレシーバをもつ
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// レシーバの変数を変更する場合はポインタレシーバにする必要がある
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
