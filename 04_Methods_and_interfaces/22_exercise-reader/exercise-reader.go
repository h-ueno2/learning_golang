package main

import "fmt"

type MyReader struct{}

func (m MyReader) Read(rb []byte) (n int, e error) {
	for n, e = 0, nil; n < len(rb); n++ {
		rb[n] = 'A'
	}
	return
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func main() {
	m := MyReader{}
	b := make([]byte, 1024, 2048)
	m.Read(b)
	fmt.Println(fmt.Sprintf("%s", b))
}
