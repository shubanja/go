// Пользователь вводит числа числа одно за другим (каждое с новой строки, все числа ∈ N) до тех пор, пока не будет введен нуль. Ваша программа должна выводить вводимые числа до тех пор, пока не будет введен нуль (Сам нуль выводить не требуется). Никаких арифметических операций над числами производить не требуется.

package main

import (
	"fmt"
)

func main() {

	var x int
outer:
	for {
		fmt.Scan(&x)
		if x == 0 {
			break outer
		}
		fmt.Println(x)
	}
}
