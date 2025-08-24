package provider_a

import (
	"encoding/json"
	"fmt"
	"github.com/garvit4540/simplepay/internal/payments"
	"math/rand/v2"
)

type ProviderAService struct{}

func NewProviderAService() *ProviderAService {
	return &ProviderAService{}
}

type ProviderASuccessResponse struct {
	TransactionID string `json:"transaction_id"`
	Status        string `json:"status"`
	Amount        int64  `json:"amount"`
	Currency      string `json:"currency"`
	Timestamp     string `json:"timestamp"`
}

type ProviderAErrorResponse struct {
	ErrorCode string `json:"error_code"`
	Message   string `json:"message"`
}

func (s *ProviderAService) ProcessPayment(payment *payments.PaymentModel) error {
	randomNumber := rand.IntN(10)

	// Mocking Provider A response here
	var respJSON []byte
	if randomNumber%2 == 0 {
		resp := ProviderASuccessResponse{
			TransactionID: "TXN123456",
			Status:        "APPROVED",
			Amount:        payment.Amount,
			Currency:      payment.Currency,
			Timestamp:     "2024-01-15T10:30:00Z",
		}
		respJSON, _ = json.Marshal(resp)
	} else {
		resp := ProviderAErrorResponse{
			ErrorCode: "INSUFFICIENT_FUNDS",
			Message:   "Not enough balance",
		}
		respJSON, _ = json.Marshal(resp)
	}

	payment.GatewayResponse = string(respJSON)

	var success ProviderASuccessResponse
	if err := json.Unmarshal(respJSON, &success); err == nil && success.Status == "APPROVED" {
		payment.Status = payments.PaymentCompleted
		return nil
	}

	var errResp ProviderAErrorResponse
	if err := json.Unmarshal(respJSON, &errResp); err == nil {
		payment.Status = payments.PaymentFailed
		return fmt.Errorf("provider a error: %s", errResp.Message)
	}

	return fmt.Errorf("unknown response from provider a")
}
