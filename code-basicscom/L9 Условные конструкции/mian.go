//Реализуйте функцию Greetings(name string) string, которая вернет строку-приветствие.
//Например, при передаче имени Иван, возвращается Привет, Иван!. Учтите, что имя может быть написано в разном регистре и содержать пробелы.

package main

import (
	"fmt"
	"strings"
)

func main() {
	var msg string
	fmt.Scanln(&msg)
	fmt.Println(reetings(msg))
}

func reetings(name string) string {
	name = strings.Trim(name, " ")
	name = strings.ToLower(name)
	name = strings.Title(name)

	return fmt.Sprintf("Привет, %s!", name)
}
