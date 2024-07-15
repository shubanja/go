package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите строку: ")
	line, _ := reader.ReadString('\n')

	fmt.Printf("Введена строка:%s, type %T", line, line)
}
