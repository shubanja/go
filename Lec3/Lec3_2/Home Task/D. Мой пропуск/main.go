package main

import (
	"fmt"
)

func main() {
	var (
		firstname string
		lastname  string
		age       int
	)

	fmt.Print("Ваше имя?\n")
	fmt.Scan(&firstname)
	fmt.Print("Ваша Фамилия?\n")
	fmt.Scan(&lastname)
	fmt.Print("Ваша возраст?\n")
	fmt.Scan(&age)

	fmt.Printf("Имя: %s , Фамилия: %s , Возраст: %d . Студент BPS", firstname, lastname, age)
}
