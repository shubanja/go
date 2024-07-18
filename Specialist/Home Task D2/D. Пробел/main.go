package main

import (
	"fmt"
	"strings"
)

func main() {
	var text string

outer:
	for {
		fmt.Scanln(&text)
		fmt.Printf("type: %T, out:%s \n", text, text)
		if Validate(text) {
			break outer
		}
	}
}

func Validate(req string) bool {
	if req == "" || strings.Contains(req, "\n") {
		return true
	}
	return false
}
