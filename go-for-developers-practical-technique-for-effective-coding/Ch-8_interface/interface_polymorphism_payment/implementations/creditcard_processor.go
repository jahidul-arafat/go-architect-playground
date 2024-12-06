package implementations

import "fmt"

type CreditCardProcessor struct {
}

// ProcessPayment
// Implement the method "ProcessPayment(amount float64)" from PaymentProcessor interface
func (cc CreditCardProcessor) ProcessPayment(amount float64) {
	fmt.Printf("Processing credit card payment of $%.2f\n", amount)
}
