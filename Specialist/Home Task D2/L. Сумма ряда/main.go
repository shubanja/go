//Напишите программу, которая считывает число n ∈ N, затем n чисел (целых), каждое с новой строки.
//
//После чего программа вычисляет и печатает знакочередующуюся сумму ряда (первое число прибавить, второе вычесть, третье прибавить, четвертое вычесть, и т.д.) .
//Например, для чисел 1,2,3,4,5  сумма будет выглядеть следующим образом:

package main

import (
	"fmt"
)

func main() {
	var aArg int
	var sumArg int
	var nArg int

outer:
	for {
		_, err := fmt.Scan(&aArg)
		if err != nil {
			break outer
		}
		nArg = nArg + 1
		if nArg%2 == 0 {
			sumArg = sumArg - aArg // чётное
			fmt.Println("сумм -:", sumArg)
		} else {
			sumArg = sumArg + aArg // не четное
			fmt.Println("сумм +:", sumArg)
		}

	}
	fmt.Println("Total num:", nArg)
	fmt.Print("Summ num:", sumArg)
}
