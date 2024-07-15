package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите строку: ")
	line, _ := reader.ReadString(' ')

	fmt.Println("Введена строка:", line)
}
