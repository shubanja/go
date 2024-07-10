package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

outer:
	for {
		text, space := reader.ReadString(' ')
		fmt.Print(text)
		if space == nil {
			break outer
		}
	}
}
