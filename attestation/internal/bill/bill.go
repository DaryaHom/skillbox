package bill

import (
	"fmt"
	"io/ioutil"
)

type BillingData struct {
	CreateCustomer bool `json:"create_customer"`
	Purchase       bool `json:"purchase"`
	Payout         bool `json:"payout"`
	Recurring      bool `json:"recurring"`
	FraudControl   bool `json:"fraud_control"`
	CheckoutPage   bool `json:"checkout_page"`
}

func NewBillingData() *BillingData {
	return &BillingData{}
}

func (d *BillingData) SetCreateCustomer(createCustomer bool) {
	d.CreateCustomer = createCustomer
}

func (d *BillingData) SetPurchase(purchase bool) {
	d.Purchase = purchase
}

func (d *BillingData) SetPayout(payout bool) {
	d.Payout = payout
}

func (d *BillingData) SetRecurring(recurring bool) {
	d.Recurring = recurring
}

func (d *BillingData) SetFraudControl(fraudControl bool) {
	d.FraudControl = fraudControl
}

func (d *BillingData) SetCheckoutPage(checkoutPage bool) {
	d.CheckoutPage = checkoutPage
}

// IsValid - checks the data validity
func IsValid(data byte) bool {
	if data == 49 {
		return true
	}
	return false
}

// GetStatus - function to get billing status data from CSV file
func GetStatus() (BillingData, error) {
	fmt.Println()
	fmt.Println("****************")
	fmt.Println("Billing Status:")

	var bill BillingData
	data, err := ioutil.ReadFile("../attestation/assets/billing.data")
	if err != nil {
		return bill, err
	}

	var store []bool
	for _, d := range data {
		store = append(store, IsValid(d))
	}

	bill.SetCheckoutPage(store[0])
	bill.SetFraudControl(store[1])
	bill.SetRecurring(store[2])
	bill.SetPayout(store[3])
	bill.SetPurchase(store[4])
	bill.SetCreateCustomer(store[5])

	// Testing the function
	fmt.Printf("%+v\n", bill)

	return bill, nil
}
