package factory_pattern

import "errors"

func GetPaymentMehtod(paymentMethodType string, amount float64) (IPaymentMethod, error) {
	switch paymentMethodType {
	case Cash:
		return NewCashPM(amount), nil
	case DebitCard:
		return NewDebitCardPM(amount), nil
	}
	return nil, errors.New("not implemented yet")
}
