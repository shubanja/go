// Пора начинать проникаться искусством математики. Для этого вам предложено решить простейшую задачу - определить количество корней указанного уравнения.

// Напишите программу, которая принимает на вход три вещественных числа

//  и отвечает на вопрос : сколько корней имеет соответствующее этим числам уравнение вида

// если уравнение имеет 2 вещественных корня - вывести "два корня"
// если уравнение имеет 1 корень (совпадающие  корни тоже считаются за один) - вывести "один корень"
// если уравнение не имеет вещественных корней или не имеет решений - "корней нет"
// Гарантируется, что как минимум один из коэффициентов не равен нулю.

package main

import (
	"fmt"
)

func main() {
	var aArg, bArg, cArg, dArg int
	fmt.Scan(&aArg, &bArg, &cArg)
	dArg = bArg*bArg - 4*aArg*cArg
	fmt.Println(dArg)

	if aArg != 0 {
		if dArg > 0 {
			fmt.Println("Два коня")
		} else if dArg == 0 {
			fmt.Println("Один корень")
		} else {
			fmt.Println("Ноль корней")
		}
	} else if aArg == 0 || cArg == 0 {
		fmt.Println("простое уравнение")
	} else if aArg == 0 && bArg == 0 {
		fmt.Println("Два коэфицента нули? ну нет корней для тебя ноль")
	}
}