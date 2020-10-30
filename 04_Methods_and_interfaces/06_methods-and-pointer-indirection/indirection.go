package main

import "fmt"

type Vertex struct {
	X, Y float64
}

// レシーバの変数を変更する場合はポインタレシーバにする必要がある
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := Vertex{3, 4}
	p.Scale(2)
	ScaleFunc(&p, 10)

	fmt.Println(v, p)
}
