package keys

import (
	"fmt"
	"github.com/garvit4540/simplepay/internal/utils"
)

// KeysService handles merchant-related operations
type KeysService struct {
	repo *KeysRepo
}

// NewKeysService creates a new merchant service
func NewKeysService(repo *KeysRepo) *KeysService {
	return &KeysService{
		repo: repo,
	}
}

func (svc *KeysService) CreateKeysForMerchant(merchantId string) (string, error) {
	var keyValue string
	keyValue = utils.GenerateSimplePayID()
	keyModel := KeysModel{
		ID:         utils.GenerateSimplePayID(),
		MerchantId: merchantId,
		KeyValue:   keyValue,
	}
	err := svc.repo.CreateKeys(&keyModel)
	if err != nil {
		return "", fmt.Errorf("failed to create key for merchant %v - %w", merchantId, err)
	}
	return keyValue, nil
}

func (svc *KeysService) GetKeysForMerchant(merchantId string) (string, error) {
	keyModel, err := svc.repo.GetKeyByMerchantId(merchantId)
	if err != nil {
		return "", fmt.Errorf("failed to get key by merchant id %v - ", err)
	}
	return keyModel.KeyValue, nil
}
