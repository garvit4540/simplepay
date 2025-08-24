package merchants

import (
	"context"
	"github.com/garvit4540/simplepay/internal/keys"
	"github.com/garvit4540/simplepay/internal/registry"
	"github.com/garvit4540/simplepay/internal/terminals"
	"github.com/garvit4540/simplepay/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// MerchantService handles merchant-related operations
type MerchantService struct {
	repo *MerchantRepo
}

// NewMerchantService creates a new merchant service
func NewMerchantService(repo *MerchantRepo) *MerchantService {
	return &MerchantService{
		repo: repo,
	}
}

func (svc *MerchantService) CreateMerchant(c *gin.Context) {
	_, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var merchant MerchantModel
	if err := c.ShouldBindJSON(&merchant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	// generate random 10-digit ID
	merchant.ID = utils.GenerateSimplePayID()

	// generate secret key for merchant
	service := registry.GetServiceFromRegister(registry.KeysService)
	keysService, ok := service.(*keys.KeysService)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch keys service from registry"})
	}
	merchantSecretKey, err := keysService.CreateKeysForMerchant(merchant.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create secret key for merchant"})
	}

	// Create a terminal for this merchant for future
	service = registry.GetServiceFromRegister(registry.TerminalService)
	terminalService, ok := service.(*terminals.TerminalService)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch terminal service from registry"})
	}
	err = terminalService.CreateTerminalForMerchant(merchant.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create terminal for merchant"})
	}

	// insert into DB
	if err := svc.repo.CreateMerchant(&merchant); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":             "merchant created successfully",
		"merchant":            merchant,
		"merchant_secret_key": merchantSecretKey,
	})
}
