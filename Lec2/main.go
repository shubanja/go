package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("A", "Привет! Я новый студент!")
	fmt.Println("B", "Это моя вторая программа! Я рад, что учусь здесь!!")
	// fmt.Print("first")
	// fmt.Print("Third")
	// fmt.Printf("\n Hello, my name is %s \n My age is %d \n Hi", "Bob", 42)
	//форматированный вывод
	////////////////
	//Декларирование переменной
	var age int
	fmt.Println("My age is:", age)
	age = 32
	fmt.Println("Age after assignment:", age)

	// Декларирование и инициализация пользовательским значением
	var height int = 183
	fmt.Println("My height is:", height)

	// в чем полустрогость
	var uid = 123123
	fmt.Println("My passport:", uid)
	fmt.Printf("%T", uid)

	///
	var firstVar, secondVar int = 20, 30
	fmt.Printf("firstVar:%d, secondVar:%d\n", firstVar, secondVar)
	////
	var (
		personName string = "bob"
		personAge  int    = 19
		personId   int
	)
	fmt.Printf("personName: %s\nAge: %d\nUID: %d\n", personName, personAge, personId)
	///немного странного
	var a, b = 30, "vova"
	fmt.Println(a, b)
	///немного хорошего. повторное декларирование переменой
	///var a = 200

	//коротаая декларация переменной
	count := 10
	fmt.Println(count)
	count = count + 2
	fmt.Println(count)
	//////
	//множественное присваивание через короткую декларацию :=
	aArg, bArg := 1, 5
	fmt.Println(aArg, bArg)
	aArg, bArg = 5, 5
	fmt.Println(aArg, bArg)
	//aArg, bArg := 1, 5 не даст уже переменная заделарирована
	//fmt.Println(aArg, bArg)
	//Исключения из этого правила
	bArg, cArg := 300, 200
	fmt.Println(aArg, bArg, cArg)
	//Пример
	width, length := 20.2, 30.3
	fmt.Printf("Min dimensional of rectangle is: %.0f\n", math.Min(width, length))
	///задание С
	word1, word2, word3, word4 := "имя", "твое", "мне", "знакомо"
	fmt.Println(word4, word3, word2, word1)
	fmt.Println(word3, word1, word4, word2)
	fmt.Println(word2, word4, word3, word1)
	///

}
