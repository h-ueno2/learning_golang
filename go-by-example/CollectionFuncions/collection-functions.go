package main

import (
	"fmt"
	"strings"
)

// GoにはJavaでいうstream()のようなコレクション制御の機能はない。
// そのため、自作する必要がある。

// Index ...
// 文字列tを始めに見つけたインデックスを返却。
// 見つからなければ-1を返す。
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// Include ...
// スライスの中に文字列tが含まれるならtrueを返却する
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// Any ...
// スライスの中に引数の関数が真になる文字列があればtrueを返却
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// All ...
// スライスの中の文字列がいずれも引数の関数を真にするならtrueを返却
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// Filter ...
// スライスの中の文字列から条件が真となる文字列だけを含むスライスを返却する
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0) //返却用のスライス
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// Map ...
// スライスの文字列に関数fを実行した結果で新しいスライスを作成して返却する
func Map(vs []string, f func(string) string) []string {
	vsf := make([]string, 0) // 返却用のスライス
	for _, v := range vs {
		vsf = append(vsf, f(v))
	}
	return vsf
}

func main() {
	var strs = []string{"peach", "apple", "pear", "plum"}
	fmt.Println(Index(strs, "pear"))
	fmt.Println(Include(strs, "grape"))
	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))
	fmt.Println(Map(strs, strings.ToUpper))

}
