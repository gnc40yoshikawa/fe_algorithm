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

	fmt.Println(Text)
	fmt.Println(Pat)

	BitapMatch(Text, Pat)
}

func BitapMatch(Text []string, Pat []string) int {
	var Goal, Status int
	var Mask [26]int
	var i, TextLen, PatLen int

	TextLen = len(Text)
	PatLen = GenerateBitMask(Pat, Mask)
	Status = 0
	Goal = 1 << PatLen

	for i = 0; i < TextLen; i++ {
		Status = (Status << 1) | 1
		Status = Status & Mask[Index(Text[i])]
	}
	fmt.Printf("%016b\n", Status)
	fmt.Printf("%016b\n", Goal)
	return Goal
}

func GenerateBitMask(Pat []string, Mask [26]int) int {
	//検索文字列の要素数を取得
	var PatLen = len(Pat)
	//ビットマスクを初期化
	for i := 0; i < 26; i++ {
		Mask[i] = 0
	}
	fmt.Printf("%016b\n", Mask)

	for i := 0; i < PatLen; i++ {
		Mask[Index(Pat[i])] = (1 << i) | Mask[Index(Pat[i])]
		//何番目のアルファベットに位置するか
		//fmt.Println(Index(Pat[i]))
		//fmt.Printf("%016b\n", Mask[Index(Pat[i])])
	}
	fmt.Printf("%016b\n", Mask)
	return PatLen
}

func Index(s string) int {
	var obj string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return strings.Index(obj, s)
}
