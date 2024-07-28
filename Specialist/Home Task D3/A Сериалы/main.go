package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	KekFlexFile := "C:\\Users\\MENDOZA\\Go\\src\\github\\go\\Specialist\\Home Task D3\\A Сериалы\\KekFlex.txt"
	myFilmsFile := "C:\\Users\\MENDOZA\\Go\\src\\github\\go\\Specialist\\Home Task D3\\A Сериалы\\MyFilms.txt"

	KekFlexSlice := readFile(KekFlexFile)
	myFilmsSlice := readFile(myFilmsFile)

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	nStr := input.Text()
	input.Scan()
	mStr := input.Text()

	n, _ := strconv.Atoi(nStr)
	m, _ := strconv.Atoi(mStr)

	nlist := KekFlexSlice[0:n]
	fmt.Println(nlist)
	mList := myFilmsSlice[0:m]
	fmt.Println(mList)

	fmt.Println(isArrayEquals(nlist, mList))

}
func isArrayEquals(arr1 []string, arr2 []string) bool {
	if len(arr1) < len(arr2) {
		fmt.Println("arr2 больше")
		return false
	}
	for _, v := range arr1 {
		for x := 0; x < len(arr2); x++ {
			if v == arr2[x] {
				fmt.Printf("ЕСТЬ %s\n", v)
				return true
			}

		}
		fmt.Printf("НЕТ В НАЛИЧИИ %s\n", v)
	}
	fmt.Println("НЕТ В НАЛИЧИИ")
	return false
}

func readFile(f string) []string {
	file, err := os.Open(f)
	FileSlice := make([]string, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		x := strings.Split(strings.TrimSpace(scanner.Text()), "\n")
		FileSlice = append(FileSlice, x[0])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return FileSlice
}
