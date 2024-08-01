//Профессиональная игра в слова
//Иван и Петр по очереди вводят слова в программу. Начиная со второго введенного слова программа анализирует , совпадает ли первая буква нового слова с последней буквой предыдущего слова.
//Если да, то программа работает дальше (считывает еще слово). Если нет - выводит последнее введенное слово и завершает работу.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	firstSlice := make([]rune, 0)
	secondSlice := make([]rune, 0)
	//firstSlice = append(firstSlice, secondSlice)

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	firstSlice = []rune(strings.TrimSpace(line))
	fmt.Println(lenRuneElement(firstSlice))

	for {
		line, _ := reader.ReadString('\n')
		secondSlice = []rune(strings.TrimSpace(line))
		fmt.Println(firstSlice[lenRuneElement(secondSlice)])
		if secondSlice[0] == firstSlice[lenRuneElement(secondSlice)] {
			firstSlice = append(firstSlice, secondSlice...)
			fmt.Println(string(firstSlice[:]))
		}

	}

	fmt.Println(firstSlice, secondSlice)
}

func lenRuneElement(s []rune) int {
	return len(s) - 1
}
