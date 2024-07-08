package main

import "fmt"

func main() {
	//aArg := 12345
	third := 1.0 / 3
	fmt.Println(third)            // Выводит: 0.3333333333333333
	fmt.Printf("%v\n", third)     // Выводит: 0.3333333333333333
	fmt.Printf("%f\n", third)     // Выводит: 0.333333
	fmt.Printf("%.3f\n", third)   // Выводит: 0.333
	fmt.Printf("%f\n", third)     // Выводит: 0.333333
	fmt.Printf("%4.4f\n", third)  // Выводит: 0.33
	fmt.Printf("%07.4f\n", third) // Выводит: 0.3333
	fmt.Printf("%6.2f\n", third)  // Выводит: 000.33
	fmt.Printf("Ограниченный вывод %d\n", 12345%8080)
	fmt.Printf("Ограниченный вывод %d\n", 12345%70)
	fmt.Printf("Ограниченный вывод %d\n", 12345%60)
	fmt.Printf("Ограниченный вывод %d\n", 12345%50)
	fmt.Printf("Ограниченный вывод %d\n", 12345%40)
	fmt.Printf("Ограниченный вывод %d\n", 12345%30)
	fmt.Printf("Ограниченный вывод %d\n", 12345%20)
	fmt.Printf("Ограниченный вывод %d\n", 12345%10)
	fmt.Printf("Ограниченный вывод %d\n", 12345%1)
	fmt.Printf("Ограниченный вывод %d\n", 12345%100)
	fmt.Printf("Ограниченный вывод %d\n", 12345%1000)
	fmt.Printf("Ограниченный вывод %d\n", 12345%100000)
	fmt.Printf("Ограниченный вывод %d\n", 12345%00111)
}
