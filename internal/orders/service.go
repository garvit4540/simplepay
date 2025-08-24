package orders

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

const Created string = "created"
const Completed string = "completed"

// OrdersService handles merchant-related operations
type OrdersService struct {
	repo *OrdersRepo
}

// NewOrdersService creates a new merchant service
func NewOrdersService(repo *OrdersRepo) *OrdersService {
	return &OrdersService{
		repo: repo,
	}
}

func (svc *OrdersService) ValidateOrder(ctx *gin.Context, order *OrderModel) error {
	if order.Amount <= 0 {
		return fmt.Errorf("amount must be greater than 0")
	}

	if order.MerchantID == "" {
		return fmt.Errorf("merchant_id is required")
	}
	
	if order.Currency == "" {
		return fmt.Errorf("currency is required")
	}

	basicAuthMerchantId, _ := ctx.Get("merchant_id")
	if order.MerchantID != basicAuthMerchantId {
		return fmt.Errorf("please use your own key id and secret to create order")
	}

	return nil
}

func (svc *OrdersService) CreateOrder(order *OrderModel) error {
	err := svc.repo.CreateOrder(order)
	if err != nil {
		return fmt.Errorf("order creation failed - %v", err)
	}

	return nil
}
