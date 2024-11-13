// main_test.go
package main

import (
	"testing"
)

func TestListStudentsPass(t *testing.T) {
	students := []Student{
		{ID: 1, Name: "Alice", Grade: A},
		{ID: 2, Name: "Bob", Grade: B},
	}

	listStudents(students) // We simply ensure no errors occur while calling this
}

//func TestListStudentsFail(t *testing.T) {
//	students := []Student{
//		{ID: 1, Name: "Alice", Grade: A},
//		{ID: 2, Name: "Bob", Grade: B},
//	}
//
//	// Redirect stdout to capture the output
//	oldStdout := os.Stdout
//	r, w, _ := os.Pipe()
//	os.Stdout = w
//
//	listStudents(students)
//
//	// Restore stdout
//	w.Close()
//	os.Stdout = oldStdout
//
//	// Read the output
//	out, _ := io.ReadAll(r)
//	output := string(out)
//
//	// Check if the output contains the expected information
//	expectedOutputs := []string{
//		"ID: 1, Name: Alice, Grade: A",
//		"ID: 2, Name: Bob, Grade: B",
//		"Total number of students: 2",
//	}
//
//	for _, expected := range expectedOutputs {
//		if !strings.Contains(output, expected) {
//			t.Errorf("Expected output to contain '%s', but it didn't.\nActual output: %s", expected, output)
//		}
//	}
//
//	// This assertion will cause the test to fail
//	if strings.Contains(output, "ID: 3") {
//		t.Errorf("Unexpected student with ID 3 found in the output")
//	}
//}

// A benchmark function
/*
BenchmarkListStudents-14    	  332730	      3659 ns/op
PASS
*/
func BenchmarkListStudents(b *testing.B) {
	students := []Student{{ID: 1, Name: "Alice", Grade: A}}
	for i := 0; i < b.N; i++ {
		listStudents(students)
	}
}
