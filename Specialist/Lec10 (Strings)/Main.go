//Scanf считать строку "31 of month"

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	name := "Hello World / Привет мир!"
	fmt.Println(name)
	fmt.Printf("%x \n", name)
	//Строка это набор байтовый слайс со своими особенностями при обращении
	//к нижележащему массиву
	word := "Hello World / Привет мир!"
	fmt.Printf("String %s\n", word)
	//Какие значения байт сейчас находятся в слайсе word?
	fmt.Printf("Bytes: ")
	for i := 0; i < len(word); i++ {
		fmt.Printf("%x ", word[i]) // %x - формат представления 16ти ричного байта
		//Bytes: 48 65 6c 6c 6f 20 57 6f 72 6c 64 20 2f 20 d0 9f d1 80 d0 b8 d0 b2 d0 b5 d1 82 20 d0 bc d0 b8 d1 80 21
	}
	fmt.Println()
	// Каким образом получить доступ к отдельно стоящим сиволам?
	fmt.Printf("Characters: ")
	for i := 0; i < len(word); i++ {
		fmt.Printf("%c ", word[i]) // %c - формат представления символа
		//Characters: S a m p l   W o r l d
	}
	fmt.Println()
	//2. строки в го харнятся как наборы ЮТФ-8 симмволов.Каждый символ может занимать  больше чем 1 байт
	//Руна (Rune) - это стандартный встроенный тип в Го (alias над int32), Позваляет хранить
	//Единый неделимый  UTF символ ВНЕ зависимости от того сколько байт он занимает
	fmt.Printf("Runes: ")
	runSlice := []rune(word) // Преобразование слайса байт к слайсу рун
	for i := 0; i < len(runSlice); i++ {
		fmt.Printf("%c ", runSlice[i])
	}
	fmt.Println()
	//4. Итерирование по строке с использованием рун
	for id, runeVal := range word { // id это индекс  байта с которого начинается руна reuneVal,
		fmt.Printf("%c | %b | %x starts at position %d\n", runeVal, runeVal, runeVal, id) // %b - binary %x - HEX
	}
	//5. Создание строки из слайса байтов
	myByteSlice := []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x57, 0x6f, 0x72, 0x6c, 0x64} // Исходное HEX
	myStr := string(myByteSlice)
	fmt.Println(myStr)

	myDecimalByteSlice := []byte{100, 200, 101, 102, 202} // nin
	myDec := string(myDecimalByteSlice)
	fmt.Println(myDec)

	//6.Создание строки из слайса рун
	//Руны как hex
	runneHexSlice := []rune{0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x57, 0x6f, 0x72, 0x6c, 0x64}
	myStrFromRune := string(runneHexSlice)
	fmt.Println("From Runes(hex):", myStrFromRune)
	//7. Руны как литералы
	runeLiteralSlice := []rune{'f', 'o', 'l', 'o'} // '' - такм образом обозначают руны
	myStrFromLitral := string(runeLiteralSlice)
	fmt.Printf("From Runes(literals): %s type:%T\n", myStrFromLitral, myStrFromLitral)
	//7.Длина емкость строки
	fmt.Println("Length of Вася:", len("Вася"), "bytes")
	//Длина RuneCounter - количество  !рун!
	fmt.Println("Length of Вася:", utf8.RuneCountInString("Вася"), "runes")
	// Попробовал перебором постучать так посчитать
	myString := "Вася"
	lenString := 0
	for id, _ := range myString {
		if id >= 0 {
			lenString = lenString + 1
		}
	}
	fmt.Printf("Количество символов:%d \n", lenString)
	// Синтатический сахар - вычесление емкости строки бесмысленно, т.к. строка базовый тип
	fmt.Println(cap([]rune("Васе")))

	//8.Сравнение строк
	word1, word2 := "Вася", "Петя"
	if word1 == word2 {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not Equal")
	}
	fmt.Println()
	//9.Конкатинация
	word3 := word1 + word2
	fmt.Println(word3)

	//10.Строитель строк
	firstName := "Alex"
	secondName := "Johnson"
	fullName := fmt.Sprintf("%s ##### %s", firstName, secondName)
	fmt.Println(fullName)
	//11. Строки не изменяемые
	//fullName[0] = "Q"

	//12. А слайсы изменяемые
	fullNameSlice := []rune(firstName)
	fullNameSlice[0] = 'F'
	firstName = string(fullNameSlice)
	fmt.Println(firstName)
	//13. Сравнение Runes
	if 'Q' == 'Q' {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not Equal")
	}
	fmt.Println()
	//14. Где живут полезные методы со строками??
	// import "strings"

}
