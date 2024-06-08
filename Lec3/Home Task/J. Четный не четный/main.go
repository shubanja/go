package main

import "fmt"

func main() {

	var (
		number int
	)

	fmt.Println("Програма для проверки четности или нечетности.\n Введите число:5")
	fmt.Scan(&number)

	if number%2 == 0 {
		fmt.Printf("Это число %d четное.\n", number)
	} else {
		fmt.Printf("Это число %d нечетное.\n", number)
	}

}
