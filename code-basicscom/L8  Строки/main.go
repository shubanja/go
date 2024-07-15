//Реализуйте функцию IsValid(id int, text string) bool, которая проверяет,
//что переданный идентификатор id является положительным числом и текст text не пустой.

package main

import (
	"fmt"
)

func main() {
	id := 1
	text := "Hi"
	fmt.Println(IsValid(id, text))
}
func IsValid(id int, text string) bool {
	return text != "" && id > 0
}
