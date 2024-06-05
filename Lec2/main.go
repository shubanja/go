package main

import "fmt"

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
}
