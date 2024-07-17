package main

import "fmt"

func main() {
	var inMsg string
	fmt.Scanln(&inMsg)
	fmt.Print(ErrorMessageToCode(inMsg))
}

func ErrorMessageToCode(msg string) int {
	const (
		OK = iota
		CANCELLED
		UNKNOWN
	)
	switch msg {
	case "OK":
		return OK
	case "CANCELLED":
		return CANCELLED
	case "UNKNOWN":
		return UNKNOWN
	default:
		return UNKNOWN
	}
}
