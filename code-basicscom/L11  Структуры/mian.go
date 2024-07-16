package main

import "strings"

type UserCreateRequest struct {
	FirstName string // не может быть пустым; не может содержать пробелы
	Age       int    // не может быть 0 или отрицательным; не может быть больше 150
}

var invalRequest = "invalid request"

func main() {

}

func Validate(req UserCreateRequest) string {

	if req.FirstName == "" || strings.Contains(req.FirstName, " ") {
		return invalRequest
	}

	if req.Age == 0 || req.Age < 0 || req.Age > 150 {
		return invalRequest
	}

	return " "
}
