package atm

import "sync"

type BankingService struct {
	accounts sync.Map // key: string, value: *Account
	cards    sync.Map // key: string (cardNumber), value: *Card
}

func NewBankingService() *BankingService {
	return &BankingService{
		accounts: sync.Map{},
		cards:    sync.Map{},
	}
}

func (b *BankingService) AddAccount(account *Account) {
	b.accounts.Store(account.accountNumber, account)
}

func (b *BankingService) GetAccount(accountNumber string) (*Account, error) {
	account, ok := b.accounts.Load(accountNumber)
	if !ok {
		return nil, ErrAccountNotFound
	}
	return account.(*Account), nil
}

func (b *BankingService) UpdateAccount(account *Account) {
	b.accounts.Store(account.accountNumber, account)
}

func (b *BankingService) DeleteAccount(accountNumber string) {
	b.accounts.Delete(accountNumber)
}

func (b *BankingService) AddCard(card *Card) {
	b.cards.Store(card.cardNumber, card)
}

func (b *BankingService) GetCard(cardNumber string) (*Card, error) {
	card, ok := b.cards.Load(cardNumber)
	if !ok {
		return nil, ErrCardNotFound
	}
	return card.(*Card), nil
}

func (b *BankingService) ValidateCard(cardNumber, pin string) (*Card, error) {
	card, err := b.GetCard(cardNumber)
	if err != nil {
		return nil, err
	}
	if !card.ValidatePIN(pin) {
		return nil, ErrInvalidPIN
	}
	return card, nil
}

func (b *BankingService) ProcessTransaction(transaction Transaction) error {
	return transaction.Execute()
}
