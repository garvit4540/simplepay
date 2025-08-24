package provider_b

import (
	"encoding/json"
	"fmt"
	"github.com/garvit4540/simplepay/internal/payments"
	"math/rand/v2"
)

type ProviderBService struct{}

func NewProviderBService() *ProviderBService {
	return &ProviderBService{}
}

type ProviderBSuccessResponse struct {
	PaymentID string `json:"paymentId"`
	State     string `json:"state"`
	Value     struct {
		Amount       string `json:"amount"`
		CurrencyCode string `json:"currencyCode"`
	} `json:"value"`
	ProcessedAt int64 `json:"processedAt"`
}

type ProviderBErrorResponse struct {
	ErrorType string `json:"errorType"`
	Reason    string `json:"reason"`
	Details   struct {
		Code string `json:"code"`
	} `json:"details"`
}

func (s *ProviderBService) ProcessPayment(payment *payments.PaymentModel) error {
	randomNumber := rand.IntN(10)

	// Mocking Provider B response here
	var respJSON []byte
	if randomNumber%2 == 0 {
		resp := ProviderBSuccessResponse{
			PaymentID: "PAY-789-XYZ",
			State:     "SUCCESS",
		}
		resp.Value.Amount = fmt.Sprintf("%.2f", float64(payment.Amount)/100)
		resp.Value.CurrencyCode = payment.Currency
		resp.ProcessedAt = 1705318200
		respJSON, _ = json.Marshal(resp)
	} else {
		resp := ProviderBErrorResponse{
			ErrorType: "PAYMENT_FAILED",
			Reason:    "Card Declined",
		}
		resp.Details.Code = "E001"
		respJSON, _ = json.Marshal(resp)
	}

	payment.GatewayResponse = string(respJSON)

	var success ProviderBSuccessResponse
	if err := json.Unmarshal(respJSON, &success); err == nil && success.State == "SUCCESS" {
		payment.Status = payments.PaymentCompleted
		return nil
	}

	var errResp ProviderBErrorResponse
	if err := json.Unmarshal(respJSON, &errResp); err == nil {
		payment.Status = payments.PaymentFailed
		return fmt.Errorf("provider b error: %s", errResp.Reason)
	}

	return fmt.Errorf("unknown response from provider b")
}
