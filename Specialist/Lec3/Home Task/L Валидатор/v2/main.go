package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		login   string
		argMail string
	)
	fmt.Print("Ваш логин: ")
	fmt.Scan(&login)

	lookFor := "@"
	lookFor2 := "."
	containLogin := strings.Contains(login, lookFor)

	if containLogin == true || len(login) <= 9 {
		panic("Некорректный логин")
	}

	fmt.Print("Ваша почта: ")
	fmt.Scan(&argMail)

	containMail := strings.Contains(argMail, lookFor)
	containMail2 := strings.Contains(argMail, lookFor2)

	if containMail == false || containMail2 == false {
		panic("Некорректная почта")
	} else {
		fmt.Print("ОК")
	}
}
