//Scanf считать строку "31 of month"

package main

import "fmt"

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
		//Bytes: 53 61 6d 70 6c 20 57 6f 72 6c 64
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
	myByteSlice := []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x57, 0x6f, 0x72, 0x6c, 0x64} // HEX
	myStr := string(myByteSlice)
	fmt.Print(myStr)
}
