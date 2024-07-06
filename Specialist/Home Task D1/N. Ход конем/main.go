package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2 uint
	fmt.Scan(&x1, &y1, &x2, &y2)
	hodX := math.Abs(float64(x1) - float64(x2))
	hodY := math.Abs(float64(y1) - float64(y2))
	fmt.Println(hodX, hodY)
	if hodX == 1 && hodY == 2 || hodX == 2 && hodY == 1 {
		fmt.Println("ДА")
	} else {
		fmt.Println("НЕТ")
	}
}
