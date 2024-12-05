package tests

import (
	"encoding/json"
	"os"
	utf "tabledrivendesignandcleanup/utils"
	"testing"
)

// Define the struct for the table data
// +marshalling(serializing) and unmarshalling(deserializing) principles
// serializing: convert the struct data to JSON
// deserializing: convert the JSON data back to struct
type AddTable struct {
	Title    string `json:"title,omitempty"` // means, if title is NULL, omit that from the JSON too
	A        int    `json:"a"`               // means serialized to 'a' in json file
	B        int    `json:"b"`               // serialized to 'b' in json file
	Expected int    `json:"expected"`        // serialized to 'expected' in json file
}

// Helper-02
func testAddHelperController(t testing.TB, loadedCases []AddTable) {
	t.Helper()
	for _, tc := range loadedCases {
		result := utf.Add(tc.A, tc.B)
		t.Logf("Testing: %v,Got:[%v], Expected:[%v]", tc.Title, result, tc.Expected)
		if result != tc.Expected {
			t.Errorf("[%s] >> Add(%d,%d)=%d[Got], expected: %d",
				tc.Title, tc.A, tc.B, result, tc.Expected)
		}
	}
}

// Helper-01
// create the test helper function
func testAddHelperWithFile(t testing.TB, tmpFile string, testCases []AddTable) {
	// tag it as helper function
	t.Helper()

	// Create the tmpFile and make sure it cleanup after all the test case execution done
	file, err := os.Create(tmpFile)
	if err != nil {
		t.Fatalf("Failed to crate test case file: %v", tmpFile)
	}
	// now schedule the file cleanup after all the test cases completed
	// use annonymous function inside the call
	t.Cleanup(func() {
		os.Remove(tmpFile)
	})

	// write the test cases to the file
	// create a JSON encoder/serializer to serializer the data in the testCases into the file
	encoder := json.NewEncoder(file)
	// make sure the err variable is scoped only to the if statement and its block
	// this prevents the error variable from leaking into the boarder scope, which can help avoid accidental use of an outdated error value
	if err := encoder.Encode(testCases); err != nil { // this is a "common ok" or "commom error" idoms
		t.Fatalf("Failed to write test data to file: %v", err)
	}
	// close the file 'tmpFile'
	file.Close()

	// Unmarshalling
	// read the contents from 'tmpFile'
	data, err := os.ReadFile(tmpFile)
	if err != nil {
		t.Fatalf("Failed to unmarshal(deserialzie) test data from file: %v: %v", tmpFile, err)
	}

	// create a variable of type AddTable to load the unmarshalled data
	var loadedCases []AddTable

	// now we have to unmarshal these json data into a struct
	if err := json.Unmarshal(data, &loadedCases); err != nil {
		t.Fatalf("Failed to unmarshal test data: %v", err)
	}

	testAddHelperController(t, loadedCases)

}

// define the test case for the Add() function to test with a bundle of data written into the 'test_data.json' file
// the 'test_data.json' file will be automatically created during the execution
// and will be automatically deleted after the execution
// while cleaning up the `test_data.json` make sure the file is only deleted after all the test cases execution completed
// parameter: *testing.T pointer; this has all the required testing functionalities
func Test_AddWithFileToBeCleanedAfter(t *testing.T) {
	// make sure the test running parallel to fasten the test executions
	t.Parallel()

	// create a slice of test cases
	// aka. Table of test cases
	testCases := []AddTable{
		{"Adding two positives", 2, 3, 5},
		{"Adding positive and negative", 10, -5, 15},
		{"Adding two negatives", -2, -3, -5},
		{"Adding with zero", 0, 5, 5},
	}

	// pass the struct to the helper function to create the `json`
	// + to do necessary cleanup with TB interfaces (this interface hosting both the Testing.T and Testing.B)
	testAddHelperWithFile(t, "test_data.json", testCases)
}
