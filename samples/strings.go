package main

import (
	"fmt"
	"strings"
)

func main() {

	// 文字列長
	fmt.Println(len("abcdefg"))                  // 7
	fmt.Println(len("あいうえお"))                    // 15
	fmt.Println(len(strings.Split("あいうえお", ""))) // 5

	// 文字列を含むかどうかの判定
	fmt.Println(strings.Contains("hello, world", "hello")) // true

	// 文字列のカウント
	fmt.Println(strings.Count("hello, world", "lo")) // 1

	// 文字列の位置を取得
	fmt.Println(strings.Index("hello, world", "ld")) // 10

	// 文字列の連結
	var ary = []string{"hello", "go"}
	fmt.Println(strings.Join(ary, ", ")) // hello,go

	// 大文字変換
	fmt.Println(strings.ToUpper("hello, world")) //   HELLO, WORLD

	// 小文字変換
	fmt.Println(strings.ToLower("GO")) // go

	// 単語の頭の文字だけをtitle caseにする
	fmt.Println(strings.Title("effective go programming")) // Effective Go Programming

	// トリムは文字列を指定する
	fmt.Println(strings.Trim("***hello***", "*")) // hello

	// split
	fmt.Println(strings.Split("a,b,c,d,e,f,g", ",")) // [a b c d e f g]

	//文字列の繰り返し
	fmt.Println(strings.Repeat("オラ", 10)) // オラオラオラオラオラオラオラオラオラオラ

	// 置換
	r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
	fmt.Println(r.Replace("<script>alert('hello');</script>")) // &lt;script&gt;alert('hello');&lt;/script&gt;

}
