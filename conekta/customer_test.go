package conekta_test

import (
	"os"
	"testing"

	"github.com/grodriguez85/go-conekta/conekta"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCustomer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Create a customer")
}

// Public testing tokens available in
// https://developers.conekta.com/resources/testing
var _ = Describe("Creating customer", func() {
	//Testing key
	conekta.APIKey = os.Getenv("CONEKTAKEY")
	var cusid, paysrc string
	Context("Post customer", func() {
		//New customer
		customer := new(conekta.Customer)
		customer.Name = "Fulanito Pérez"
		customer.Email = "fulanito@conekta.com"
		customer.Phone = "+52181818181"
		//Testing payment src
		payment := conekta.PaymentSource{
			Type:    "card",
			TokenID: "tok_test_visa_4242",
		}
		customer.PaymentSources = append(customer.PaymentSources, payment)
		//Send to conekta
		customerResponse := new(conekta.CustomerResponse)
		statusCode, _ := customerResponse.Create(customer)
		cusid = customerResponse.ID
		It("Should response 200", func() {
			Expect(statusCode).Should(Equal(200))
		})
	})

	Context("Update a customer", func() {
		//New customer
		customer := new(conekta.Customer)
		customer.ID = cusid
		customer.Name = "Zutano Pérez"
		customer.Email = "zutano@conekta.com"
		customer.Phone = "+52181818181"
		customerResponse := new(conekta.CustomerResponse)
		statusCode, _ := customerResponse.Update(customer)
		It("Should response 200", func() {
			Expect(statusCode).Should(Equal(200))
		})
	})

	Context("Create a subscription", func() {
		customer := new(conekta.Customer)
		customer.ID = cusid
		subscriptionResponse := new(conekta.SubscriptionResponse)
		statusCode, _ := subscriptionResponse.CreateSubscription(customer, "399")
		It("Should response 200", func() {
			Expect(statusCode).Should(Equal(200))
		})
	})

	Context("Update a subscription", func() {
		customer := new(conekta.Customer)
		customer.ID = cusid
		subscriptionResponse := new(conekta.SubscriptionResponse)
		statusCode, _, _ := subscriptionResponse.UpdateSubscription(customer, "400")
		It("Should response 200", func() {
			Expect(statusCode).Should(Equal(200))
		})
	})

	Context("Pause a subscription", func() {
		customer := new(conekta.Customer)
		customer.ID = cusid
		subscriptionResponse := new(conekta.SubscriptionResponse)
		statusCode, _, _ := subscriptionResponse.PauseSubscription(customer)
		It("Should response 200", func() {
			Expect(statusCode).Should(Equal(200))
		})
	})

	Context("Resume a subscription", func() {
		customer := new(conekta.Customer)
		customer.ID = cusid
		subscriptionResponse := new(conekta.SubscriptionResponse)
		statusCode, _, _ := subscriptionResponse.ResumeSubscription(customer)
		It("Should response 200", func() {
			Expect(statusCode).Should(Equal(200))
		})
	})

	Context("Cancel a subscription", func() {
		customer := new(conekta.Customer)
		customer.ID = cusid
		subscriptionResponse := new(conekta.SubscriptionResponse)
		statusCode, _, _ := subscriptionResponse.CancelSubscription(customer)
		It("Should response 200", func() {
			Expect(statusCode).Should(Equal(200))
		})
	})

	Context("Create a payment source", func() {
		paymentSourceResponse := new(conekta.PaymentSourceResponse)
		statusCode, _ := paymentSourceResponse.CreatePaymentSource(&conekta.PaymentSource{
			Type:     "card",
			TokenID:  "tok_test_mastercard_4444",
			ParentID: cusid,
		})
		It("Should response 200", func() {
			Expect(statusCode).Should(Equal(200))
		})
	})

	Context("Delete a payment source", func() {
		paymentSourceResponse := new(conekta.PaymentSourceResponse)
		statusCode, _ := paymentSourceResponse.DeletePaymentSource(cusid, paysrc)
		It("Should response 200", func() {
			Expect(statusCode).Should(Equal(200))
		})
	})

	Context("Delete a customer", func() {
		customerResponse := new(conekta.CustomerResponse)
		statusCode, _ := customerResponse.Delete(cusid)
		It("Should response 200", func() {
			Expect(statusCode).Should(Equal(200))
		})
	})
})
