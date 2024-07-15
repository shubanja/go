// Чтобы создать буферизованный считыватель, вы можете использовать bufio.NewReader функцию. Эта функция принимает io.Reader в качестве аргумента.
//Это означает, что вы можете передавать любой тип, реализующий io.Reader интерфейс. Сюда входят os.File, strings.Reader и bytes.Buffer.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("C:\\Users\\shubin\\Documents\\go\\go\\Exampl\\Bufio\\file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	//writer := bufio.NewWriter(file)
	//fmt.Println(writer)

	reader := bufio.NewReader(file)
	data := make([]byte, 100)
	_, err = reader.Read(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}
