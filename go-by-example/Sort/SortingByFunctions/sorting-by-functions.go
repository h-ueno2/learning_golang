package main

import (
	"fmt"
	"sort"
)

// 独自のソート順を使いたい場合はそのための型を定義する
type byLength []string

// sort.Interfaceの Len, Lessm Swapの3メソッドを実装する
// このサンプルでは[]stringだけど、独自のtypeでも同じように出来るのだろう。
func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
