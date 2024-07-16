//Реализуйте функцию Greetings(name string) string, которая вернет строку-приветствие....
//Например, при передаче имени Иван, возвращается Привет, Иван!. Учтите, что имя может быть написано в разном регистре и содержать пробелы.

package main

import (
	"fmt"
)

func main() {
	var a string
	var b string
	fmt.Print("Enter domain:")
	fmt.Scanln(&a)
	fmt.Print("Enter local:")
	fmt.Scanln(&b)
	fmt.Print(DomainForLocale(a, b))
}

func DomainForLocale(domain, locale string) string {
	if locale == "" {
		locale = "en"
	}
	return fmt.Sprintf("%s.%s", locale, domain)

}
