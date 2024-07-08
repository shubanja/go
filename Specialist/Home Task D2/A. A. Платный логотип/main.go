// "Бу'сайз" - дизайнерское агенство, основным полем деятельности которого является разработка логотипов с текстовой начинкой. У данного агенства вполне интересный прайс-лист  - за каждый символ (не важно какой) , агенство берет плату в размере 23 коп.

// Вам поручено написать программу, которая выполняет вычислительную операцию подсчета стоимости логотипа в зависимости от количества символов в строке клиента.

// Напишите программу, которая считывает строку и выводит ее предполагаемую цену в следующем формате:

// Y коп. - если цена не дотягивает до рубля.
// X р. Y коп. - если цена превышает 1 рубль.

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	var praicRUB int = 23
	var praicCOP int = 23
	var samplText string
	var resultCop int
	var result int
	//var totalPrice string
	fmt.Scan(&samplText)
	result = utf8.RuneCountInString(samplText)
	var resultRub float64 = float64(result) * praic

	if resultCop > 100 {
		fmt.Printf("%d р. %d коп.", int64(resultRub), int64(resultCop)-int64(resultRub))
	} else {
		fmt.Printf("цена %d копеек", resultRub)
	}
}
