package conekta

import (
	"encoding/json"
)

type PaymentSourceResponse struct {
	ID        string  `json:"id,omitempty"`
	Object    string  `json:"object,omitempty"`
	TokenID   string  `json:"token_id,omitempty"`
	Type      string  `json:"type,omitempty"`
	CreatedAt int64   `json:"created_at,omitempty"`
	Last4     string  `json:"last4,omitempty"`
	Name      string  `json:"name,omitempty"`
	ExpMonth  string  `json:"exp_month,omitempty"`
	ExpYear   string  `json:"exp_year,omitempty"`
	Brand     string  `json:"brand,omitempty"`
	ParentID  string  `json:"parent_id,omitempty"`
	Address   Address `json:"address,omitempty"`
	Deleted   bool    `json:"deleted,omitempty"`
}

type PaymentSource struct {
	ID        string  `json:"id,omitempty"`
	Object    string  `json:"object,omitempty"`
	TokenID   string  `json:"token_id,omitempty"`
	Type      string  `json:"type,omitempty"`
	CreatedAt int64   `json:"created_at,omitempty"`
	Last4     string  `json:"last4,omitempty"`
	Name      string  `json:"name,omitempty"`
	ExpMonth  string  `json:"exp_month,omitempty"`
	ExpYear   string  `json:"exp_year,omitempty"`
	Brand     string  `json:"brand,omitempty"`
	ParentID  string  `json:"parent_id,omitempty"`
	Address   Address `json:"address,omitempty"`
}

type PaymentSources struct {
	Object  string                   `json:"object,omitempty"`
	HasMore *bool                    `json:"has_more,omitempty"`
	Total   int64                    `json:"total,omitempty"`
	Data    []*PaymentSourceResponse `json:"data,omitempty"`
}

func (ps *PaymentSources) FindPaymentSources(customerID string) (statusCode int, conektaError ConektaError) {
	statusCode, response := request("GET", "/customers/"+customerID+"/payment_sources", nil)
	if statusCode != 200 {
		err := json.Unmarshal(response, &conektaError)
		checkError(err)
	} else {
		err := json.Unmarshal(response, &ps)
		checkError(err)
	}
	return
}

func (psr *PaymentSourceResponse) CreatePaymentSource(ps *PaymentSource) (statusCode int, conektaError ConektaError) {
	statusCode, response := request("POST", "/customers/"+ps.ParentID+"/payment_sources/", ps)
	if statusCode != 200 {
		err := json.Unmarshal(response, &conektaError)
		checkError(err)
	} else {
		err := json.Unmarshal(response, &psr)
		checkError(err)
	}
	return
}

func (psr *PaymentSourceResponse) DeletePaymentSource(customerID string, paymentSourceID string) (statusCode int, conektaError ConektaError) {
	statusCode, response := request("DELETE", "/customers/"+customerID+"/payment_sources/"+paymentSourceID, nil)
	if statusCode != 200 {
		err := json.Unmarshal(response, &conektaError)
		checkError(err)
	} else {
		err := json.Unmarshal(response, &psr)
		checkError(err)
	}
	return
}
