package providerfactory

import (
	"fmt"
	"time"
)

// BaseProvider provides common functionality for all providers
type BaseProvider struct {
	ID   string
	Name string
}

// GetName returns the provider name
func (bp *BaseProvider) GetName() string {
	return bp.Name
}

// GetID returns the provider ID
func (bp *BaseProvider) GetID() string {
	return bp.ID
}

// ProviderOneImpl implementation
type ProviderOneImpl struct {
	BaseProvider
	APIKey    string
	SecretKey string
	Endpoint  string
}

// NewProviderOne creates a new ProviderOneImpl instance
func NewProviderOne(id, apiKey, secretKey, endpoint string) *ProviderOneImpl {
	return &ProviderOneImpl{
		BaseProvider: BaseProvider{
			ID:   id,
			Name: string(ProviderOne),
		},
		APIKey:    apiKey,
		SecretKey: secretKey,
		Endpoint:  endpoint,
	}
}

// ProcessPayment processes a payment through ProviderOneImpl
func (p1 *ProviderOneImpl) ProcessPayment(amount int64, currency string, orderID string) error {
	// Simulate payment processing
	fmt.Printf("Processing payment through ProviderOne: Amount=%d, Currency=%s, OrderID=%s\n",
		amount, currency, orderID)

	// Here you would implement the actual payment processing logic
	// For now, we'll just simulate a successful payment
	time.Sleep(100 * time.Millisecond) // Simulate API call

	return nil
}

// ValidatePayment validates a payment through ProviderOneImpl
func (p1 *ProviderOneImpl) ValidatePayment(paymentID string) error {
	// Simulate payment validation
	fmt.Printf("Validating payment through ProviderOne: PaymentID=%s\n", paymentID)

	// Here you would implement the actual payment validation logic
	// For now, we'll just simulate a successful validation
	time.Sleep(50 * time.Millisecond) // Simulate API call

	return nil
}

// ProviderTwoImpl implementation
type ProviderTwoImpl struct {
	BaseProvider
	ClientID     string
	ClientSecret string
	BaseURL      string
}

// NewProviderTwo creates a new ProviderTwoImpl instance
func NewProviderTwo(id, clientID, clientSecret, baseURL string) *ProviderTwoImpl {
	return &ProviderTwoImpl{
		BaseProvider: BaseProvider{
			ID:   id,
			Name: string(ProviderTwo),
		},
		ClientID:     clientID,
		ClientSecret: clientSecret,
		BaseURL:      baseURL,
	}
}

// ProcessPayment processes a payment through ProviderTwoImpl
func (p2 *ProviderTwoImpl) ProcessPayment(amount int64, currency string, orderID string) error {
	// Simulate payment processing
	fmt.Printf("Processing payment through ProviderTwo: Amount=%d, Currency=%s, OrderID=%s\n",
		amount, currency, orderID)

	// Here you would implement the actual payment processing logic
	// For now, we'll just simulate a successful payment
	time.Sleep(150 * time.Millisecond) // Simulate API call

	return nil
}

// ValidatePayment validates a payment through ProviderTwoImpl
func (p2 *ProviderTwoImpl) ValidatePayment(paymentID string) error {
	// Simulate payment validation
	fmt.Printf("Validating payment through ProviderTwo: PaymentID=%s\n", paymentID)

	// Here you would implement the actual payment validation logic
	// For now, we'll just simulate a successful validation
	time.Sleep(75 * time.Millisecond) // Simulate API call

	return nil
}

// InitializeDefaultProviders creates and registers default providers with the factory
func InitializeDefaultProviders(factory *ProviderFactory) {
	// Create default ProviderOne
	providerOne := NewProviderOne(
		"PROV001",
		"pk_test_provider_one_key",
		"sk_test_provider_one_secret",
		"https://api.providerone.com/v1",
	)

	// Create default ProviderTwo
	providerTwo := NewProviderTwo(
		"PROV002",
		"client_id_provider_two",
		"client_secret_provider_two",
		"https://api.providertwo.com/v2",
	)

	// Register providers with the factory
	factory.RegisterProvider(ProviderOne, providerOne)
	factory.RegisterProvider(ProviderTwo, providerTwo)
}
