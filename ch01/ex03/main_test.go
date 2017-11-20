package main

import (
	"fmt"
	"strings"
	"testing"
)

var args = []string{"run", "arg0", "arg1"}

//TODO:GOのテスト命名規則にそってないが例題クリアを優先する・・・ https://qiita.com/Jxck_/items/8717a5982547cfa54ebc
func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println(strings.Join(args, " "))
	}
}

func BenchmarkLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s, sep := "", ""
		for _, arg := range args[1:] {
			s += sep + arg
			sep = " "
		}
		fmt.Println(s)
	}
}
