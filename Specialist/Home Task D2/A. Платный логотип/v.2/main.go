package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var samplText string
	var result int // (1)
	fmt.Scan(&samplText)
	result = utf8.RuneCountInString(samplText) * 23
	if result > 100 {
		fmt.Printf("%d р. %d коп.\n", int(result/100), result%100)
	} else {
		fmt.Printf("%d коп.\n", result)
	}
}
