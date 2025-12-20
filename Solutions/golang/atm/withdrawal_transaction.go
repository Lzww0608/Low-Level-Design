package atm

import "errors"

type WithdrawalTransaction struct {
	BaseTransaction
}

func NewWithdrawalTransaction(txnID string, account *Account, amount float64) *WithdrawalTransaction {
	return &WithdrawalTransaction{
		BaseTransaction: BaseTransaction{
			TransactionID: txnID,
			Account:       account,
			Amount:        amount,
		},
	}
}

func (t *WithdrawalTransaction) Execute() error {
	if t.Account == nil {
		return errors.New("account is nil")
	}
	if t.Amount <= 0 {
		return ErrInvalidAmount
	}
	return t.Account.Debit(t.Amount)
}
