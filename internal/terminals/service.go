package terminals

import (
	"fmt"
	"github.com/garvit4540/simplepay/internal/providerfactory"
	"github.com/garvit4540/simplepay/internal/registry"
	"github.com/garvit4540/simplepay/internal/utils"
)

// TerminalService handles merchant-related operations
type TerminalService struct {
	repo *TerminalRepo
}

// NewTerminalService creates a new merchant service
func NewTerminalService(repo *TerminalRepo) *TerminalService {
	return &TerminalService{
		repo: repo,
	}
}

func (svc *TerminalService) CreateTerminalForMerchant(merchantId string) error {
	terminalId := utils.GenerateSimplePayID()

	// Generate Random Provider id
	service := registry.GetServiceFromRegister(registry.ProviderService)
	providerService, ok := service.(*providerfactory.ProviderService)
	if !ok {
		return fmt.Errorf("error converting to provider service")
	}
	providerId, err := providerService.GetRandomProviderId()
	if err != nil {
		return fmt.Errorf("failed fetching random provider err - %v", err)
	}

	terminalModel := &TerminalModel{
		ID:         terminalId,
		MerchantId: merchantId,
		ProviderId: providerId,
	}
	err = svc.repo.CreateTerminal(terminalModel)
	if err != nil {
		return fmt.Errorf("failed to create terminal for merchant %v - %w", merchantId, err)
	}

	return nil
}
