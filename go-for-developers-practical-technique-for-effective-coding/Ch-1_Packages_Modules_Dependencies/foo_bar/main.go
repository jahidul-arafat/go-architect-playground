package main

import (
	bar1 "demo/bar"
	bar2 "demo/foo/bar"
	"fmt"
)

func main() {
	var p bar2.Pub
	var b bar1.Bar

	fmt.Println(p, b)
}
