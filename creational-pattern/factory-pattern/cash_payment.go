package factory_pattern

import "fmt"

type CashPM struct {
	amount            float64
	paymentMethodType string
}

func NewCashPM(amount float64) *CashPM {
	return &CashPM{
		amount:            amount,
		paymentMethodType: Cash,
	}
}

func (c *CashPM) Pay() string {
	return fmt.Sprintf("Paid %0.1f via %s", c.amount, c.paymentMethodType)
}
