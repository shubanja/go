// Hi

package main

import "fmt"

const (
	zero = iota
	one
	two
	three
)

const (
	a = iota
	b = 42
	c = iota
	d
)

func main() {
	fmt.Println(zero, one, two, three) // 0 1 2 3
	fmt.Println(a, b, c, d)            // 0 42 2 3
}
