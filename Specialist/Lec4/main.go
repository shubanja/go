package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func main() {
	//Boolean => Defalt false
	var firsBoolean bool
	fmt.Println(firsBoolean)
	//Boolean operands
	aBoolean, bBoolean := true, true
	fmt.Println("AND:", aBoolean && bBoolean) // логическое умнажение  И - &&
	fmt.Println("OR:", aBoolean || bBoolean)  // логическое ИЛИ
	fmt.Println("NOT:", !aBoolean)            // ! Отрицание

	//Numerics. Integres
	//int8 (8bit), int16, int32, int64 (8byte)
	//uint8, uint16, uint32, uint64
	var a int = 32
	b := 92
	fmt.Println("Value of a:", a, "Value of b:", b, "Sum of a+b:", a+b)
	//Вывод типа через %T
	fmt.Printf("Type is %T\n", a)
	//Узнает сколько байт занимает переменая типа int
	fmt.Printf("type %T size of %d bytes\n", a, unsafe.Sizeof(a))
	//Эксперемент 1 при кородком объявлеии плавтформа зависимы
	fmt.Printf("type %T size of %d bytes\n", a, unsafe.Sizeof(b))

	//Эксперемент 2
	var first32 int32 = 12
	var second64 int64 = 13
	fmt.Println(int64(first32) + second64) // явно преведение типов. возможно колизии

	//Эксперемент 3. Если проводятся арифмитичные операции нужно явные приведения ипов
	// над int u intX то обязательно нужно использовать механизм приведения тк: int != int64
	var third64 int64 = 2342334
	var forthInt int = 12312312
	//fmt.Println(third64 + forthInt) надо явно
	fmt.Println(third64 + int64(forthInt))
	/// + - * / %

	//uint только положительные чила
	//Numerics. Float
	//Float32, float64
	floatFirst, floatSecond := 5.67, 12.54
	fmt.Printf("type of a %T and type of %T\n", floatFirst, floatSecond) // по умлочанию float64
	sum := floatFirst + floatSecond
	sub := floatFirst - floatSecond
	fmt.Println("Sum:", sum, "Sub:", sub)
	fmt.Printf("Sum:%.3f, Sub:%.3f\n", sum, sub)

	//Numerics. Complex
	c1 := complex(5, 7)
	c2 := 12 + 32i
	fmt.Println(c1 + c2)
	fmt.Println(c1 * c2)

	//Strings это набор БАЙТ
	Name := "Федя"
	LastName := "Pukin"
	concat := Name + " " + LastName
	fmt.Println("Full name:", concat)
	fmt.Println("Length of string:", Name, len(Name)) // len - Возращает кол элементов наборе
	fmt.Println("Amount of cahrs:", Name, utf8.RuneCountInString(Name))
	//Rune - руна это один ЮТФ символ

	//Поиск в строке конкотинация
	totalString, subString := "ABCDEF", "CD" //риеестр имеет значение
	fmt.Println(strings.Contains(totalString, subString))
	//rune -> alias int32
	var sampleRune rune
	var anotherRune rune = 'Q' // для инициализации руны в симвл значении использ ' '
	var thirdRune rune = 2345
	fmt.Printf("Run as char %c\n", sampleRune)
	fmt.Printf("Run as char %c\n", anotherRune)
	fmt.Printf("Run as char %c\n", thirdRune)
	// Сравнение стринг
	fmt.Println(strings.Compare("abcd", "bc")) // -1 не очень понял\

	var aArg byte // alias uint8
	fmt.Println("Byte:", aArg)

}
