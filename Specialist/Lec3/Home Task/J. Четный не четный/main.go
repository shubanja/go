package main

import "fmt"

func main() {

	var (
		aArg float64
		bArg float64
	)

	fmt.Println("Програма для складывания двух чисел и проверки четности.\n Введите первое число:")
	fmt.Scan(&aArg)
	fmt.Println("Введите второе число:")
	fmt.Scan(&bArg)

	number := aArg + bArg

	if number%2 == 0 {
		fmt.Printf("Это число %d четное.\n", number)
	} else {
		fmt.Printf("Это число %d нечетное.\n", number)
	}

}
