package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		result strings.Builder
	)
	//fmt.Printf("%s...", result)
	result.WriteString("s")
	result.WriteString("w")

	fmt.Print(result.String())

}
