package main

import "fmt"

func main() {
	//1.Слайсы - он же срезы
	// Слайсы - это динамическая обвязка
	startArr := [4]int{10, 20, 30, 40}
	var startSlice []int = startArr[1:2] //Слайс иницил пустым квадрат скобками
	fmt.Println(startSlice)
	// Создали слайс основываясь уже на существование массива
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
	fmt.Println("Before mod: Arr", twoArr, "fSlice", fSlice, "sSlice", sSlice) //[10 20 30 40 50 60] fSlice [10 20 30 40 50 60] sSlice [30 40]
	fSlice[3]++
	sSlice[0]++
	fmt.Println("After mod: Arr", twoArr, "fSlice", fSlice, "sSlice", sSlice) //[10 20 31 41 50 60] fSlice [10 20 31 41 50 60] sSlice [31 41]
	//5. Срез как встроенный тип
	// type slice struct {
	// 	Length int
	// 	Capacity int
	// 	ZeroElement *byte
	// }

	//6. Длина и емкость слайса
	wordsSilce := []string{"one", "two", "three"}
	fmt.Println("slice:", wordsSilce, "Length:", len(wordsSilce), "Capacity:", cap(wordsSilce)) //slice: [one two three] Length: 3 Capacity: 3
	wordsSilce = append(wordsSilce, "four")
	fmt.Println("slice:", wordsSilce, "Length:", len(wordsSilce), "Capacity:", cap(wordsSilce)) //slice: [one two three four] Length: 4 Capacity: 6
	// Capacity (cap) или ёмкость слайса - это значение, показывающее СКОЛЬКО ЭЛЕМЕНТОВ В ПРИНЦИПЕ
	// можно добавить в срез БЕЗ ВЫДЕЛЕНИЯ ДОПОЛНИТЕЛЬНОЙ ПАМЯТИ ПОД НИЖЕЛЕЖАЩИЙ МАССИВ.

	// Допустим у нас есть срез на 3 элемента (инициализировали без явного указания массива)
	// Компилятор при создании этого среза СНАЧАЛА создал массив ровно на 3 элемента
	// После этого компилятор вернул адрес, где этот масив живет
	// Срез запомнил этот адрес и теперь ссылается на него
	// Потом
	// Начинаем деформировать слайс (увеличим длину /увеличим количество элементов)
	// Проблема - в нижележащем массиве (на котором основан слайс) все 3 ячейки. Что делать?
	// Компилятор ищет в памяти место для массива размера 3*2 (в общем случае n*2, где n - первоначальный размер)
	// После того как место найдено (в нашем случае найдено место для 6 элементов), в это место копируются
	// старые 3 элемента на свои позиции. На 4-ую позицию мы добавляем новый элемент
	// После этого компилятор возвращает нашему слайсу новый адрес в памяти, где находится массив под 6 элементов.

	//Емкость всегда будет изменять как n*2.
	numerics := []int{1, 2}
	for i := 0; i < 200; i++ {
		if i%5 == 0 {
			fmt.Println("Current len:", len(numerics), "Current cap:", cap(numerics))
		}
		numerics = append(numerics, i)
	}

	//Важно: после выделения памяти под новый массив, ссылки со старым будут перенсены в новый
	// Пример
	numArr := [2]int{1, 2}
	numSlice := numArr[:]

	numSlice = append(numSlice, 3) // В этот момент numSlice больше не ссылается на numArr
	numSlice[0] = 10
	fmt.Println(numArr)
	fmt.Println(numSlice)

	// 7. Как создавать слайсы наиболее эффективно?
	// make() - это функция, позволяющая более детально создавать срезы
	sl := make([]int, 10, 15)
	// []int - тип коллекции
	// 10 - длина
	// 15 - емкость
	//Сначала инициализируется arr = [15]int
	//Затем по нему делается срез arr[0:10]
	//После чего он возаращается
	fmt.Println(sl)
	// 8. Добавление элементов в СРЕЗ
	myWords := []string{"one", "two", "three"}
	fmt.Println("myWords:", myWords)
	anotherSlice := []string{"four", "five", "six"}
	myWords = append(myWords, anotherSlice...)
	myWords = append(myWords, "seven", "eight")
	fmt.Println("myWords:", myWords)
	//9. Многомерный срез
	mSlice := [][]int{
		{1, 2},
		{3, 4, 5, 6},
		{10, 20, 30},
		{},
	}
	fmt.Println(mSlice)

}
