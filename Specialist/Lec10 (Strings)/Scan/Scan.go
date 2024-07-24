package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var name string
	input := bufio.NewScanner(os.Stdin) // буферезированный интерфейс ввода вывода стдин
	if input.Scan() {                   // Команда захвата потока ввода и сохранения в буфере (Захват идет до \n)
		name = input.Text() //Команда возвращения  элементов, помещенных в буфер (отдаст стринг)
	} else {
		fmt.Println("Input Error")
	}
	fmt.Println(name)

	fmt.Println("For loop started:")
	for {
		if input.Scan() {
			result := input.Text()
			if result == "" {
				break
			}
			fmt.Println("Input string is:", result)
		}
	}

	// преобразование строкового литерала к чему нибудь числовому
	numStr := "10"
	numInt, _ := strconv.Atoi(numStr)
	fmt.Printf("%d is %T\n", numInt, numInt)

	numInt64, _ := strconv.ParseInt(numStr, 10, 32)
	fmt.Printf("%d is %T\n", numInt64, int32(numInt64))

	numFloat32, _ := strconv.ParseFloat(numStr, 64) // но это 64 разрятное число будет гарантированно переводится 32
	fmt.Printf("%.2f is %T\n", numFloat32, float32(numFloat32))

}
