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
	var stop float32 = 000

	fmt.Print("Find average arithmetic value\n" + "Enter 000 - to calculate\n")

	for {
		fmt.Println("Enter number:")
		_, _ = fmt.Scanln(&input)
		if input == stop {
			break
		}
		fmt.Print("The data is correct?\n(Y - yes / N - No):")
		_, _ = fmt.Scanln(&inpCheck)
		switch inpCheck {
		case "y", "Y":
			sum = sum + input
			vol = vol + 1
		case "n", "N":
		default:
			//Отработает  если выше не нашел в условиях
			fmt.Println("unknown, enter number again ")
		}
	}
	//}
	fmt.Printf("Average mean: %.2f\n", sum/float32(vol))
	fmt.Printf("SUM %.2f\nVOL %d", sum, vol)

}
