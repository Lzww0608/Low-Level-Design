package atm

type Card struct {
	cardNumber    string
	pin           string
	accountNumber string
}

func NewCard(cardNumber, pin, accountNumber string) *Card {
	return &Card{
		cardNumber:    cardNumber,
		pin:           pin,
		accountNumber: accountNumber,
	}
}

func (c *Card) GetCardNumber() string {
	return c.cardNumber
}

func (c *Card) GetPIN() string {
	return c.pin
}

func (c *Card) GetAccountNumber() string {
	return c.accountNumber
}

func (c *Card) ValidatePIN(inputPIN string) bool {
	return c.pin == inputPIN
}
