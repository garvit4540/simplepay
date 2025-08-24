package providerfactory

import (
	"fmt"
	"github.com/garvit4540/simplepay/internal/payments"
	"github.com/garvit4540/simplepay/internal/providerfactory/provider_a"
	"github.com/garvit4540/simplepay/internal/providerfactory/provider_b"
	"math/rand/v2"
)

// ProviderService handles provider-related operations
type ProviderService struct {
	repo *ProviderRepo
}

type PaymentProvider interface {
	ProcessPayment(payment *payments.PaymentModel) error
}

// NewProviderService creates a new provider service
func NewProviderService(repo *ProviderRepo) *ProviderService {
	return &ProviderService{
		repo: repo,
	}
}

func (svc *ProviderService) CreateProvider(provider *ProviderModel) error {
	// insert into DB
	if err := svc.repo.CreateProvider(provider); err != nil {
		return err
	}
	return nil
}

func (svc *ProviderService) GetRandomProviderId() (string, error) {
	providers, err := svc.repo.GetAllProviders()
	if err != nil {
		return "", err
	}

	n := len(providers)
	idx := rand.IntN(n)

	return providers[idx].ID, nil
}

func (svc *ProviderService) GetPaymentProvider(provider string) (PaymentProvider, error) {
	switch provider {
	case "provider_a":
		return provider_a.NewProviderAService(), nil
	case "provider_b":
		return provider_b.NewProviderBService(), nil
	default:
		return nil, fmt.Errorf("unknown provider: %s", provider)
	}
}

func (svc *ProviderService) ProcessPaymentWithProvider(payment *payments.PaymentModel) error {
	providerId := payment.ProviderID
	provider, err := svc.repo.GetProviderByID(providerId)
	if err != nil {
		return err
	}

	providerConfig, err := svc.GetPaymentProvider(provider.Name)
	if err != nil {
		return err
	}

	return providerConfig.ProcessPayment(payment)
}
