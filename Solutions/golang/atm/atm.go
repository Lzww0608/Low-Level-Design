package atm

import (
	"fmt"
	"sync/atomic"
)

type ATM struct {
	bankingService *BankingService
	cashDispenser  *CashDispenser
	txnCounter     int64
}

func NewATM(bankingService *BankingService, cashDispenser *CashDispenser) *ATM {
	return &ATM{
		bankingService: bankingService,
		cashDispenser:  cashDispenser,
		txnCounter:     0,
	}
}

// AuthenticateUser 验证用户的卡和PIN
func (a *ATM) AuthenticateUser(cardNumber, pin string) (*Card, error) {
	return a.bankingService.ValidateCard(cardNumber, pin)
}

// GetBalance 查询账户余额
func (a *ATM) GetBalance(cardNumber, pin string) (float64, error) {
	// 验证用户
	card, err := a.AuthenticateUser(cardNumber, pin)
	if err != nil {
		return 0, err
	}

	// 获取账户
	account, err := a.bankingService.GetAccount(card.GetAccountNumber())
	if err != nil {
		return 0, err
	}

	return account.GetBalance(), nil
}

// WithdrawCash 取款
func (a *ATM) WithdrawCash(cardNumber, pin string, amount float64) error {
	// 验证金额
	if amount <= 0 {
		return ErrInvalidAmount
	}

	// 验证用户
	card, err := a.AuthenticateUser(cardNumber, pin)
	if err != nil {
		return err
	}

	// 获取账户
	account, err := a.bankingService.GetAccount(card.GetAccountNumber())
	if err != nil {
		return err
	}

	// 检查ATM现金是否充足
	if err := a.cashDispenser.DispenseCash(int(amount)); err != nil {
		return err
	}

	// 创建并执行取款交易
	transaction := NewWithdrawalTransaction(a.GenerateTransactionID(), account, amount)
	if err := a.bankingService.ProcessTransaction(transaction); err != nil {
		// 如果交易失败，将现金退回ATM
		a.cashDispenser.AddCash(int(amount))
		return err
	}

	return nil
}

// DepositCash 存款
func (a *ATM) DepositCash(cardNumber, pin string, amount float64) error {
	// 验证金额
	if amount <= 0 {
		return ErrInvalidAmount
	}

	// 验证用户
	card, err := a.AuthenticateUser(cardNumber, pin)
	if err != nil {
		return err
	}

	// 获取账户
	account, err := a.bankingService.GetAccount(card.GetAccountNumber())
	if err != nil {
		return err
	}

	// 创建并执行存款交易
	transaction := NewDepositTransaction(a.GenerateTransactionID(), account, amount)
	if err := a.bankingService.ProcessTransaction(transaction); err != nil {
		return err
	}

	// 将现金加入ATM
	a.cashDispenser.AddCash(int(amount))

	return nil
}

// GenerateTransactionID 生成唯一的交易ID
func (a *ATM) GenerateTransactionID() string {
	return fmt.Sprintf("TXN-%d", atomic.AddInt64(&a.txnCounter, 1))
}
