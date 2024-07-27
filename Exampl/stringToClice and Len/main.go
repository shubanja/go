package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	mySlice := strToSlice(line)
	fmt.Println("Slice:", mySlice)
	for i := 0; i < len(mySlice); i++ {
		fmt.Println("id:", i, "Element:", mySlice[i])
	}

}

func strToSlice(s string) []string {
	x := strings.Split(strings.TrimSpace(s), " ")
	fmt.Println(lenSlice(x))
	return x
}

func lenSlice(s []string) int {
	fmt.Println("Len:")
	return len(s)
}
