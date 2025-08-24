package providerfactory

import (
	"math/rand/v2"
)

// ProviderService handles provider-related operations
type ProviderService struct {
	repo *ProviderRepo
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
