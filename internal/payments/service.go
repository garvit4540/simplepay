package payments

import (
	"fmt"
	"github.com/garvit4540/simplepay/internal/orders"
	"github.com/gin-gonic/gin"
)

const PaymentCreated string = "created"
const PaymentCompleted string = "completed"

type PaymentsService struct {
	repo       *PaymentsRepo
	ordersRepo *orders.OrdersRepo
}

func NewPaymentsService(repo *PaymentsRepo, ordersRepo *orders.OrdersRepo) *PaymentsService {
	return &PaymentsService{
		repo:       repo,
		ordersRepo: ordersRepo,
	}
}

func (svc *PaymentsService) ValidatePayment(ctx *gin.Context, payment *PaymentModel) error {
	if payment.OrderID == "" {
		return fmt.Errorf("order_id is required")
	}
	if payment.MerchantID == "" {
		return fmt.Errorf("merchant_id is required")
	}
	if payment.Amount <= 0 {
		return fmt.Errorf("amount must be greater than 0")
	}
	if payment.Currency == "" {
		return fmt.Errorf("currency is required")
	}

	// Check merchant ID from context
	basicAuthMerchantID, _ := ctx.Get("merchant_id")
	if payment.MerchantID != basicAuthMerchantID {
		return fmt.Errorf("please use your own key id and secret to create payment")
	}

	// Fetch the order and check its status
	order, err := svc.ordersRepo.GetOrderByID(payment.OrderID)
	if err != nil {
		return fmt.Errorf("failed to fetch order: %v", err)
	}

	if order.Status != orders.Created {
		return fmt.Errorf("payment can only be created for orders in 'created' state")
	}

	return nil
}

func (svc *PaymentsService) CreatePayment(payment *PaymentModel) error {
	err := svc.repo.CreatePayment(payment)
	if err != nil {
		return fmt.Errorf("payment creation failed - %v", err)
	}
	return nil
}

func (svc *PaymentsService) UpdatePayment(payment *PaymentModel) error {
	err := svc.repo.UpdatePayment(payment)
	if err != nil {
		return fmt.Errorf("payment update failed - %v", err)
	}
	return nil
}
