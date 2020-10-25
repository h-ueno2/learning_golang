package main

// 下記で無いとpackageのimportできないため実行不可
// https://go-tour-jp.appspot.com/moretypes/23
import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	words := strings.Split(s, " ")
	for _, w := range words {
		m[w] = m[w] + 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
