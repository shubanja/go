package main

import (
	"fmt"
)

func main() {
	var aArg, bArg, cArg int
	fmt.Scan(&aArg, &bArg, &cArg)
	fmt.Printf("%d%d%d", cArg, bArg, aArg)

	var first string
	var second string
	var third string
	fmt.Scan(&first, &second, &third)
	concat := first + " " + second + " " + third
	fmt.Println("Full name:", concat)
}
