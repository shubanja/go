package main

import (
	"fmt"
	"strings"
)

func main() {
	//Классический условный оператор
	//if condition {
	//	//body
	//}

	// если условие1 {
	// 	блок кода в случае истинности условия1
	//   } иначе если условие2 {
	// 	блок кода в случае истинности условия2
	//   } иначе {
	// 	блок кода в случае ложности всех предыдущих условий
	//   }

	var value int
	fmt.Scan(&value)

	if value%2 == 0 {
		fmt.Println("The number", value, "is even")
	} else { // Оператор с блоком else (если)
		fmt.Println("The number", value, "is odd")
	}
	// if condition1 {

	// } eLse if condition2 {

	// } else if conditionN {

	// } eLse {

	// }

	var color string
	fmt.Scan(&color)
	if strings.Compare(color, "green") == 0 {
		fmt.Println("Color is green") // блок кода в случае истинности условия
	} else if strings.Compare(color, "red") == 0 {
		fmt.Println("Color is red")
	} else {
		fmt.Println("Unknown color") // блок кода в случае ложности условия
	}

	// Инициализация пепременой в блоке условного оператора2
	if num := 10; num%2 == 0 {
		fmt.Println("EVEN")
	} else {
		fmt.Println("ODD")
	}

	/*
		var age int = 10
		if age > 7 {
			fmt.Println("Go to school")
		} else {
			fmt.Println("Another case")
		}
	*/

	//не идеоматично есть правило избегать else

	if width := 100; width > 100 {
		fmt.Println("Width > 100")
	} else {
		fmt.Println("width <= 100")
	}

	//идеоматично заканчивать return
	if height := 101; height > 100 {
		fmt.Println("height > 100")
		return
	}
	fmt.Println("heigt <= 100")

}
