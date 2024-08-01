package main

import "fmt"

func main() {
	var firstStr, secondStr string
	var end, sec string

	fmt.Scan(&firstStr)

	for {
		fmt.Scan(&secondStr)
		end = fmt.Sprintf("%x ", firstStr[len(firstStr)-2:len(firstStr)])
		sec = fmt.Sprintf("%x ", secondStr[0:2])
		fmt.Println(end)
		fmt.Println(sec)
		if end != sec {
			fmt.Println("Проиграл:", secondStr)
			break
		} else {
			firstStr = secondStr
			fmt.Printf("firstStr: %s\n", firstStr)
			secondStr = ""
			fmt.Printf("secondStr: %s\n", secondStr)
		}
	}
}

func lenStr(s string) int {
	return len(s)
}
