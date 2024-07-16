package main

import (
	"fmt"
	"strings"
)

// UserCreateRequest is a request to create a new user.
type UserCreateRequest struct {
	FirstName string // не может быть пустым; не может содержать пробелы
	Age       int    // не может быть 0 или отрицательным; не может быть больше 150
}

var invalRequest = "invalid request"

func main() {
	u := UserCreateRequest{}
	u.FirstName = "Alex"
	u.Age = 18
	fmt.Println(Validate(u))
}

func Validate(req UserCreateRequest) string {

	if req.FirstName == "" || strings.Contains(req.FirstName, " ") {
		return invalRequest
	}

	if req.Age <= 0 || req.Age > 150 {
		return invalRequest
	}

	return ""
}
