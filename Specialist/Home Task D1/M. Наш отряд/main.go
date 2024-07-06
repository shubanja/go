package main

import (
	"fmt"
	"strings"
)

func main() {
	var aArg, bArg, cArg string
	var chekSum int

	fmt.Scan(&aArg)

	if strings.Compare(aArg, "раз") == 0 {
		chekSum = chekSum + 1
	} else if strings.Compare(aArg, "один") == 0 {
		chekSum = chekSum + 1
	} else if strings.Compare(aArg, "1") == 0 {
		chekSum = chekSum + 1
	}

	fmt.Scan(&bArg)

	if strings.Compare(bArg, "два") == 0 {
		chekSum = chekSum + 1
	} else if strings.Compare(bArg, "2") == 0 {
		chekSum = chekSum + 1
	}

	fmt.Scan(&cArg)

	if strings.Compare(cArg, "три") == 0 {
		chekSum = chekSum + 1
	} else if strings.Compare(cArg, "3") == 0 {
		chekSum = chekSum + 1
	}

	if chekSum == 3 {
		fmt.Println("ОК")
		return
	}
	fmt.Println("НЕ ПРАВИЛЬНО")
}

/* Напишите программу, которая проверяет правильность расчета на "раз два три".

Пользователь вводит последовательно 3 строки.

если эти строки "раз", "два", "три" - вывести "ОК" (кириллица)
если вместо строки "раз" введена "один" - вывести "ОК" (кириллица)
если вместо всех слов введены соответствующие числа "1", "2", "3" - вывести "ОК" (кириллица)
"НЕ ПРАВИЛЬНО" - во всех остальных случаях */
