package routing

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	// Auth Middleware
	auth := MerchantAuthMiddleware()

	// Order routes
	r.POST("/orders", auth, CreateOrder)

	// Payment routes
	r.POST("/payments", auth, CreatePayment)
	r.GET("/payments/:id", auth, GetPayment)

	// Provider routes
	r.POST("/provider", CreateProvider)

	// Merchant routes
	r.POST("/merchants", CreateMerchant)
}
