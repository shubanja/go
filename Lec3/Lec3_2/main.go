package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		age  int
		name string
	)
	fmt.Printf("Ваш возраст?\n")
	fmt.Scan(&age)
	fmt.Print("Ваше имя?\n")
	fmt.Scan(&name)

	fmt.Printf("my name is: %s\nMy age is: %d\n", name, age)
	fmt.Printf("name type: %T\nage type: %T\n", name, age)
	//Для ручного потока ввода Fsacn
	fmt.Print("Еще раз сколько лет?\n")
	fmt.Fscan(os.Stdin, &age)
	fmt.Println("New age:", age)
}
