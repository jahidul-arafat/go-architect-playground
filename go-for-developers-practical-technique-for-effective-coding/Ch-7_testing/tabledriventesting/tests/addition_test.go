package tests

import (
	utf "tabledriventesting/utils"
	"testing"
)

// Helper function to test additional logic
// testing.TB interface allows this helper function to be used in both *testing.T (unit testing) and *testing.B(benchmarks testing)
func testAddHelper(t testing.TB, inp1, inp2, expected int) {
	// mark this is a helper function for the calling test function
	t.Helper() // without it, if the test is failed, the compiler will print the line number of this testAddHelper()
	// if used, when the test failed, the compiler will point the line number of the caller of this testAddHelper()

	// define the business logic this helper function does
	result := utf.Add(inp1, inp2) // call the Add() function to be tested
	if result != expected {
		t.Fatalf("[Inside TestHelper]>> Add(%d,%d)=%d[Got]; expected: %d", inp1, inp2, result, expected)
	}

	// test helper should cleanup some resources when done

}

func Test_Add(t *testing.T) {
	// make the tests running parallels
	t.Parallel()
	// define the test cases in a table
	// define the structure of the table
	type TestTableForAddition struct {
		testTitle string
		input1    int
		input2    int
		expected  int
	}

	// define all the testcases
	// create a new slice of tests - it's not anonymous since a Struct is defined earlier
	// if the struct is not defined earlier and is in call definition, then this struct would be an anonymous struct
	testCases := []TestTableForAddition{
		{"Adding two positives", 2, 3, 5},
		{"Adding positive and negative", 10, -5, 5},
		{"Adding two negatives", -5, -5, -20},
		{"Adding with zero", 0, 5, 5},
	}

	// iterate over each testcases
	for _, tc := range testCases {
		// create a subtest for each test cases
		// t.Run(param1/test_title, param2/ func(t *testing.T))
		// using t.Run() we are creating a subtest for each tests to easily identify which test has failed
		t.Run(tc.testTitle, func(t *testing.T) {
			// this code snippet has been replaced by a test helper function
			//result := utf.Add(tc.input1, tc.input2) // call the Add() function to be tested
			//if result != tc.expected {
			//	t.Fatalf("Add(%d,%d)=%d[Got]; expected: %d", tc.input1, tc.input2, result, tc.expected)
			//}
			testAddHelper(t, tc.input1, tc.input2, tc.expected)

		})
	}
}
