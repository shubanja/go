// Напишите функцию IntToString, которая преобразует и возвращает входящее число в строку
package main

import (
	"fmt"
	"strconv"
)

func main() {
	y := IntToString(5)
	fmt.Printf(" type: %T \n value: %s", y, y) // %Т тип
}

func IntToString(msg int) string {
	return strconv.Itoa(msg)
}
