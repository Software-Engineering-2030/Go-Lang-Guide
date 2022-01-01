package entity

func NewCreditCard1(number string, name string, expirationMonth int, expirationYear int, expirationCVV int) (*CreditCard, error) {
	cc := &CreditCard{
		number:          number,
		name:            name,
		expirationMonth: expirationMonth,
		expirationYear:  expirationYear,
		cvv:             expirationCVV,
	}
}
