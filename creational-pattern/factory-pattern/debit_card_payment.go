package factory_pattern

import "fmt"

type DebitCardPM struct {
	amount            float64
	paymentMethodType string
}

func NewDebitCardPM(amount float64) *DebitCardPM {
	return &DebitCardPM{
		amount:            amount,
		paymentMethodType: DebitCard,
	}
}

func (d *DebitCardPM) Pay() string {
	return fmt.Sprintf("Paid %0.1f via %s", d.amount, d.paymentMethodType)
}
