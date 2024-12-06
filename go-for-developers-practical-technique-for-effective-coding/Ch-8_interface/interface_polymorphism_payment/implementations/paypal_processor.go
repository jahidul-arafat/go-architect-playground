package implementations

import "fmt"

type PayPalProcessor struct{}

// ProcessPayment
// Implement the method "ProcessPayment(amount float64)" from PaymentProcessor interface
func (pp PayPalProcessor) ProcessPayment(amount float64) {
	fmt.Printf("Processing PayPal payment of $%.2f\n", amount)
}
