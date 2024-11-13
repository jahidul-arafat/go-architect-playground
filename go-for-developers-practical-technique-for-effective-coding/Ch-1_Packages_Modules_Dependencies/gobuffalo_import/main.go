package main

import (
	"fmt"
	"github.com/gobuffalo/flect"
	"github.com/gofrs/uuid"
	"log"
)

func main() {
	s := "Hello World !"
	d := flect.Dasherize(s)
	fmt.Println(s, d)

	u, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u)
}
