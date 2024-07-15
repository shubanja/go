package main

import (
	"fmt"
)

func main() {
	id := 1
	text := "rtrt"
	fmt.Println(IsValid(id, text))
}
func IsValid(id int, text string) bool {
	if text != "" && id > 0 {
		return true
	} else {
		return false
	}
}
