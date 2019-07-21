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

//CreateSubscription makes request to create subscription
func (sr *SubscriptionResponse) CreateSubscription(c *Customer, plan string) (statusCode int, conektaError Error) {
	statusCode, response := request("POST", "/customers/"+c.ID+"/subscription", body{"plan": plan})
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
func (sr *SubscriptionResponse) UpdateSubscription(c *Customer, plan string) (statusCode int, conektaError Error, subscription Subscription) {
	statusCode, response := request("PUT", "/customers/"+c.ID+"/subscription", body{"plan": plan})
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
func (sr *SubscriptionResponse) PauseSubscription(c *Customer) (statusCode int, conektaError Error, subscription Subscription) {
	statusCode, response := request("POST", "/customers/"+c.ID+"/subscription/pause", nil)
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
func (sr *SubscriptionResponse) ResumeSubscription(c *Customer) (statusCode int, conektaError Error, subscription Subscription) {
	statusCode, response := request("POST", "/customers/"+c.ID+"/subscription/resume", nil)
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
func (sr *SubscriptionResponse) CancelSubscription(c *Customer) (statusCode int, conektaError Error, subscription Subscription) {
	statusCode, response := request("POST", "/customers/"+c.ID+"/subscription/cancel", nil)
	if statusCode != 200 {
		err := json.Unmarshal(response, &conektaError)
		checkError(err)
	} else {
		err := json.Unmarshal(response, &sr)
		checkError(err)
	}
	return
}
