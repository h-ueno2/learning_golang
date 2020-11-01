package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader *rot13Reader) Read(rb []byte) (n int, err error) {
	n, err = reader.r.Read(rb)
	if err == nil {
		for i, v := range rb {
			switch {
			case v >= 'A' && v <= 'Z':
				rb[i] = (v-'A'+13)%26 + 'A'
			case v >= 'a' && v <= 'z':
				rb[i] = (v-'a'+13)%26 + 'a'
			}
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
