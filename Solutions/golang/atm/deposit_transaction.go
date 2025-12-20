package atm

import "errors"

type DepositTransaction struct {
	BaseTransaction
}

func NewDepositTransaction(txnID string, account *Account, amount float64) *DepositTransaction {
	return &DepositTransaction{
		BaseTransaction: BaseTransaction{
			TransactionID: txnID,
			Account:       account,
			Amount:        amount,
		},
	}
}

func (t *DepositTransaction) Execute() error {
	if t.Account == nil {
		return errors.New("account is nil")
	}
	if t.Amount <= 0 {
		return ErrInvalidAmount
	}
	return t.Account.Credit(t.Amount)
}
