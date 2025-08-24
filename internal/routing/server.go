package routing

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	// Order routes
	r.POST("/orders", CreateOrder)
	r.GET("/orders/:id", OrderStatus)

	// Payment routes
	r.POST("/payments", CreatePayment)
	r.GET("/payments/:id", PaymentStatus)

	// Provider routes
	r.POST("/provider", CreateProvider)

	// Merchant routes
	r.POST("/merchants", CreateMerchant)
}
