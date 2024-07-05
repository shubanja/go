package main

/* Напишите программу, которая проверяет правильность расчета на "раз два три".

Пользователь вводит последовательно 3 строки.

если эти строки "раз", "два", "три" - вывести "ОК" (кириллица)
если вместо строки "раз" введена "один" - вывести "ОК" (кириллица)
если вместо всех слов введены соответствующие числа "1", "2", "3" - вывести "ОК" (кириллица)
"НЕ ПРАВИЛЬНО" - во всех остальных случаях */

import (
	"fmt"
)

func main() {
	var (
		aArg      string
		lookForA  string = "Раз"
		lookForA2 string = "Один"
		lookForA3 int    = 1
		//bArg int
		//cArg int
	)

	fmt.Scan(&aArg)

	if aArg == lookForA || lookForA2 || lookForA3 {
		fmt.Print("ok")
	} else {
		fmt.print("NOT")
	}

	fmt.Printf("%d", aArg)
}
