package test

import (
	utf "demotestproject/util" // <modulename>/<packagename>
	"testing"
)

/*
Several test execution commands
----------------------------------
 1081  go test ./...
 1082  go test -v ./test
 1083  go test -v ./...
 1084  go test -v ./..
 1085  g
 1086  go test ./...
 1087  go test -v ./test
 1088  go test -v ./test
 1089  go test -v -run="GetUser" ./...
 1090  go test -v -run="odd" ./...
 1091  go test -v -run="Odd" ./...
 1092  go test -v -run="Odd" ./test
 1093  go test -v -run="Logging" ./test
 1094  go test -v -run="User" ./test
 1095  go test -v -run="Infinite" ./test
 1096  go test -v -timeout 50ms -run="Infinite" ./test

*/

/*
Go test coverage report
1101  go test -cover ./test
 1102  go test -cover
 1103  go test -cover ./...
 1104  go test -coverprofile=coverage.out ./...
 1105  go tool cover -func=coverage.out
 1106  go tool cover -func=coverage.out
 1107  go test -coverprofile=coverage.out ./...
 1108  go tool cover -func=coverage.out
 1109  go test -coverprofile=coverage.out ./...
 1110  go tool cover -func=coverage.out
 1111  go tool cover -html=coverage.out
 1112  go tool cover -func=coverage.out
 1113  go tool cover -html=coverage.out
*/

// create a test for the isOdd() function
// test function takes a pointer to testing.T type
// why this pointer of testing.T type?
// to make test's failure, cleanup and even run subtests
func Test_IsOdd(t *testing.T) {
	// make the test as parallel
	// marking test as parallel will increase the spreed of the test
	t.Parallel()

	// testing
	if !utf.IsOdd(1) {
		t.Fatal("1 should be odd")
	}
	if utf.IsOdd(2) {
		t.Fatal("2 should not be odd")
	}

	if !utf.IsOdd(5) {
		t.Fatal("5 should be odd")
	}

}

// this test will fail since the User struct has no "id"
func Test_GetUser(t *testing.T) {
	user, err := utf.GetUser(1)
	if err != nil { // mark the test as failed with a error message
		t.Fatal(err) // Fatal method will stop the execution here
		// t.Error(err) // even if the test failed, the execution will not stop here
		// this will reach to user.Name where a nil cant nave a Name attribute; thus the system would panic

	}

	exp := "Ally"
	actual := user.Name // since user is NIL, user.Name would panic
	// this will never be reached if "Fatal" method is used
	if exp != actual {
		t.Fatalf("Expected: %v, got: %v", exp, actual)
	}

}

func Test_Logging(t *testing.T) {
	t.Log("This is a Test log message")
}

// create an infinite test
//func Test_Infinite(t *testing.T) {
//	for {
//		// infinite loop
//	}
//}
