package main

import "fmt"

func main() {
	var inputInt int = 10
	for id, arr := range inputInt {
		fmt.Printf("%d element of arrRange is: %.2f \n", id, arr)
		inputInt += arr
	}
}
