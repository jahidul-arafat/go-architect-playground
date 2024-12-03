package main

import (
	"fmt"
	"strings"
)

func main() {
	// this is the Entry point of the application

	//c := make([]string, 1, 3)           // length=1, capacity=3

	fmt.Println("k8s BookStore Operator v0.1")

	var slice1 []string
	slice1 = append(slice1, "I", "J")
	fmt.Printf("Slice: %[1]T, %#[1]v, Capacity: %#[3]v, length: %[2]v\n\n", slice1, len(slice1), cap(slice1))

	slice2 := make([]string, 0) // length=0
	slice2 = append(slice2, "A", "B", "C")
	fmt.Printf("Slice: %[1]T, %#[1]v, "+
		"Capacity: %#[2]v, Length: %#[3]v\n",
		slice2, cap(slice2), len(slice2))
	for _, item := range slice2 {
		fmt.Printf("item: %v\n", item)
	}

	fmt.Println("\nGenerating Random ID...", GenerateRandomID("ID-"))

}

func GenerateRandomID(prefix string) string {
	const CHARSET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	id := make([]string, 1)
	id = append(id, "A", "B")
	fmt.Println(id)

	return prefix + strings.Join(id, "")
}
