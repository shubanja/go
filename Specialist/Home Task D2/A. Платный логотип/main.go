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

	var praic float64 = 0.23
	var samplText string
	var result int
	fmt.Scan(&samplText)
	result = utf8.RuneCountInString(samplText)
	var resultRub float64 = float64(result) * praic
	var RUB int = int(resultRub)
	if resultRub > 1.00 {
		fmt.Printf("%d р. %d коп.", int64(RUB), int((resultRub-float64(RUB))*100))
	} else {
		fmt.Printf("%d коп.", int(resultRub*100))
	}
}
