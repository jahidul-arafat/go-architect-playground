package implementations

import "interface_polymorphism_payment/interfaces"

/*
	•	Polymorphism: You can use a single function for multiple implementations.
	•	Flexibility: Code doesn’t need to know about concrete types, making it easier to extend.
	•	Testability: Mock implementations of the interface can be passed for testing.
*/
// Checkout
// function that accept the interface to behave like polymorphic
func Checkout(processorInterface interfaces.PaymentProcessor, amount float64) {
	processorInterface.ProcessPayment(amount)
}
