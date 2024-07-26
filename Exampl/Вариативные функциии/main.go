package main

import "fmt"

func main() {
	inStr := "Hello World"
	stringSlice := words(inStr)
	fmt.Println(stringSlice)
	//stringSlice[0] = `Aa` // [Aa]
	anotherSlice := []string{"four", "five", "six"}
	myWordsSlice := append(stringSlice, anotherSlice...)
	fmt.Println("String to slice:", myWordsSlice) //[Hello World four five six]
	fmt.Println()
	stringSlice2 := words2(inStr)
	fmt.Println("Print stringSlice2:", stringSlice2)
	fmt.Println()
	slice := append([]byte("hello "), "world"...)
	fmt.Println(slice)
}

func words(w ...string) []string {
	return w
}

func words2(w ...string) []string {
	fmt.Println("words2:", printWord(w...))
	return w
}

func printWord(w ...string) int {
	for id, word := range w {
		fmt.Println("id:", id, "printWord:", word)
	}
	return len(w)
}
