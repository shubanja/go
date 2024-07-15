//Реализуйте функцию ModifySpaces(s, mode string) string, которая изменяет строку s в зависимости от переданного режима mode. Подробности в примере:
//
//// когда передается mode "dash", нужно заменить все пробелы на дефисы "-"
//ModifySpaces("hello world", "dash") // hello-world
//
//// когда передается mode "underscore", нужно заменить все пробелы на нижние подчеркивания "_"
//ModifySpaces("hello world", "underscore") // hello_world
//
//// когда передается неизвестный или пустой mode, заменяем все пробелы на звездочки "*"
//ModifySpaces("hello world", "unknown") // hello*world
//ModifySpaces("hello world", "") // hello*world

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	var mode string
	fmt.Scanln(&mode)
	fmt.Println(ModifySpaces(line, mode))
}

func ModifySpaces(s, mode string) string {
	switch {
	case mode == "dash":
		return strings.ReplaceAll(s, " ", "-")
	case mode == "underscore":
		return strings.ReplaceAll(s, " ", "_")
	case mode == "unknown" || mode == "":
		return strings.ReplaceAll(s, " ", "*")
	}
	return s
}
