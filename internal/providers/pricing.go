package providers

// Interface to providing pricing data
type Price interface {
	IProvider
	// GetExchange gets the value of currency1 in currency2
	GetExchange(currency1, currency2 string) (string, error)
}
