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
	mySice := []rune{'Д', 'д', 'А', 'а'}
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	mySlice := []rune(strings.TrimSpace(line))
	x := len(mySlice) - 1
	fmt.Println(len(mySlice))
	fmt.Println(x)
	fmt.Printf("%s _ %s\n", string(mySlice[0]), string(mySlice[x]))

	/*for _, run := range mySice[0:1] {
		run ==
	}*/

	/*	if mySlice[0] == 'Д' || mySlice[0] == 'д' && mySlice[x] == 'А' || mySlice[x] == 'а' {
			fmt.Println("СОГЛАСЕН")
		} else {
			fmt.Println("НЕ СОГЛАСЕН")
		}*/

	for i, v := range mySlice {
		mySlice[i]++
		fmt.Printf("id: %d. Element:%s\n", i, string(v))
	}

}
