//Значение переменной можно изменять в любой момент:
//
//// двоеточие используется только при инициализации
//num := 22
//num = 33

package main

import "fmt"

func main() {
	var firstName, lastName = "Johon", "Smith"
	fmt.Println(firstName, lastName)
}
