// main.go
//go:generate go run main.go -generate-schema

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/alecthomas/jsonschema"
)

// Define a new type Grade
type Grade int

const (
	A Grade = iota
	B
	C
	D
	F
)

type Student struct {
	ID    int
	Name  string
	Grade Grade
}

// GenerateSchema creates a JSON schema for the Student struct
func GenerateSchema() {
	schema := jsonschema.Reflect(&Student{})
	f, err := os.Create("student_schema.json")
	if err != nil {
		log.Fatalf("failed to create schema file: %v", err)
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(schema); err != nil {
		log.Fatalf("failed to encode schema: %v", err)
	}
	fmt.Println("JSON schema generated: student_schema.json")
}

func listStudents(studentList []Student) {
	fmt.Println("Listing Students Records:...")
	for _, student := range studentList {
		fmt.Printf("ID: %d, Name: %s, Grade: %s\n",
			student.ID, student.Name, student.Grade)
	}
}

func main() {
	// Parse flag to check if schema generation is requested
	generateSchema := flag.Bool("generate-schema", false, "Generate JSON schema for Student")
	flag.Parse()

	if *generateSchema {
		GenerateSchema()
		return
	}

	// Regular execution for listing students
	studentList := []Student{
		{ID: 1, Name: "Ally", Grade: A},
		{ID: 2, Name: "Billy", Grade: B},
		{ID: 3, Name: "Cilly", Grade: C},
	}

	listStudents(studentList)
}
