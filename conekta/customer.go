package conekta

import (
	"encoding/json"
)

type CustomerResponse struct {
	CustomerID               string         `json:"id,omitempty"`
	Name                     string         `json:"name,omitempty"`
	Phone                    string         `json:"phone,omitempty"`
	Email                    string         `json:"email,omitempty"`
	Corporate                bool           `json:"corporate,omitempty"`
	DefaultPaymentSourceID   string         `json:"default_payment_source_id,omitempty"`
	DefaultShippingContactID string         `json:"default_shipping_contact_id,omitempty"`
	PaymentSources           PaymentSources `json:"payment_sources,omitempty"`
}

type Customer struct {
	CustomerID               string            `json:"id,omitempty"`
	Name                     string            `json:"name,omitempty"`
	Phone                    string            `json:"phone,omitempty"`
	Email                    string            `json:"email,omitempty"`
	Corporate                bool              `json:"corporate,omitempty"`
	DefaultPaymentSourceID   string            `json:"default_payment_source_id,omitempty"`
	DefaultShippingContactID string            `json:"default_shipping_contact_id,omitempty"`
	PaymentSources           []PaymentSource   `json:"payment_sources,omitempty"`
	ShippingContacts         []ShippingContact `json:"shipping_contacts,omitempty"`
}

type ShippingContact struct {
	ID             string   `json:"id,omitempty"`
	Object         string   `json:"object,omitempty"`
	CreatedAt      int64    `json:"created_at,omitempty"`
	UpdatedAt      int64    `json:"updated_at,omitempty"`
	Phone          string   `json:"phone,omitempty"`
	Receiver       string   `json:"receiver,omitempty"`
	BetweenStreets string   `json:"between_streets,omitempty"`
	Address        Address  `json:"address,omitempty"`
	Metadata       Metadata `json:"metadata,omitempty"`
}

type Address struct {
	Street1     string `json:"street1,omitempty"`
	Street2     string `json:"street2,omitempty"`
	City        string `json:"city,omitempty"`
	State       string `json:"string,omitempty"`
	Country     string `json:"country,omitempty"`
	PostalCode  string `json:"postal_code,omitempty"`
	Residential bool   `json:"residential,omitempty"`
	Object      string `json:"object,omitempty"`
}

func (cr *CustomerResponse) Find(id string) (statusCode int, conektaError ConektaError) {
	statusCode, response := request("GET", "/customers/"+id, nil)
	if statusCode != 200 {
		err := json.Unmarshal(response, &conektaError)
		checkError(err)
	} else {
		err := json.Unmarshal(response, &cr)
		checkError(err)
	}
	return
}

func (cr *CustomerResponse) Create(c *Customer) (statusCode int, conektaError ConektaError) {
	statusCode, response := request("POST", "/customers", c)
	if statusCode != 200 {
		err := json.Unmarshal(response, &conektaError)
		checkError(err)
	} else {
		err := json.Unmarshal(response, &cr)
		checkError(err)
	}
	return
}

func (cr *CustomerResponse) Update(c *Customer) (statusCode int, conektaError ConektaError) {
	statusCode, response := request("PUT", "/customers/"+c.CustomerID, c)
	if statusCode != 200 {
		err := json.Unmarshal(response, &conektaError)
		checkError(err)
	} else {
		err := json.Unmarshal(response, &cr)
		checkError(err)
	}
	return
}

func (cr *CustomerResponse) Delete(customerID string) (statusCode int, conektaError ConektaError) {
	statusCode, response := request("DELETE", "/customers/"+customerID, nil)
	if statusCode != 200 {
		err := json.Unmarshal(response, &conektaError)
		checkError(err)
	} else {
		err := json.Unmarshal(response, &cr)
		checkError(err)
	}
	return
}
