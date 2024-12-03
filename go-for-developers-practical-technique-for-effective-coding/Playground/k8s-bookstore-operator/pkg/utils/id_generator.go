package utils

import "math/rand"

// this is an utility function to generate the ID of all the struct objects
// Note- I am not gonna use it for now, parked for later use

// GenerateRandomID - generates an ID of length 8 with a prefix
/*
Type         Prefix
---------------------
Book		 BK_
BookStore	 BKS_
Student		 STD_
Librarian    LIB_
*/

func GenerateRandomID(prefix string) string {
	const CHARSET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	// allocate a slice of 8 bytes (since 8 character = 8 X 1Byte per character = 8 Byte)
	// create a slice of 8 bytes without initialization
	var id []byte
	return prefix + string(bytes)
}
