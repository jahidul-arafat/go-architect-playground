package interfaces

// PaymentProcessor interface defines the behavior for processing payments
type PaymentProcessor interface {
	ProcessPayment(amount float64) //method to be implemented
}
