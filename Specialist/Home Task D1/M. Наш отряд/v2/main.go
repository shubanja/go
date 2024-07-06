package main

import (
	"bufio"
	"log"
	"os"
)

func correct(lines []string) (ok bool) {
	if len(lines) != 3 {
		return false
	} else if (lines[0] == "один" || lines[0] == "раз") && (lines[1] == "два") && (lines[2] == "три") {
		return true
	} else if (lines[0] == "1") && (lines[1] == "2") && (lines[2] == "3") {
		return true
	} else {
		return false
	}
}

func main() {
	var err error

	var lines = make([]string, 0, 3)
	var s = bufio.NewScanner(os.Stdin)
	for i := 0; s.Scan(); i++ {
		lines = append(lines, s.Text())
	}
	if err = s.Err(); err != nil {
		log.Fatalf("scanning error: %s\n", err)
	}

	if !correct(lines) {
		log.Fatalln("НЕ ПРАВИЛЬНО")
	}

	log.Println("ОК")
}
