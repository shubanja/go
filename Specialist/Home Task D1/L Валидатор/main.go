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
	fmt.Print("Ваша почта: ")
	fmt.Scan(&argMail)
	lookFor := "@"
	lookFor2 := "."
	contain := strings.Contains(login, lookFor)
	containMail := strings.Contains(argMail, lookFor)
	containMail2 := strings.Contains(argMail, lookFor2)

	if contain == true {
		fmt.Printf("Некорректный логин содержит [%s] ", lookFor)
	} else if len(login) <= 9 {
		fmt.Print("Некорректный логин содержит < 9 символов")
	} else if containMail != true {
		fmt.Printf("Некорректный адресс не содержит [%s]", lookFor)
	} else if containMail2 != true {
		fmt.Printf("Некорректный адресс не содержит [%s]", lookFor2)
	} else {
		fmt.Print("ОК")
	}

}
