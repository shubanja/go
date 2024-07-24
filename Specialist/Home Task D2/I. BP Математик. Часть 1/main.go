//Уровень - математик с большой буквы
//Всегда интересно , является ли число простым. Простые числа - натуральные числа, которые имеют только 2 делителя : 1 и само это число (т.к. они обязаны отличаться - 1 это не простое число). Ваша задача вывести все делители числа N. В случае, если число N - простое, нужно дополнительно сообщить об этом.
//
//Формат ввода
//Одно чсило N - натуральное число (< 1000000).
//
//Формат вывода
//Все делители числа N в порядке возрастания. Если число N - простое, то еще вывести слово ACHTUNG

package main

import (
	"fmt"
)

func main() {
	var value int // checked number
	var i int
	var naturalNUB int

	fmt.Println("Enter NUMber")
	_, err := fmt.Scan(&value) // enter number
	if err != nil {
		return
	}

	for i = 1; i <= value; i++ {
		if value%i == 0 {
			fmt.Printf(" %d", i)
			naturalNUB = naturalNUB + 1

		}

	}
	fmt.Print("\n")
	if naturalNUB == 2 {
		fmt.Println("ACHTUNG")
	}
}
