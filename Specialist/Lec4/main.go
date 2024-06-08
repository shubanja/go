package main

import "fmt"

func main() {
	//Boolean
	var firsBoolean bool
	fmt.Println(firsBoolean)
	//Boolean operand
	aBoolean, bBoolean := true, true
	fmt.Println("AND:", aBoolean && bBoolean) // логическое умнажение  И - &&
	fmt.Println("OR:", aBoolean || bBoolean)  // логическое И
	fmt.Println("NOT:", !aBoolean)            // ! Отрицание

	//Numerics. Integres
	//int8, int16, int32, int64
	//uint8, uint16, uint32, uint64
	var a int = 32
	b := 92
	fmt.Println("Value of a:", a, "Value of b:", b, "Sum of a+b:", a+b)
}
