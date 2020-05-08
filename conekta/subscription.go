package conekta

import (
	"encoding/json"
)

//SubscriptionResponse subscriptions api response struct
type SubscriptionResponse struct {
	ID                string `json:"id,omitempty"`
	Object            string `json:"object,omitempty"`
	CreatedAt         int64  `json:"created_at,omitempty"`
	UpdatedAt         int64  `json:"updated_at,omitempty"`
	PausedAt          int64  `json:"paused_at,omitempty"`
	BillingCycleStart int64  `json:"billing_cycle_start,omitempty"`
	BillingCycleEnd   int64  `json:"billing_cycle_end,omitempty"`
	TrialStart        int64  `json:"trial_start,omitempty"`
	TrialEnd          int64  `json:"trial_end,omitempty"`
	PlanID            string `json:"plan_id,omitempty"`
	Status            string `json:"status,omitempty"`
	CustomerID        string `json:"customer_id,omitempty"`
}

//Subscription subscription struct
type Subscription struct {
	ID                string `json:"id,omitempty"`
	Object            string `json:"object,omitempty"`
	CreatedAt         int64  `json:"created_at,omitempty"`
	UpdatedAt         int64  `json:"updated_at,omitempty"`
	PausedAt          int64  `json:"paused_at,omitempty"`
	BillingCycleStart int64  `json:"billing_cycle_start,omitempty"`
	BillingCycleEnd   int64  `json:"billing_cycle_end,omitempty"`
	TrialStart        int64  `json:"trial_start,omitempty"`
	TrialEnd          int64  `json:"trial_end,omitempty"`
	PlanID            string `json:"plan_id,omitempty"`
	Status            string `json:"status,omitempty"`
}

type SubscriptionRequestBody struct {
	Plan            string `json:"plan,omitempty"`
	CustomerID      string `json:"customer_id,omitempty"`
	PaymentSourceID string `json:"card,omitempty"`
}

//SubscriptionWebhook struct for webhooks payload
type SubscriptionWebhook struct {
	Data SubscriptionWebhookData `json:"data,omitempty"`
	Type string                  `json:"type,omitempty"`
}

type SubscriptionWebhookData struct {
	Subscription SubscriptionResponse `json:"object,omitempty"`
}

//CreateSubscription makes request to create subscription
func (sr *SubscriptionResponse) CreateSubscription(conektaID string, plan string, paymentSourceID interface{}) (statusCode int, conektaError Error) {

	body := SubscriptionRequestBody{
		Plan:       plan,
		CustomerID: conektaID,
	}

	if paymentSourceID != nil {
		body.PaymentSourceID = paymentSourceID.(string)
	}

	statusCode, response := request(
		"POST",
		"/customers/"+conektaID+"/subscription",
		body,
	)
	if statusCode != 200 {
		err := json.Unmarshal(response, &conektaError)
		checkError(err)
	} else {
		err := json.Unmarshal(response, &sr)
		checkError(err)
	}
	return
}

//UpdateSubscription makes request to update subscription
func (sr *SubscriptionResponse) UpdateSubscription(conektaID string, plan string) (statusCode int, conektaError Error) {

	body := SubscriptionRequestBody{
		Plan:       plan,
		CustomerID: conektaID,
	}

	statusCode, response := request("PUT", "/customers/"+conektaID+"/subscription", body)
	if statusCode != 200 {
		err := json.Unmarshal(response, &conektaError)
		checkError(err)
	} else {
		err := json.Unmarshal(response, &sr)
		checkError(err)
	}
	return
}

//PauseSubscription makes request to pause subscription
func (sr *SubscriptionResponse) PauseSubscription(conektaID string) (statusCode int, conektaError Error) {
	statusCode, response := request("POST", "/customers/"+conektaID+"/subscription/pause", nil)
	if statusCode != 200 {
		err := json.Unmarshal(response, &conektaError)
		checkError(err)
	} else {
		err := json.Unmarshal(response, &sr)
		checkError(err)
	}
	return
}

//ResumeSubscription makes request to resume subscription
func (sr *SubscriptionResponse) ResumeSubscription(conektaID string) (statusCode int, conektaError Error) {
	statusCode, response := request("POST", "/customers/"+conektaID+"/subscription/resume", nil)
	if statusCode != 200 {
		err := json.Unmarshal(response, &conektaError)
		checkError(err)
	} else {
		err := json.Unmarshal(response, &sr)
		checkError(err)
	}
	return
}

//CancelSubscription makes request to cancel subscription
func (sr *SubscriptionResponse) CancelSubscription(conektaID string) (statusCode int, conektaError Error) {
	statusCode, response := request("POST", "/customers/"+conektaID+"/subscription/cancel", nil)
	if statusCode != 200 {
		err := json.Unmarshal(response, &conektaError)
		checkError(err)
	} else {
		err := json.Unmarshal(response, &sr)
		checkError(err)
	}
	return
}
