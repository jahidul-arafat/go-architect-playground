package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// create a struct User with JSON tag
// make sure the json to omit the field if it found to be empty

type User struct {
	ID        int       `json:"id"`                // non-reference field
	Username  string    `json:"username"`          // non-reference field
	Email     string    `json:"email"`             // non-reference field
	CreatedAT time.Time `json:"createdAT"`         // non-reference field
	IsActive  bool      `json:"isActive"`          // non-reference field
	Roles     []string  `json:"roles"`             // reference field // slice of strings
	IgnoreMe  string    `json:"-"`                 // telling the marshaller and unmarshaller to totally ignore this
	Profile   *Profile  `json:"profile,omitempty"` // pointing to the <Profile> struct// means this can be empty and if empty-> means a nil poiter
}

type Profile struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age,omitempty"` // omit the field if it has ZERO value
	Password  string `json:"-"`             // means, this field should be totally ignored during both marshaling/unmarshaling
}

// struct method-01
//
//	func (u User) PrintUserDetails() {
//		fmt.Printf("%+v\n", u)
//		if u.Profile != nil {
//			// u.Profile -> is an address the pointer points to
//			// *u.Profile -> is dereferencing the pointer to get the actual values
//			fmt.Printf("Profile: %+v\n", *u.Profile) // deferencing the pointer manually while printing
//		} else {
//			fmt.Println("Profile: nil")
//		}
//
// }
func (u User) PrintUserDetails() {
	fmt.Printf("ID: %d\n", u.ID)
	fmt.Printf("Username: %s\n", u.Username)
	fmt.Printf("Email: %s\n", u.Email)
	fmt.Printf("CreatedAT: %s\n", u.CreatedAT)
	fmt.Printf("IsActive: %v\n", u.IsActive)
	fmt.Printf("Roles: %v\n", u.Roles)
	fmt.Printf("IgnoreMe: %s (This field is ignored in JSON)\n", u.IgnoreMe)
	if u.Profile != nil {
		fmt.Printf("Profile:\n")
		fmt.Printf("  FirstName: %s\n", u.Profile.FirstName)
		fmt.Printf("  LastName: %s\n", u.Profile.LastName)
		fmt.Printf("  Age: %d\n", u.Profile.Age)
		fmt.Printf("  Password: %s (This field is ignored in JSON)\n", u.Profile.Password)
	} else {
		fmt.Println("Profile: nil")
	}
}

// a value to a pointer (*) should always be a memory address (&) storing the actual values

func main() {
	fmt.Println("Struct with JSONz Tag")

	// create a new user
	user := User{
		ID:        1,
		Username:  "jahidapon",
		Email:     "jahidapon@gmail.com",
		CreatedAT: time.Now(),
		IsActive:  true,
		Roles:     []string{"Admin"},
		IgnoreMe:  "IgnoreMe",
		Profile: &Profile{ // create a new Profile struct and takes the address of this newly created struct
			//Age:       30,
			FirstName: "Jahidul",
			LastName:  "Arafat",
			Password:  "12324",
		},
	}

	user.Roles = append(user.Roles, "NewRole")

	// can go automatically derefence values? -> Yes
	user.Profile.FirstName = "Test" // go is automatically dereferencing the pointer for you

	user.PrintUserDetails()

	/*
		When marshaling (converting struct to JSON), you pass the struct directly: json.Marshal(user)
		When unmarshaling (converting JSON to struct), you pass a pointer to the struct: json.Unmarshal(jsonData, &newUser)
	*/

	// Serealizing
	// Marshal the <user> to JSON
	// marshaling the user struct into JSON format (jsonData). This could be used to send the user data over a network or save it to a file.
	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error marshaling(serializing) the struct <user>!!")
		return // early return
	}
	// what a serialzied raw object looks like
	fmt.Printf("User JSON [Raw, Serialized []bytes]: %v\n", jsonData)
	fmt.Printf("User JSON: %v\n", string(jsonData))

	// Deserialize
	// Unmarshal JSON back to a User struct
	// unmarshaling this JSON data back into a new User struct (newUser). This simulates receiving JSON data and converting it back into a Go struct.
	var newUser User
	err = json.Unmarshal(jsonData, &newUser)
	if err != nil {
		fmt.Println("Error unmarshaling(deserializing) the JSON to struct....!!!")
		return
	}

	fmt.Println("\n\nUnmarshalled User:")
	newUser.PrintUserDetails()
}
