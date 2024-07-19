package main

import (
	"fmt"
)

func main() {
	//Массивы основа
	//1. Определение массива
	//Создадим масив подхранение 5ти целочсленных элементов
	var arr [5]int
	fmt.Println("This is my array: ", arr)
	//2. Определение элементов массива (после предварительной инциализации)
	// Необходимо обратиться к элементу массива через синтаксис arr[i] = elem
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40
	arr[4] = 50
	fmt.Println("This is my array: ", arr)
	//3. Определение массва с указанием элементов на месте, в случаее если элементов меньше номинальной длины массива, элементы заполнятся нулями.
	newArr := [5]int{1, 2, 3}
	fmt.Println("This is my short array : ", newArr)
	//4. Создание массива через инициализацию переменных
	arrWithValues := [...]int{10, 20, 30, 40}
	fmt.Println("Arr declaration with [...]:", arrWithValues, "Len:", len(arrWithValues))
	arrWithValues[1] = 100
	fmt.Println("Arr declaration with [...]:", arrWithValues, "Len:", len(arrWithValues))
	//Массив это набор значений.копируется строго при присваивании
	first := [...]int{1, 2, 3}
	second := first
	second[0] = 10000
	fmt.Println("first:", first)
	fmt.Println("second:", second)
	//6. Массив и его размер это две составляющие  одного типа ( Размер массива - часть типа)
	//var aArr [5]int
	//var bArr [5]int
	//aArr[0] = 100
	//bArr = aArr
	//7. Интерирование по массиву (перебор масива с нумерацией) Одна из задач, справа на лево и как угодно.
	floatArr := [...]float64{12.5, 13.5, 10.2, 37.6, 19.84}
	for i := 1; i < len(arr); i++ {
		fmt.Printf("%d element of arr is: %.2f \n", i, floatArr[i])
	}
	//8.Итерирование по масиву через оператор range
	var sumArr float64
	for id, arr := range floatArr {
		fmt.Printf("%d element of arrRange is: %.2f \n", id, arr)
		sumArr += arr
	}
	fmt.Printf("Tottal sm is SumArr: %.2f \n", sumArr)
	// 9. Игнорирование id в randge for цикле, _ - blank identifier
	for _, arr := range floatArr {
		fmt.Printf("Element of arrRange is: %.2f \n", arr)
		sumArr += arr
	}
	fmt.Printf("Tottal sm is SumArr: %.2f \n", sumArr)
	// Многомерные массивы
	words := [2][2]string{
		{"Bob", "Alice"},
		{"Vicror", "Jo"},
	}
	fmt.Println("Multidimensional array ", words)
	//. Итерирование по многомерному массиву
	for _, val1 := range words {
		for _, val2 := range val1 {
			fmt.Printf(" %s", val2)
		}
		fmt.Println()
	}
}
