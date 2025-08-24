package routing

import (
	"context"
	"github.com/garvit4540/simplepay/internal/keys"
	"github.com/garvit4540/simplepay/internal/merchants"
	"github.com/garvit4540/simplepay/internal/orders"
	"github.com/garvit4540/simplepay/internal/providerfactory"
	"github.com/garvit4540/simplepay/internal/registry"
	"github.com/garvit4540/simplepay/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func MerchantAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		merchantID, secret, ok := c.Request.BasicAuth()
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization required"})
			return
		}

		service := registry.GetServiceFromRegister(registry.KeysService)
		keysService, ok := service.(*keys.KeysService)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load keys service"})
			return
		}

		merchantKey, err := keysService.GetKeysForMerchant(merchantID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load keys for this merchant in auth middleware"})
			return
		}

		if merchantKey != secret {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid key or secret"})
			return
		}

		c.Set("merchant_id", merchantID)
		c.Next()
	}
}

func CreateOrder(c *gin.Context) {
	service := registry.GetServiceFromRegister(registry.OrdersService)
	orderService, ok := service.(*orders.OrdersService)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load provider service"})
		return
	}

	_, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var order orders.OrderModel
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	if err := orderService.ValidateOrder(c, &order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload - " + err.Error()})
		return
	}

	order.ID = utils.GenerateSimplePayID()
	order.Status = orders.Created

	err := orderService.CreateOrder(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create order with err " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "order created successfully",
		"order":   order,
	})
	return
}

func CreatePayment(c *gin.Context) {

}

func GetPayment(c *gin.Context) {

}

func CreateProvider(c *gin.Context) {
	providerService := registry.GetServiceFromRegister(registry.ProviderService)
	service, ok := providerService.(*providerfactory.ProviderService)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load provider service"})
		return
	}

	_, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var provider providerfactory.ProviderModel
	if err := c.ShouldBindJSON(&provider); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	// generate random 10-digit ID
	provider.ID = utils.GenerateSimplePayID()

	err := service.CreateProvider(&provider)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create provider with error - " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "provider created successfully",
		"provider": provider,
	})
	return
}

func CreateMerchant(c *gin.Context) {
	merchantService := registry.GetServiceFromRegister(registry.MerchantService)
	service, ok := merchantService.(*merchants.MerchantService)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load merchants service"})
		return
	}

	service.CreateMerchant(c)
	return
}
