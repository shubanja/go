package main

import "fmt"

func main() {
	//1.Слайсы - он же срезы
	// Слайсы - это динамическая обвязка
	startArr := [4]int{10, 20, 30, 40}
	var startSlice []int = startArr[1:2] //Слайс иницил пустым квадрат скобками
	fmt.Println(startSlice)
	// Создали слайс основываясь уже на сужествование массива
	//2. Создание слайса без явной иниц массива
	secondSlice := []int{10, 20, 30, 40}
	fmt.Println(secondSlice)
	//3. Изменение элемента среза
	originArr := [...]int{10, 20, 30, 40, 50, 60}
	firstSlice := originArr[1:4] // [:] - весь массив
	for i, v := range firstSlice {
		firstSlice[i]++
		fmt.Println(i, v)
	}
	fmt.Println("originArr:", originArr)  // originArr: [10 21 31 41 50 60] Срезы это набор ссылок на элементы ниже лежащего массива
	fmt.Println("firstSlice", firstSlice) // firstSlice [21 31 41]
	//4. Один массив и два производных среза
	twoArr := [...]int{10, 20, 30, 40, 50, 60}
	fSlice := twoArr[:]
	sSlice := twoArr[2:4]
	fmt.Println("Before mod: Arr", twoArr, "fSlice", fSlice, "sSlice", sSlice)
	fSlice[3]++
	sSlice[0]++
	fmt.Println("Before mod: Arr", twoArr, "fSlice", fSlice, "sSlice", sSlice)
	//5. Срез как встроенный тип
	// type slice struct {
	// 	Length int
	// 	Capacity int
	// 	ZeroElement *byte
	// }

}
