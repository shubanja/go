//На вход программе подается строка.
//Если из первого и последнего символов этой строки можно собрать слово "да" ("Да", "дА", "ДА", "да") - то программа выводит "СОГЛАСЕН" и "НЕ СОГЛАСЕН" в противном случае.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	mySlice := []rune(strings.TrimSpace(line))
	lenSlice := len(mySlice) - 1

	fmt.Println("Len:", len(mySlice))
	fmt.Printf(" first:%s  last:%s\n ", string(mySlice[0]), string(mySlice[lenSlice]))
	fmt.Print(approval(mySlice))

}

func approval(a []rune) string {
	argSlice := []rune{'Д', 'д', 'А', 'а'}
	lenSlice := len(a) - 1
	for _, run := range argSlice {
		if run == a[0] {
			for _, run := range argSlice {
				if run == a[lenSlice] {
					return "СОГЛАСЕН"
				}
			}

		}
	}
	return "НЕ СОГЛАСЕН"
}
