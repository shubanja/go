//Есть такая вещь, как параметрическая функция в го. Так вот задача проста.
//Реализовать функцию которая считает среднее арифметическое переданных в нее значений. Входных значений может быть от 2 до n

package main

import (
	"fmt"
)

func main() {
	var vol int
	var input float32
	var sum float32
	var inpCheck string
	fmt.Print("Find average arithmetic value\n" + "To complete enter => exit\n")
	for {
		fmt.Println("Enter number:")
		_, _ = fmt.Scanln(&input)
		fmt.Print("Данные верны?(Y/N):")
		_, _ = fmt.Scanln(&inpCheck)
		switch inpCheck {
		case "y":
			fmt.Println("Yes")
		case "n":
			fmt.Println("No")
			input = 0
		default:
			//Отработает  если выше не нашел в условиях
			fmt.Println("unknown")
		}
		sum = sum + input
		vol = vol + 1
	}
	fmt.Printf("Average mean: %.2f\n", sum/float32(vol))
	fmt.Printf("SUM %.2f\nVOL %d", sum, vol)

}
