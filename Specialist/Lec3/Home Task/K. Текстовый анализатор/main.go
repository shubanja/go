package main

import (
	"fmt"
	"strings"
)

func main() {
	var text string
	fmt.Scan(&text)
	lookFor := "чат"
	contain := strings.Contains(text, lookFor)
	if contain == true {
		fmt.Print("БОТ")
	} else {
		fmt.Print("НЕ БОТ")
	}

}
