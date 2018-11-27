package conekta

import (
	"context"
	"encoding/json"
)

type ChargeResponse struct {
	ID                  string        `json:"id,omitempty"`
	Object              string        `json:"object,omitempty"`
	Description         string        `json:"description,omitempty"`
	CreatedAt           int64         `json:"created_at,omitempty"`
	UpdatedAt           int64         `json:"updated_at,omitempty"`
	ExpiresAt           int64         `json:"expires_at,omitempty"`
	Currency            string        `json:"currency,omitempty"`
	Amount              int64         `json:"amount,omitempty"`
	MonthlyInstallments int64         `json:"monthly_installments,omitempty"`
	Livemode            *bool         `json:"livemode,omitempty"`
	Status              string        `json:"status,omitempty"`
	Fee                 int64         `json:"fee,omitempty"`
	OrderID             string        `json:"order_id,omitempty"`
	CustomerID          string        `json:"customer_id,omitempty"`
	DeviceFingerprint   string        `json:"device_fingerprint,omitempty"`
	PaidAt              int64         `json:"paid_at,omitempty"`
	PaymentMethod       PaymentMethod `json:"payment_method,omitempty"`
}

type Charge struct {
	ID                  string        `json:"id,omitempty"`
	Object              string        `json:"object,omitempty"`
	Description         string        `json:"description,omitempty"`
	CreatedAt           int64         `json:"created_at,omitempty"`
	UpdatedAt           int64         `json:"updated_at,omitempty"`
	ExpiresAt           int64         `json:"expires_at,omitempty"`
	Currency            string        `json:"currency,omitempty"`
	Amount              int64         `json:"amount,omitempty"`
	MonthlyInstallments int64         `json:"monthly_installments,omitempty"`
	Livemode            *bool         `json:"livemode,omitempty"`
	Status              string        `json:"status,omitempty"`
	Fee                 int64         `json:"fee,omitempty"`
	OrderID             string        `json:"order_id,omitempty"`
	CustomerID          string        `json:"customer_id,omitempty"`
	DeviceFingerprint   string        `json:"device_fingerprint,omitempty"`
	PaidAt              int64         `json:"paid_at,omitempty"`
	PaymentMethod       PaymentMethod `json:"payment_method,omitempty"`
}

type Charges struct {
	Object  string           `json:"object,omitempty"`
	HasMore *bool            `json:"has_more,omitempty"`
	Total   int64            `json:"total,omitempty"`
	Data    []ChargeResponse `json:"data,omitempty"`
}

func (cr *ChargeResponse) CreateCharge(ctx context.Context, c *Charge) (statusCode int, conektaError ConektaError) {
	statusCode, response := request(ctx, "POST", "/orders/"+c.OrderID+"/charges/", c)
	if statusCode != 200 {
		err := json.Unmarshal(response, &conektaError)
		checkError(err)
	} else {
		err := json.Unmarshal(response, &cr)
		checkError(err)
	}
	return
}
