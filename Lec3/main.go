package main

import "fmt"

func main() {
	var (
		age  int
		name string
	)
	fmt.Printf("Ваш возраст?:")
	fmt.Scan(&age)
	fmt.Print("Ваше имя?:")
	fmt.Scan(&name)

	fmt.Printf("my name is: %s\nMy age is: %d\n", name, age)
	fmt.Printf("name type: %T\nage type: %T\n", name, age)
}
