package main

import (
	"fmt"
	"strings"
)

func main() {
	// for init; condition; post {
	// init - блок инициализации переменных цикла
	// condition - условие (если верно - то тело цикла выполняется, если нет - то цикл завершается)
	// post - изменение переменной цикла (инкрементарное действие, декрементарное действие)
	// }

	for i := 0; i <= 6; i++ {
		fmt.Printf("Current value: %d\n", i)
	}
	//Важный момент : в качестве init может быть использовано выражение присваивания ТОЛЬКО через :=

	//break - команда, прерывающая текущее выполнение тела цикла и передающая управление инструкциям, следующим
	// за циклом

	for i := 0; i > 9; i++ {
		if i > 5 {
			break
		}
		fmt.Printf("Current value: %d\n", i)
	}
	fmt.Println("Ater for loop with BREAK")

	//continue - команда, прерывающая текущее выполнение тела цикла и передающая управления СЛЕДУЮЩЕЙ ИТЕРАЦИИ ЦИКЛА
	for i := 0; i <= 20; i++ {
		if i > 10 && i < 15 {
			continue
		}
		fmt.Printf("Current value: %d\n", i)
	}
	fmt.Println("Ater for loop with continue")

	//v.2

	for i := 0; i <= 20; i++ {
		if i < 10 || i > 15 {
			continue
		}
		fmt.Printf("Current value: %d\n", i)
	}
	fmt.Println("Ater for loop with continue")

	// Вложенный цикл и лейблы
	for i := 0; i < 10; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Print("Конец\n")

	// Оставновить внешний цикл из внутреннего можно лейблами
outer:
	for i := 0; i <= 2; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("i : %d and j : %d sum i+j=%d\n ", i, j, i+j)
			if i == j {
				break outer
			}
		}
	}

	var stopFlag bool
	var x = 0
	fmt.Println(stopFlag)
	fmt.Println(x)
	for i := 0; i <= 2; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("i : %d and j : %d sum i+j=%d\n ", i, j, i+j)
			if i == j {
				stopFlag = true // присваевает true
				x = 10
				break
			}
		}
		if stopFlag {
			break
		}
	}
	fmt.Println(stopFlag)
	fmt.Println(x)

	// Модифиаация цикла for
	//1. Классчиеский цикл while do
	var loopVar int = 0
	// while (loopVar < 10){
	// 	....
	// 	loopVar++
	// }
	for loopVar <= 10 {
		fmt.Printf("In while like loop %d\n", loopVar)
		loopVar++
	}
	//2. Классический бесконечный цикл
	//for {
	//	fmt.Println("KEK")
	//}

	var password string
outer2:
	for {
		fmt.Print("Insert password: ")
		fmt.Scan(&password)
		if strings.Contains(password, "1234") {
			fmt.Println("Weak password . Try again")
		} else {
			fmt.Println("Password Accepted")
			break outer2
		}
	}
	//цикл с моножественными переменными
	for x, y := 0, 2; x <= 10 && y <= 8; x, y = x+3, y+1 {
		fmt.Printf("%d + %d = %d\n", x, y, x+y)
	}

}
