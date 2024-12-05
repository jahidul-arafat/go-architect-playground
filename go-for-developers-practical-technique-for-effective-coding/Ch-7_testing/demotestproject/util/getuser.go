package util

import "fmt"

type User struct {
	Name string
	Age  int
}

func GetUser(id int) (*User, error) {
	return nil, fmt.Errorf("Not implemented")
}
