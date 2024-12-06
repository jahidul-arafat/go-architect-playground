package main

import (
	"fmt"
	"interface_polymorphism_payment/implementations"
)

func main() {
	fmt.Println("Payment Processing .... Interface ...Polymorphism")

	// create two instances of type CC and PayPal
	creditCardProcessor := implementations.CreditCardProcessor{}
	paypalProcessor := implementations.PayPalProcessor{}

	implementations.Checkout(creditCardProcessor, 32.44)
	implementations.Checkout(paypalProcessor, 67.90)

}
