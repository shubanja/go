package main

import "fmt"

func main() {
	var (
		aArg   int
		bArg   int
		perArg int
		squArg int
	)

	fmt.Scan(&aArg)
	fmt.Scan(&bArg)
	perArg = aArg*2 + bArg*2
	squArg = aArg * bArg
	fmt.Println("Периметр:", perArg)
	fmt.Println("Площадь:", squArg)

}
