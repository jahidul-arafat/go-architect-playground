package main

import (
	"fmt"
)

// Not all pointers necessarily lead to heap allocation, and not all value types are always on the stack.
// create a struct type "User"
type User struct {
	Name string
	Age  int
}

// create a new type Named "UserEnhanced" of type "User" with the same fields as "User" and an additional field "Email"
type UserEnhanced struct {
	User
	Email string
}

// create a new type "MyUser" of "User" type and check if it can reuse the <User> methods
type MyUser User

// create few struct methods for the user
// Method-01 : value receiver
func (u User) PrintString() string {
	// check if the type of u is not of type User then it will panic
	//if reflect.TypeOf(u) == reflect.TypeOf(UserEnhanced{}) {
	//	return fmt.Sprintf("%s is %d, Email: %s", u.Name, u.Age,
	//		reflect.ValueOf(u).FieldByName("Email").String())
	//}
	return fmt.Sprintf("%s is %d", u.Name, u.Age)
}

// create a few independent functions - not methods
// call by value- will make a copy of the <user> in the function's local scope which will not be available after the function scope exists
// not available in the global scope like pointer
// will be stored in STACK (LIFO), not in HEAP
// no GC is required, will be automatically cleaned when function scope closes
func ChangeName(u *User) { // if pointer used as argumet, then it will attain a global lifecycle and will be at HEAP
	u.Name = "New Name"
	fmt.Println("New Name: ", u.Name)
}

func main() {
	fmt.Println("Go lang simulation with Pointer Parameters ....")

	// create a new user
	user1 := User{"Ally", 28}
	fmt.Println("Before: ", user1.PrintString())
	ChangeName(&user1)
	fmt.Println("After: ", user1.PrintString())

	// create a nw User of type <UserEnhanced>
	enhancedUser01 := UserEnhanced{
		User:  User{"eAlly", 20},
		Email: "eally@ally.com",
	}

	// copy th enhancedUser01 into a new variable of type <UserEnhanced>
	enhancedUser02 := enhancedUser01
	enhancedUser02.Name = "New Enhanced Name"

	// print the enhancedUser01 details
	// it will print the original details since the changes are not reflected in the copy enhancedUser02

	fmt.Printf("Enhanced User Detials: %+v\n", enhancedUser01)
	fmt.Println(enhancedUser01.User.PrintString())

	// create a new user of type "MyUser" which simply a clone of type <User>
	// even though its a clone of type <User>, since this is a new type, i cant reuse the methods of <User>
	myUser01 := MyUser{
		Name: "MyUser01",
		Age:  27,
	}

	fmt.Printf("My User Detials: %+v", myUser01)
	//myUser01.PrintString() // this PrintString() is mnethods for <User>, not for <MyUser>

}
