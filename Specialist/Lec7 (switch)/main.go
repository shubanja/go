package main

import "fmt"

func main() {
	//Что бы так не писать есть Switch!
	// price := 120
	// if price == 100 {
	// 	fmt.Println("First case")
	// } else if price == 110 {
	// 	fmt.Println("Second case")
	// } else if price == 120 {
	// 	fmt.Println("Third price")
	// } ...

	//var price int
	//fmt.Scan(&price)

	// //for {
	// switch price { // Запрещены дубли в case'ах
	// case 100:
	// 	fmt.Println("First case")
	// 	//continue
	// case 200:
	// 	fmt.Println("Second case")
	// 	//continue
	// case 300:
	// 	fmt.Println("Third case")
	// 	//continue
	// case 400:
	// 	fmt.Println("Forth case")
	// 	//continue
	// default:
	// 	//Отработает  если выше не нашел в условиях
	// 	fmt.Println("Default case")
	// }

	//Case с множеством вариантов

	var price string = "a"
	switch price {
	case "a", "b":
		fmt.Println("A - B")
		//continue
	case "c", "d":
		fmt.Println("C - D")
		//continue
	case "e", "f":
		fmt.Println("F")
		//continue
	case "g", "j":
		fmt.Println("G - J")
		//continue
	default:
		//Отработает  если выше не нашел в условиях
		fmt.Println("Default case")
	}

	// Case можно использовать с вырожениями
	var age int
	fmt.Scan(age)

	switch {
	case age <= 18:
		fmt.Println("Too yong")
	case age > 18 && age <= 30:
		fmt.Println("Second case")
	case age > 30 && age <= 100:
		fmt.Println("You old")
	default:
		fmt.Println("Who are you")
	}

	//case c проваливаниями . Проваливаются даже ложные кейсы
	// В момент выполнения fallthrough у следующего кейсв условие не проверяется а ваполняется сразу
	//решение brack или outer:
	var number int
	fmt.Scan(&number)
	switch {
	case number < 100:
		fmt.Printf("%d is less then 100\n", number)
		if number%2 == 0 {
			break
		}
		fallthrough
	case number < 200:
		fmt.Printf("%d is less then 200\n", number)
		fallthrough
	default:
		fmt.Print("DEFAULT")
	}

	//Гадость терминации цикла for из switch
	// Используй лейблы при двух циклах, если один не надо
	var i int
uberloop:
	for {
		fmt.Scan(&i)
		switch {
		case i%2 == 0:
			fmt.Printf("Value is %d and it's even\n", i)
			break uberloop
		}
	}

	fmt.Println("END")
}
