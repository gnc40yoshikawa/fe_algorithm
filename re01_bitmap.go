package main

import (
	"fmt"
	"strings"
)

func main() {
	//対象文字列
	var Text = []string{"A", "A", "C", "B", "B", "A", "A", "C", "A", "B", "A", "B", "A", "B"}
	//検索文字列
	var Pat = []string{"A", "C", "A", "B", "A", "B"}
	//ビットマスク
	var Mask [26]int
	fmt.Println(Text)
	fmt.Println(Pat)

	GenerateBitMask(Pat, Mask)
}

func GenerateBitMask(Pat []string, Mask [26]int) {
	//検索文字列の要素数を取得
	var PatLen int = len(Pat)
	//ビットマスクを初期化
	for i := 0; i < 26; i++ {
		Mask[i] = 0b0
	}
	fmt.Printf("%016b\n", Mask)

	for i := 0; i < PatLen; i++ {
		//何番目のアルファベットに位置するか
		fmt.Println(Index(Pat[i]))
	}
}

func Index(s string) any {
	var obj string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return strings.Index(obj, s)
}
