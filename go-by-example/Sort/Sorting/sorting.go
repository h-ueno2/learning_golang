package main

import (
	"fmt"
	"sort"
)

func main() {
	// ソートは対象の型によって使用する関数が異なる
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	// ソート済みかどうかの判定もできる
	// booleanが返却される関数の名称が"is"始まりではないのはGoの文化なのかな？
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}
