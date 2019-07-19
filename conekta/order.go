package conekta

import (
	"encoding/json"
)

type OrderResponse struct {
	Livemode        *bool           `json:"livemode,omitempty"`
	Amount          int64           `json:"amount,omitempty"`
	Currency        string          `json:"currency,omitempty"`
	PaymentStatus   string          `json:"payment_status,omitempty"`
	AmountRefunded  int64           `json:"amount_refunded,omitempty"`
	CustomerInfo    CustomerInfo    `json:"customer_info,omitempty"`
	ShippingContact ShippingContact `json:"shipping_contact,omitempty"`
	Object          string          `json:"object,omitempty"`
	ID              string          `json:"id,omitempty"`
	Metadata        Metadata        `json:"metadata,omitempty,omitempty"`
	CreatedAt       int64           `json:"created_at,omitempty"`
	UpdatedAt       int64           `json:"updated_at,omitempty"`
	LineItems       LineItems       `json:"line_items,omitempty"`
	ShippingLines   ShippingLines   `json:"shipping_lines,omitempty"`
	Charges         Charges         `json:"charges,omitempty"`
}

type Order struct {
	ID              string           `json:"id,omitempty"`
	Object          string           `json:"object,omitempty"`
	CreatedAt       int64            `json:"created_at,omitempty"`
	UpdatedAt       int64            `json:"updated_at,omitempty"`
	Currency        string           `json:"currency,omitempty"`
	LineItems       []LineItem       `json:"line_items,omitempty"`
	ShippingLines   []ShippingLine   `json:"shipping_lines,omitempty"`
	TaxLines        []TaxLine        `json:"tax_lines,omitempty"`
	DiscountLines   []DiscountLine   `json:"discount_lines,omitempty"`
	Livemode        *bool            `json:"livemode,omitempty"`
	PreAuthorize    *bool            `json:"pre_authorize,omitempty"`
	ShippingContact *ShippingContact `json:"shipping_contact,omitempty"`
	Amount          float64          `json:"amount,omitempty"`
	Reason          string           `json:"reason,omitempty"`
	AmountRefunded  float64          `json:"amount_refunded,omitempty"`
	PaymentStatus   string           `json:"payment_status,omitempty"`
	CustomerInfo    CustomerInfo     `json:"customer_info,omitempty"`
	Charges         []Charge         `json:"charges,omitempty"`
	Metadata        Metadata         `json:"metadata,omitempty"`
}

type TaxLine struct {
	ID          string   `json:"id,omitempty"`
	Object      string   `json:"object,omitempty"`
	Description string   `json:"description,omitempty"`
	Amount      float64  `json:"amount,omitempty"`
	ParentID    string   `json:"parent_id,omitempty"`
	Metadata    Metadata `json:"metadata,omitempty"`
}

type DiscountLine struct {
	ID       string   `json:"id,omitempty"`
	Object   string   `json:"object,omitempty"`
	Code     string   `json:"code,omitempty"`
	Type     string   `json:"type,omitempty"`
	Amount   int64    `json:"amount,omitempty"`
	ParentID string   `json:"parent_id,omitempty"`
	Metadata Metadata `json:"metadata,omitempty"`
}

type Tags map[string]interface{}

type Metadata map[string]interface{}

type PaymentMethod struct {
	Type                   string  `json:"type,omitempty"`
	TokenId                string  `json:"token_id,omitempty"`
	PaymentSourceID        string  `json:"payment_source_id,omitempty"`
	ServiceName            string  `json:"service_name,omitempty"`
	BarcodeURL             string  `json:"barcode_url,omitempty"`
	Object                 string  `json:"object,omitempty"`
	ExpiresAt              int64   `json:"expires_at,omitempty"`
	StoreName              string  `json:"store_name,omitempty"`
	Reference              string  `json:"reference,omitempty"`
	Name                   string  `json:"name,omitempty"`
	ExpMonth               string  `json:"exp_month,omitempty"`
	ExpYear                string  `json:"exp_year,omitempty"`
	AuthCode               string  `json:"auth_code,omitempty"`
	Last4                  string  `json:"last4,omitempty"`
	Brand                  string  `json:"brand,omitempty"`
	Issuer                 string  `json:"issuer,omitempty"`
	AccountType            string  `json:"account_type,omitempty"`
	Country                string  `json:"country,omitempty"`
	FraudScore             float64 `json:"fraud_score,omitempty"`
	ReceivingAccountBank   string  `json:"receiving_account_bank,omitempty"`
	ReceivingAccountNumber string  `json:"receiving_account_number,omitempty"`
}

// Creates a new Order
func (or *OrderResponse) Create(o *Order) (statusCode int, conektaError ConektaError) {
	statusCode, response := request("POST", "/orders", o)
	if statusCode != 200 {
		err := json.Unmarshal(response, &conektaError)
		checkError(err)
	} else {
		err := json.Unmarshal(response, &or)
		checkError(err)
	}
	return
}

// Updates an existing Order
func (or *OrderResponse) Update(o *Order) (statusCode int, conektaError ConektaError) {
	statusCode, response := request("PUT", "/orders/"+o.ID, o)
	if statusCode != 200 {
		err := json.Unmarshal(response, &conektaError)
		checkError(err)
	} else {
		err := json.Unmarshal(response, &or)
		checkError(err)
	}
	return
}

// Process a pre-authorized order.
func (or *OrderResponse) Capture(orderID string) (statusCode int, conektaError ConektaError) {
	statusCode, response := request("POST", "/orders/"+orderID+"/capture", nil)
	if statusCode != 200 {
		err := json.Unmarshal(response, &conektaError)
		checkError(err)
	} else {
		err := json.Unmarshal(response, &or)
		checkError(err)
	}
	return
}

// A Refund details the amount and reason why an order was refunded.
func (or *OrderResponse) Refund(o *Order) (statusCode int, conektaError ConektaError) {
	statusCode, response := request("POST", "/orders/"+o.ID+"/refunds", o)
	if statusCode != 200 {
		err := json.Unmarshal(response, &conektaError)
		checkError(err)
	} else {
		err := json.Unmarshal(response, &or)
		checkError(err)
	}
	return
}
