package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	type User struct {
		ID        int64  `json:"id"`
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
	}

	u := User{}
	u.ID = 22
	u.Email = "test@test.com"
	u.FirstName = "John"

	bs, _ := json.Marshal(u)

	fmt.Println(string(bs)) // {"id":22,"email":"test@test.com","first_name":"John"}
}
