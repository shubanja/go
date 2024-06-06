package main

import (
	"fmt"
	"math"
)

func main() {
	var a int
	var b int
	fmt.Scan(&a)
	fmt.Scan(&b)

	fmt.Println(math.Pow(float64(a+b), 2))

}
