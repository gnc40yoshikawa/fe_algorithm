package main

import "fmt"

func main() {
	Text := [...]string{"A", "A", "C", "B", "B", "A", "A", "C", "A", "B", "A", "B", "A", "B"}
	Pat := [...]string{"A", "C", "A", "B", "A", "B"}
	fmt.Println(Text)
	fmt.Println(Pat)
}
