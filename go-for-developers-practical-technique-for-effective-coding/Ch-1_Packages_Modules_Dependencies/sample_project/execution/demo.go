package main

import (
	"fmt"
	"path"
)

func demo() {
	a := path.Join("a", "b")
	fmt.Println(a)
}

//func Join(root string, path string) string {
//	return path.Join(root, path)
//}
