//Значение переменной можно изменять в любой момент:
//
//// двоеточие используется только при инициализации
//num := 22
//num = 33

package main

import (
	"fmt"
	"math"
)

func main() {
	aArg := 4
	bArg := 5
	fmt.Println(MinInt(aArg, bArg))
}

func MinInt(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}
