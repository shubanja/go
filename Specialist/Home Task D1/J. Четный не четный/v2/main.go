package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Print("Введите целое число: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		panic("Проверьте ввод данных. Это не число!")
	}
	if number%2 == 0 {
		fmt.Printf("Число [%d] - чётное", number)
	} else {
		fmt.Printf("Число [%d] - нечётное", number)
	}
}
