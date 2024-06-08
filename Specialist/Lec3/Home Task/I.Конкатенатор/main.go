package main

import (
	"fmt"
)

func main() {
	var aArg, bArg, cArg int
	fmt.Scan(&aArg, &bArg, &cArg)
	fmt.Printf("%d%d%d", cArg, bArg, aArg)
}
