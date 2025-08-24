package routing

import (
	"github.com/garvit4540/simplepay/internal/merchants"
	"github.com/garvit4540/simplepay/internal/providerfactory"
	"github.com/garvit4540/simplepay/internal/registry"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOrder(c *gin.Context) {

}

func OrderStatus(c *gin.Context) {

}

func CreatePayment(c *gin.Context) {

}

func PaymentStatus(c *gin.Context) {

}

func CreateProvider(c *gin.Context) {
	providerService := registry.GetServiceFromRegister(registry.ProviderService)
	service, ok := providerService.(*providerfactory.ProviderService)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load provider service"})
	}

	service.CreateProvider(c)
}

func CreateMerchant(c *gin.Context) {
	merchantService := registry.GetServiceFromRegister(registry.MerchantService)
	service, ok := merchantService.(*merchants.MerchantService)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load merchants service"})
	}

	service.CreateMerchant(c)
}
